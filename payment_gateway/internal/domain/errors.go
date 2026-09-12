package domain

import "errors"

// Definiciones para errores comunes
var (
	ErrInvalidAmount     = errors.New("el monto debe ser mayor a cero")
	ErrInvalidCurrency   = errors.New("código de moneda inválido")
	ErrPaymentNotFound   = errors.New("pago no encontrado")
	ErrInvalidTransition = errors.New("transición de estado inválida")
	ErrNotRefundable     = errors.New("el pago no admite reembolso en su estado actual")
)
