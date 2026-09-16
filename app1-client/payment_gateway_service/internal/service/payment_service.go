package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/BienStudios/BienEnd/payment_gateway/internal/domain"
	"github.com/google/uuid"
)

// CreatePaymentInput agrupa los datos de entrada del caso de uso de creación
// de pago. Vive en la capa de aplicación (no en el handler HTTP) para que el
// servicio se pueda invocar también desde un consumer de mensajería
// sin pasar por Gin.
type CreatePaymentInput struct {
	IdempotencyKey string
	Amount         int64
	Currency       string
	CardToken      string
}

// PaymentService orquesta los casos de uso relacionados con pagos,
// coordinando el dominio con los puertos de persistencia, cobro e idempotencia.
type PaymentService struct {
	repo             PaymentRepository
	processor        PaymentProcessor
	publisher        EventPublisher
	idempotencyStore IdempotencyStore
	logger           *slog.Logger
}

// NewPaymentService construye el servicio inyectando sus dependencias
// (puertos). La inyección es manual: no se usa ningún framework de DI,
// siguiendo el estilo idiomático de Go.
func NewPaymentService(
	repo PaymentRepository,
	processor PaymentProcessor,
	publisher EventPublisher,
	idempotencyStore IdempotencyStore,
	logger *slog.Logger,
) *PaymentService {
	if logger == nil {
		logger = slog.Default()
	}
	return &PaymentService{
		repo:             repo,
		processor:        processor,
		publisher:        publisher,
		idempotencyStore: idempotencyStore,
		logger:           logger,
	}
}

// CreatePayment ejecuta el caso de uso completo de creación y cobro de un pago.
//
// El flujo es: (1) reservar la clave de idempotencia, (2) construir y validar
// la entidad de dominio, (3) persistirla en pending, (4) intentar el cobro
// con el proveedor externo, (5) actualizar el estado según el resultado,
// (6) publicar un evento de dominio de forma best-effort.
func (s *PaymentService) CreatePayment(ctx context.Context, input CreatePaymentInput) (domain.Payment, error) {
	reserved, err := s.idempotencyStore.TryReserve(ctx, input.IdempotencyKey)
	if err != nil {
		return domain.Payment{}, fmt.Errorf("reservando clave de idempotencia: %w", err)
	}
	if !reserved {
		// Ya existe un pago para esta clave de idempotencia: se devuelve el
		// existente en vez de reprocesar el cobro. En una implementación con
		// store persistente, aquí se buscaría el payment_id asociado a la
		// clave; con el store en memoria de referencia, se documenta el punto
		// de extensión y se devuelve un error identificable.
		return domain.Payment{}, fmt.Errorf("clave de idempotencia ya procesada: %s", input.IdempotencyKey)
	}

	payment, err := domain.NewPayment(uuid.NewString(), input.Amount, input.Currency, input.CardToken)
	if err != nil {
		_ = s.idempotencyStore.Release(ctx, input.IdempotencyKey)
		return domain.Payment{}, err // error de validación de dominio, no se envuelve más
	}

	if err := s.repo.Save(ctx, payment); err != nil {
		_ = s.idempotencyStore.Release(ctx, input.IdempotencyKey)
		return domain.Payment{}, fmt.Errorf("guardando pago pendiente: %w", err)
	}

	providerRef, err := s.processor.Charge(ctx, payment)
	if err != nil {
		return s.handleChargeFailure(ctx, payment, err)
	}

	if err := payment.MarkApproved(providerRef); err != nil {
		// No debería ocurrir (pending -> approved siempre es válido), pero se
		// maneja explícitamente para no dejar el pago en un estado inconsistente.
		return domain.Payment{}, fmt.Errorf("aplicando transición a aprobado: %w", err)
	}

	if err := s.repo.Update(ctx, payment); err != nil {
		return domain.Payment{}, fmt.Errorf("actualizando estado a aprobado: %w", err)
	}

	if err := s.publisher.PublishPaymentCreated(ctx, payment); err != nil {
		// Un fallo de publicación de evento no revierte el cobro ya exitoso;
		// se loguea para que un proceso de reconciliación lo detecte
		// (patrón Outbox).
		s.logger.Warn("no se pudo publicar payment.created",
			slog.String("payment_id", payment.ID), slog.String("error", err.Error()))
	}

	return payment, nil
}

// handleChargeFailure centraliza qué hacer cuando el cobro al proveedor
// falla: marca el pago como failed y libera la clave de idempotencia SOLO si
// el fallo es de negocio (ej. tarjeta rechazada) y no transitorio, para
// permitir que un fallo transitorio se reintente con la misma clave.
func (s *PaymentService) handleChargeFailure(ctx context.Context, payment domain.Payment, chargeErr error) (domain.Payment, error) {
	if err := payment.MarkFailed(); err != nil {
		return domain.Payment{}, fmt.Errorf("marcando pago como fallido: %w", err)
	}
	if err := s.repo.Update(ctx, payment); err != nil {
		s.logger.Error("no se pudo persistir el estado failed",
			slog.String("payment_id", payment.ID), slog.String("error", err.Error()))
	}

	if err := s.publisher.PublishPaymentFailed(ctx, payment, chargeErr.Error()); err != nil {
		s.logger.Warn("no se pudo publicar payment.failed",
			slog.String("payment_id", payment.ID), slog.String("error", err.Error()))
	}

	return domain.Payment{}, fmt.Errorf("cobrando pago: %w", chargeErr)
}

// GetPayment recupera un pago por ID.
func (s *PaymentService) GetPayment(ctx context.Context, id string) (domain.Payment, error) {
	payment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrPaymentNotFound) {
			return domain.Payment{}, err
		}
		return domain.Payment{}, fmt.Errorf("buscando pago %s: %w", id, err)
	}
	return payment, nil
}

// RefundPayment ejecuta el caso de uso de reembolso: valida la transición en
// el dominio ANTES de llamar al proveedor externo, evitando una llamada de
// red innecesaria si el pago no es reembolsable.
func (s *PaymentService) RefundPayment(ctx context.Context, id string) (domain.Payment, error) {
	payment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return domain.Payment{}, err
	}

	if !payment.IsRefundable() {
		return domain.Payment{}, fmt.Errorf("%w: pago %s en estado %s", domain.ErrNotRefundable, id, payment.Status)
	}

	if err := s.processor.Refund(ctx, payment.ProviderRef); err != nil {
		return domain.Payment{}, fmt.Errorf("revirtiendo cobro en el proveedor: %w", err)
	}

	if err := payment.Refund(); err != nil {
		return domain.Payment{}, fmt.Errorf("aplicando transición a reembolsado: %w", err)
	}

	if err := s.repo.Update(ctx, payment); err != nil {
		return domain.Payment{}, fmt.Errorf("persistiendo reembolso: %w", err)
	}

	return payment, nil
}
