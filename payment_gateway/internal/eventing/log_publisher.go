// Package eventing contiene los adaptadores de salida que implementan el
// puerto service.EventPublisher. LogPublisher solo deja constancia por log
// de cada evento — útil para desarrollo local y para el arranque por defecto
// del servicio sin depender de RabbitMQ.
package eventing

import (
	"context"
	"log/slog"

	"github.com/BienStudios/BienEnd/payment_gateway/internal/domain"
)

// LogPublisher implementa service.EventPublisher registrando cada evento
// como una línea de log estructurado, sin publicarlo a ningún broker real.
type LogPublisher struct {
	logger *slog.Logger
}

// NewLogPublisher construye el publisher de referencia.
func NewLogPublisher(logger *slog.Logger) *LogPublisher {
	if logger == nil {
		logger = slog.Default()
	}
	return &LogPublisher{logger: logger}
}

func (p *LogPublisher) PublishPaymentCreated(_ context.Context, payment domain.Payment) error {
	p.logger.Info("evento: payment.created",
		slog.String("payment_id", payment.ID),
		slog.Int64("amount", payment.Amount),
		slog.String("currency", payment.Currency),
		slog.String("status", string(payment.Status)),
	)
	return nil
}

func (p *LogPublisher) PublishPaymentFailed(_ context.Context, payment domain.Payment, reason string) error {
	p.logger.Info("evento: payment.failed",
		slog.String("payment_id", payment.ID),
		slog.String("reason", reason),
	)
	return nil
}
