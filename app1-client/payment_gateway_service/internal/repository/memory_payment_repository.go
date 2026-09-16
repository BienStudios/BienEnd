// Package repository contiene los adaptadores de salida que implementan el
// puerto service.PaymentRepository. MemoryPaymentRepository es la
// implementación de referencia para desarrollo y tests: no requiere
// infraestructura externa y permite correr el servicio completo con
// `go run ./cmd/api` sin una base de datos real.
//
// Para producción, ver postgres_payment_repository.go — implementa el mismo
// puerto contra PostgreSQL.
package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/BienStudios/BienEnd/payment_gateway/internal/domain"
)

// MemoryPaymentRepository implementa service.PaymentRepository con un mapa
// protegido por mutex — seguro para acceso concurrente desde múltiples
// goroutines (cada request HTTP se atiende en su propia goroutine en Gin).
type MemoryPaymentRepository struct {
	mu       sync.RWMutex
	payments map[string]domain.Payment
}

// NewMemoryPaymentRepository construye un repositorio en memoria vacío.
func NewMemoryPaymentRepository() *MemoryPaymentRepository {
	return &MemoryPaymentRepository{
		payments: make(map[string]domain.Payment),
	}
}

// Save persiste un pago nuevo. Devuelve error si el ID ya existe, para
// imitar la restricción de clave primaria que tendría una base real.
func (r *MemoryPaymentRepository) Save(_ context.Context, p domain.Payment) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.payments[p.ID]; exists {
		return errors.New("el pago ya existe")
	}
	r.payments[p.ID] = p
	return nil
}

// FindByID busca un pago por ID, devolviendo domain.ErrPaymentNotFound si no existe.
func (r *MemoryPaymentRepository) FindByID(_ context.Context, id string) (domain.Payment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.payments[id]
	if !ok {
		return domain.Payment{}, domain.ErrPaymentNotFound
	}
	return p, nil
}

// Update sobrescribe un pago existente (usado tras cada transición de estado).
func (r *MemoryPaymentRepository) Update(_ context.Context, p domain.Payment) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.payments[p.ID]; !exists {
		return domain.ErrPaymentNotFound
	}
	r.payments[p.ID] = p
	return nil
}
