package repository

import (
	"context"
	"sync"
	"time"
)

// MemoryIdempotencyStore implementa service.IdempotencyStore en memoria, con
// expiración perezosa de claves (se limpian al ser consultadas, no con un
// job de fondo). Es la implementación de referencia para desarrollo; en
// producción con múltiples réplicas se reemplaza por Valkey para que la reserva
// sea atómica entre instancias del servicio.
type MemoryIdempotencyStore struct {
	mu      sync.Mutex
	entries map[string]time.Time // clave -> momento de expiración
	ttl     time.Duration
}

// NewMemoryIdempotencyStore construye el store con un TTL fijo por clave.
func NewMemoryIdempotencyStore(ttl time.Duration) *MemoryIdempotencyStore {
	return &MemoryIdempotencyStore{
		entries: make(map[string]time.Time),
		ttl:     ttl,
	}
}

// TryReserve reserva la clave de forma atómica (protegida por mutex): dos
// goroutines que llamen con la misma clave al mismo tiempo nunca reciben
// ambas `true` — solo la primera en tomar el lock.
func (s *MemoryIdempotencyStore) TryReserve(_ context.Context, key string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if expiresAt, exists := s.entries[key]; exists && time.Now().Before(expiresAt) {
		return false, nil // ya reservada y aún vigente
	}

	s.entries[key] = time.Now().Add(s.ttl)
	return true, nil
}

// Release libera una clave reservada, permitiendo que un reintento legítimo
// (ej. tras un error transitorio) vuelva a procesarse con la misma clave.
func (s *MemoryIdempotencyStore) Release(_ context.Context, key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.entries, key)
	return nil
}
