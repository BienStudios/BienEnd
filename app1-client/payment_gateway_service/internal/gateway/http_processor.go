package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/BienStudios/BienEnd/payment_gateway/internal/domain"
)

// HTTPProcessor implementa service.PaymentProcessor contra un proveedor real
// expuesto por HTTP/JSON (ej. una pasarela tipo Stripe/MercadoPago). Se deja
// como referencia de producción: MockProcessor es la que usa cmd/api/main.go
// por defecto para poder correr el servicio sin credenciales reales.
type HTTPProcessor struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

// NewHTTPProcessor construye el cliente con un timeout explícito — nunca se
// debe dejar una llamada a un proveedor externo sin límite de tiempo.
func NewHTTPProcessor(baseURL, apiKey string) *HTTPProcessor {
	return &HTTPProcessor{
		baseURL: baseURL,
		apiKey:  apiKey,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

type chargeRequest struct {
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	CardToken string `json:"card_token"`
}

type chargeResponse struct {
	ProviderRef string `json:"provider_ref"`
	Declined    bool   `json:"declined"`
	Reason      string `json:"reason,omitempty"`
}

// Charge realiza el cobro contra el proveedor externo, distinguiendo un
// rechazo de negocio (ErrCardDeclined, no retriable) de un error de
// transporte/timeout (retriable por la capa que llama a este método).
func (h *HTTPProcessor) Charge(ctx context.Context, p domain.Payment) (string, error) {
	body, err := json.Marshal(chargeRequest{Amount: p.Amount, Currency: p.Currency, CardToken: p.CardToken})
	if err != nil {
		return "", fmt.Errorf("serializando request de cobro: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.baseURL+"/charges", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("construyendo request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+h.apiKey)

	resp, err := h.client.Do(req)
	if err != nil {
		// Error de red/timeout: se propaga sin envolver en ErrCardDeclined,
		// para que la capa de servicio lo trate como potencialmente retriable.
		return "", fmt.Errorf("llamando al proveedor de pagos: %w", err)
	}
	defer resp.Body.Close()

	var parsed chargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", fmt.Errorf("decodificando respuesta del proveedor: %w", err)
	}

	if resp.StatusCode == http.StatusPaymentRequired || parsed.Declined {
		return "", fmt.Errorf("%w: %s", ErrCardDeclined, parsed.Reason)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("respuesta inesperada del proveedor: %d", resp.StatusCode)
	}

	return parsed.ProviderRef, nil
}

// Refund revierte un cobro previamente exitoso en el proveedor externo.
func (h *HTTPProcessor) Refund(ctx context.Context, providerRef string) error {
	if providerRef == "" {
		return errors.New("providerRef vacío: no hay cobro que revertir")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.baseURL+"/refunds/"+providerRef, nil)
	if err != nil {
		return fmt.Errorf("construyendo request de reembolso: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+h.apiKey)

	resp, err := h.client.Do(req)
	if err != nil {
		return fmt.Errorf("llamando al proveedor para reembolsar: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("el proveedor rechazó el reembolso: %d", resp.StatusCode)
	}
	return nil
}
