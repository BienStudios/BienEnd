//go:build postgres

// Este archivo solo se compila con el tag de build "postgres"
// (`go build -tags postgres ./...`), para que el proyecto corra por defecto
// con MemoryPaymentRepository sin requerir una base de datos real ni la
// dependencia jackc/pgx/v5. Antes de usarlo:
//
//	go get github.com/jackc/pgx/v5/pgxpool
//	go build -tags postgres ./...
package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/BienStudios/BienEnd/payment_gateway/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresPaymentRepository implementa service.PaymentRepository usando pgx,
// contra el esquema `payments`.
type PostgresPaymentRepository struct {
	pool *pgxpool.Pool
}

// NewPostgresPaymentRepository construye el repositorio a partir de un pool
// ya configurado (ver cmd/api/main.go para cómo se abre el pool).
func NewPostgresPaymentRepository(pool *pgxpool.Pool) *PostgresPaymentRepository {
	return &PostgresPaymentRepository{pool: pool}
}

// Save inserta un pago nuevo en estado pending.
func (r *PostgresPaymentRepository) Save(ctx context.Context, p domain.Payment) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO payments.payment (id, amount, currency, status, card_token, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		p.ID, p.Amount, p.Currency, p.Status, p.CardToken, p.CreatedAt, p.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("insertando pago: %w", err)
	}
	return nil
}

// FindByID busca un pago por ID, traduciendo sql.ErrNoRows (vía pgx.ErrNoRows)
// al error de dominio correspondiente — la capa de servicio nunca ve un error
// crudo de infraestructura para el caso "no encontrado".
func (r *PostgresPaymentRepository) FindByID(ctx context.Context, id string) (domain.Payment, error) {
	var p domain.Payment
	err := r.pool.QueryRow(ctx, `
		SELECT id, amount, currency, status, card_token, COALESCE(provider_ref, ''), created_at, updated_at
		FROM payments.payment WHERE id = $1`, id,
	).Scan(&p.ID, &p.Amount, &p.Currency, &p.Status, &p.CardToken, &p.ProviderRef, &p.CreatedAt, &p.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Payment{}, domain.ErrPaymentNotFound
	}
	if err != nil {
		return domain.Payment{}, fmt.Errorf("buscando pago %s: %w", id, err)
	}
	return p, nil
}

// Update sobrescribe el estado y la referencia del proveedor de un pago
// existente. En un sistema real, esta escritura se hace dentro de la misma
// transacción que inserta la entrada del ledger.
func (r *PostgresPaymentRepository) Update(ctx context.Context, p domain.Payment) error {
	tag, err := r.pool.Exec(ctx, `
		UPDATE payments.payment
		SET status = $1, provider_ref = $2, updated_at = $3
		WHERE id = $4`,
		p.Status, p.ProviderRef, p.UpdatedAt, p.ID,
	)
	if err != nil {
		return fmt.Errorf("actualizando pago %s: %w", p.ID, err)
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrPaymentNotFound
	}
	return nil
}
