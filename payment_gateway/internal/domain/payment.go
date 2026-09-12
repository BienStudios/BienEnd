// Package domain contiene las entidades y reglas de negocio puras del
// microservicio de pagos. No importa nada de Gin, bases de datos ni
// proveedores externos — es el centro del hexágono.
package domain

import (
	"fmt"
	"time"
)

// PaymentStatus enumera los estados posibles de un pago.
type PaymentStatus string

const (
	StatusPending  PaymentStatus = "pending"
	StatusApproved PaymentStatus = "approved"
	StatusFailed   PaymentStatus = "failed"
	StatusRefunded PaymentStatus = "refunded"
)

// validTransitions define la tabla de transiciones de estado permitidas.
// Cualquier transición que no figure aquí es inválida y TransitionTo la rechaza.
var validTransitions = map[PaymentStatus][]PaymentStatus{
	StatusPending:  {StatusApproved, StatusFailed},
	StatusApproved: {StatusRefunded},
	StatusFailed:   {}, // estado terminal
	StatusRefunded: {}, // estado terminal
}

// Payment es la entidad central del dominio.
//
// Amount se expresa siempre en la unidad mínima de la moneda (ej. centavos)
// para evitar imprecisiones de punto flotante en operaciones monetarias.
type Payment struct {
	ID          string
	Amount      int64
	Currency    string
	Status      PaymentStatus
	CardToken   string // token opaco del método de pago; nunca el PAN crudo
	ProviderRef string // referencia devuelta por el proveedor externo tras el cobro
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewPayment construye un Payment válido en estado pending, o devuelve un
// error de dominio si los datos de entrada violan una invariante básica.
//
// La validación de forma (JSON bien formado, tipos correctos) ocurre en la
// capa HTTP; esta validación es la de negocio, y se aplica sin importar quién
// llame a este constructor.
func NewPayment(id string, amount int64, currency, cardToken string) (Payment, error) {
	if amount <= 0 {
		return Payment{}, fmt.Errorf("%w: %d", ErrInvalidAmount, amount)
	}
	if len(currency) != 3 {
		return Payment{}, fmt.Errorf("%w: %s", ErrInvalidCurrency, currency)
	}

	now := time.Now().UTC()
	return Payment{
		ID:        id,
		Amount:    amount,
		Currency:  currency,
		CardToken: cardToken,
		Status:    StatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// CanTransitionTo verifica si una transición de estado es válida según las
// reglas de negocio, sin mutar el pago.
func (p Payment) CanTransitionTo(next PaymentStatus) bool {
	for _, allowed := range validTransitions[p.Status] {
		if allowed == next {
			return true
		}
	}
	return false
}

// TransitionTo aplica la transición si es válida, o devuelve
// ErrInvalidTransition envuelto con el detalle de origen/destino.
func (p *Payment) TransitionTo(next PaymentStatus) error {
	if !p.CanTransitionTo(next) {
		return fmt.Errorf("%w: de %s a %s", ErrInvalidTransition, p.Status, next)
	}
	p.Status = next
	p.UpdatedAt = time.Now().UTC()
	return nil
}

// MarkApproved transiciona el pago a aprobado, registrando la referencia del
// proveedor externo que efectivamente realizó el cobro.
func (p *Payment) MarkApproved(providerRef string) error {
	if err := p.TransitionTo(StatusApproved); err != nil {
		return err
	}
	p.ProviderRef = providerRef
	return nil
}

// MarkFailed transiciona el pago a fallido. No requiere providerRef porque
// el cobro nunca llegó a completarse del lado del proveedor.
func (p *Payment) MarkFailed() error {
	return p.TransitionTo(StatusFailed)
}

// Refund transiciona el pago a reembolsado, validando primero que el estado
// actual lo permita — es el dominio, no la capa de servicio, quien protege
// esta invariante.
func (p *Payment) Refund() error {
	if !p.CanTransitionTo(StatusRefunded) {
		return fmt.Errorf("%w: estado actual %s", ErrNotRefundable, p.Status)
	}
	return p.TransitionTo(StatusRefunded)
}

// IsRefundable expone la misma regla de negocio de forma consultable, sin
// mutar el pago — útil para que la capa HTTP decida si mostrar la acción.
func (p Payment) IsRefundable() bool {
	return p.CanTransitionTo(StatusRefunded)
}
