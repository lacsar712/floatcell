package app

import (
	"context"
	"errors"
	"testing"

	"github.com/lacsar712/floatcell/internal/config"
	"github.com/lacsar712/floatcell/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	err = a.ValidateMoldDrift(context.Background(), 25.0)
	if err == nil {
		t.Fatal("expected moisture drift violation")
	}
	if !errors.Is(err, model.ErrPulpDrift) {
		t.Fatalf("expected ErrPulpDrift, got %v", err)
	}
}
