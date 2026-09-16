// Package handler contiene el adaptador de entrada HTTP: traduce requests
// de Gin hacia llamadas al service.PaymentService, y las respuestas del
// dominio hacia JSON — nunca al revés.
package handler

import "github.com/BienStudios/BienEnd/payment_gateway/internal/domain"

// CreatePaymentRequest es el DTO de entrada para crear un pago. Los tags
// `binding` son evaluados por el validador integrado de Gin antes de que el
// handler vea el request.
type CreatePaymentRequest struct {
	Amount    int64  `json:"amount" binding:"required,gt=0,lte=100000000"`
	Currency  string `json:"currency" binding:"required,len=3"`
	CardToken string `json:"card_token" binding:"required"`
}

// PaymentResponse es el DTO de salida: expone solo lo que un cliente externo
// necesita ver, sin filtrar detalles internos del dominio (ej. nunca expone
// CardToken, aunque el dominio lo tenga).
type PaymentResponse struct {
	ID          string `json:"id"`
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Status      string `json:"status"`
	ProviderRef string `json:"provider_ref,omitempty"`
	CreatedAt   string `json:"created_at"`
}

// toPaymentResponse convierte la entidad de dominio al DTO de salida.
func toPaymentResponse(p domain.Payment) PaymentResponse {
	return PaymentResponse{
		ID:          p.ID,
		Amount:      p.Amount,
		Currency:    p.Currency,
		Status:      string(p.Status),
		ProviderRef: p.ProviderRef,
		CreatedAt:   p.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// ErrorResponse estandariza el formato de error de toda la API.
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
