package domain_test

import (
	"testing"

	"github.com/BienStudios/BienEnd/payment_gateway/internal/domain"
)

func TestPayment_CanTransitionTo(t *testing.T) {
	tests := []struct {
		name    string
		current domain.PaymentStatus
		target  domain.PaymentStatus
		want    bool
	}{
		{"pending a approved es válido", domain.StatusPending, domain.StatusApproved, true},
		{"pending a failed es válido", domain.StatusPending, domain.StatusFailed, true},
		{"approved a refunded es válido", domain.StatusApproved, domain.StatusRefunded, true},
		{"approved a pending es inválido", domain.StatusApproved, domain.StatusPending, false},
		{"failed es estado terminal", domain.StatusFailed, domain.StatusApproved, false},
		{"refunded es estado terminal", domain.StatusRefunded, domain.StatusApproved, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := domain.Payment{Status: tt.current}
			got := p.CanTransitionTo(tt.target)
			if got != tt.want {
				t.Errorf("CanTransitionTo(%s→%s) = %v, quería %v", tt.current, tt.target, got, tt.want)
			}
		})
	}
}

func TestNewPayment_ValidaInvariantes(t *testing.T) {
	_, err := domain.NewPayment("pay_1", 0, "USD", "tok_x")
	if err == nil {
		t.Error("se esperaba error por monto inválido")
	}

	_, err = domain.NewPayment("pay_1", 1500, "US", "tok_x")
	if err == nil {
		t.Error("se esperaba error por moneda inválida")
	}

	p, err := domain.NewPayment("pay_1", 1500, "USD", "tok_x")
	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if p.Status != domain.StatusPending {
		t.Errorf("estado inicial = %s, quería %s", p.Status, domain.StatusPending)
	}
}

func TestPayment_Refund_RequiereEstadoAprobado(t *testing.T) {
	p := domain.Payment{Status: domain.StatusPending}
	if err := p.Refund(); err == nil {
		t.Error("se esperaba error al reembolsar un pago pendiente")
	}

	p.Status = domain.StatusApproved
	if err := p.Refund(); err != nil {
		t.Errorf("no se esperaba error: %v", err)
	}
	if p.Status != domain.StatusRefunded {
		t.Errorf("estado tras Refund = %s, quería %s", p.Status, domain.StatusRefunded)
	}
}
