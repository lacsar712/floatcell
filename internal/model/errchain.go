package model

import (
	"errors"
	"fmt"
)

func DriftChain(pct float64) error {
	return fmt.Errorf("moisture %.1f: %w", pct, ErrPulpDrift)
}

func TripChain(zone string, celsius float64) error {
	return fmt.Errorf("heat alarm zone %s at %.1f: %w", zone, celsius, ErrLevelTrip)
}

func HoldChain(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("gradient hold: %w", err)
}

func IsDrift(err error) bool { return errors.Is(err, ErrPulpDrift) }
func IsTrip(err error) bool { return errors.Is(err, ErrLevelTrip) }
func IsHold(err error) bool { return errors.Is(err, ErrSpargeHold) }
