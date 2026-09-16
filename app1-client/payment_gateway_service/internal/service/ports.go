// Package service contiene la lógica de aplicación: orquesta las entidades
// del dominio con los puertos de persistencia, procesamiento de pagos y
// publicación de eventos. No conoce Gin, SQL concreto, ni el proveedor de
// cobro real — solo las interfaces que se implementan en las capas externas
package service

import (
	"context"

	"github.com/BienStudios/BienEnd/payment_gateway/internal/domain"
)

// PaymentRepository es el puerto de salida hacia la persistencia. La capa de
// aplicación no sabe si detrás hay PostgreSQL, una base en memoria, o Mongo.
type PaymentRepository interface {
	Save(ctx context.Context, p domain.Payment) error
	FindByID(ctx context.Context, id string) (domain.Payment, error)
	Update(ctx context.Context, p domain.Payment) error
}

// PaymentProcessor es el puerto de salida hacia el proveedor externo de cobro
// (Stripe, Adyen, MercadoPago, o un mock de sandbox para desarrollo/tests).
//
// Las implementaciones deben:
//   - Ser seguras para uso concurrente.
//   - Distinguir errores retriables (timeout, 503) de no retriables (tarjeta
//     rechazada), para que la capa de aplicación no reintente lo que no debe.
//   - No loguear ni persistir el CardToken recibido.
type PaymentProcessor interface {
	// Charge intenta cobrar el monto de p usando su CardToken como método de
	// pago. Devuelve la referencia del proveedor (providerRef) en caso de éxito.
	Charge(ctx context.Context, p domain.Payment) (providerRef string, err error)

	// Refund revierte un cobro previamente exitoso, identificado por providerRef.
	Refund(ctx context.Context, providerRef string) error
}

// EventPublisher es el puerto de salida hacia el bus de eventos (RabbitMQ,
// Kafka, SNS, o un publisher no-op para desarrollo local).
type EventPublisher interface {
	PublishPaymentCreated(ctx context.Context, p domain.Payment) error
	PublishPaymentFailed(ctx context.Context, p domain.Payment, reason string) error
}

// IdempotencyStore es el puerto de salida para garantizar que un mismo
// request (identificado por una clave de idempotencia) nunca se procese dos
// veces, incluso con múltiples réplicas del servicio.
type IdempotencyStore interface {
	// TryReserve intenta reservar la clave de forma atómica. Devuelve false si
	// otra ejecución ya reservó (o completó) esta misma clave.
	TryReserve(ctx context.Context, key string) (reserved bool, err error)

	// Release libera una clave reservada que no llegó a completarse (ej. el
	// procesamiento falló con un error no relacionado al negocio), permitiendo
	// un reintento legítimo del mismo request.
	Release(ctx context.Context, key string) error
}
