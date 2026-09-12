// Package gateway contiene los adaptadores de salida que implementan el
// puerto service.PaymentProcessor — la comunicación con proveedores externos
// de cobro. MockProcessor es la implementación de referencia para desarrollo
// y demostración: no hace ninguna llamada de red real.
package gateway

import (
	"context"
	"errors"
	"fmt"

	"github.com/BienStudios/BienEnd/payment_gateway/internal/domain"
	"github.com/google/uuid"
)

// ErrCardDeclined es un error de negocio (no retriable): reintentar el mismo
// cobro no cambia el resultado, a diferencia de un timeout de red.
var ErrCardDeclined = errors.New("tarjeta rechazada por el proveedor")

// declinedTestToken es un token "mágico" reservado para pruebas: cualquier
// pago que use este CardToken se rechaza intencionalmente, para poder
// ejercitar el camino de fallo sin infraestructura real.
const declinedTestToken = "tok_test_declined"

// MockProcessor implementa service.PaymentProcessor simulando un proveedor
// de pagos en modo sandbox. Es seguro para uso concurrente: no mantiene
// estado mutable compartido entre llamadas.
type MockProcessor struct{}

// NewMockProcessor construye el procesador simulado.
func NewMockProcessor() *MockProcessor {
	return &MockProcessor{}
}

// Charge simula un cobro: aprueba cualquier pago salvo que use el token de
// prueba reservado para simular un rechazo.
func (m *MockProcessor) Charge(_ context.Context, p domain.Payment) (string, error) {
	if p.CardToken == declinedTestToken {
		return "", fmt.Errorf("%w: token de prueba de rechazo", ErrCardDeclined)
	}
	// providerRef simulado, con el mismo formato que usaría un proveedor real
	// (ej. "ch_..." en Stripe), para que el resto del sistema lo trate igual.
	return "ch_mock_" + uuid.NewString(), nil
}

// Refund simula la reversión de un cobro previamente aprobado.
func (m *MockProcessor) Refund(_ context.Context, providerRef string) error {
	if providerRef == "" {
		return errors.New("providerRef vacío: no hay cobro que revertir")
	}
	return nil
}
