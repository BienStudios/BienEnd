package handler

import (
	"errors"
	"net/http"

	"github.com/BienStudios/BienEnd/payment_gateway/internal/domain"
	"github.com/BienStudios/BienEnd/payment_gateway/internal/gateway"
	"github.com/BienStudios/BienEnd/payment_gateway/internal/service"
	"github.com/gin-gonic/gin"
)

// PaymentHandler es el adaptador de entrada HTTP para el recurso "payments".
// Depende únicamente de *service.PaymentService (una struct concreta, no una
// interfaz) — sigue la convención de Go de "aceptar interfaces, devolver
// structs concretos": quien construye el handler ya decidió qué implementación
// de cada puerto usa PaymentService.
type PaymentHandler struct {
	service *service.PaymentService
}

// NewPaymentHandler construye el handler, inyectando el servicio de aplicación.
func NewPaymentHandler(s *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: s}
}

// RegisterRoutes agrupa las rutas de pagos bajo /api/v26.00.0000_ALPHA/payments,
// incluyendo el middleware de idempotencia SOLO en la creación.
func (h *PaymentHandler) RegisterRoutes(router gin.IRouter, idempotencyMiddleware gin.HandlerFunc) {
	payments := router.Group("/payments")
	{
		payments.POST("", idempotencyMiddleware, h.CreatePayment)
		payments.GET("/:id", h.GetPayment)
		payments.POST("/:id/refund", h.RefundPayment)
	}
}

// CreatePayment godoc
//
// @Summary      Crear un nuevo pago
// @Description  Procesa un cobro con el proveedor configurado y persiste el resultado.
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        Idempotency-Key  header    string                true  "Clave de idempotencia"
// @Param        request          body      CreatePaymentRequest  true  "Datos del pago"
// @Success      201  {object}  PaymentResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      402  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/payments [post]
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: "invalid_payload", Message: err.Error()})
		return
	}

	idempotencyKey := c.GetString("idempotency_key") // provisto por el middleware de idempotencia

	payment, err := h.service.CreatePayment(c.Request.Context(), service.CreatePaymentInput{
		IdempotencyKey: idempotencyKey,
		Amount:         req.Amount,
		Currency:       req.Currency,
		CardToken:      req.CardToken,
	})
	if err != nil {
		abortWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, toPaymentResponse(payment))
}

// GetPayment godoc
//
// @Summary      Consultar un pago por ID
// @Tags         payments
// @Produce      json
// @Param        id path string true "ID del pago"
// @Success      200 {object} PaymentResponse
// @Failure      404 {object} ErrorResponse
// @Router       /api/v1/payments/{id} [get]
func (h *PaymentHandler) GetPayment(c *gin.Context) {
	id := c.Param("id")

	payment, err := h.service.GetPayment(c.Request.Context(), id)
	if err != nil {
		abortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, toPaymentResponse(payment))
}

// RefundPayment godoc
//
// @Summary      Reembolsar un pago previamente aprobado
// @Tags         payments
// @Produce      json
// @Param        id path string true "ID del pago"
// @Success      200 {object} PaymentResponse
// @Failure      404 {object} ErrorResponse
// @Failure      409 {object} ErrorResponse
// @Router       /api/v1/payments/{id}/refund [post]
func (h *PaymentHandler) RefundPayment(c *gin.Context) {
	id := c.Param("id")

	payment, err := h.service.RefundPayment(c.Request.Context(), id)
	if err != nil {
		abortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, toPaymentResponse(payment))
}

// abortWithError centraliza el mapeo de errores de dominio/aplicación a
// códigos HTTP — ningún handler decide el status a mano.
func abortWithError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrPaymentNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorResponse{Code: "not_found", Message: "pago no encontrado"})
	case errors.Is(err, domain.ErrInvalidAmount), errors.Is(err, domain.ErrInvalidCurrency):
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{Code: "invalid_payment", Message: err.Error()})
	case errors.Is(err, domain.ErrNotRefundable):
		c.AbortWithStatusJSON(http.StatusConflict, ErrorResponse{Code: "not_refundable", Message: err.Error()})
	case errors.Is(err, gateway.ErrCardDeclined):
		c.AbortWithStatusJSON(http.StatusPaymentRequired, ErrorResponse{Code: "card_declined", Message: "el proveedor rechazó el pago"})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Code: "internal_error", Message: "error interno"})
	}
}
