package model

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidID       = errors.New("floatcell: invalid identifier")
	ErrNotFound        = errors.New("floatcell: entity not found")
	ErrConflict        = errors.New("floatcell: state conflict")
	ErrInterlock       = errors.New("floatcell: interlock denied")
	ErrMoistureHold    = errors.New("floatcell: moisture hold active")
	ErrAirflowSetpoint = errors.New("floatcell: airflow setpoint violation")
	ErrFanFault        = errors.New("floatcell: fan fault")
	ErrScheduleEmpty   = errors.New("floatcell: schedule empty")
	ErrGradient        = errors.New("floatcell: moisture gradient violation")
	ErrPulpDrift   = errors.New("floatcell: moisture drift exceeded")
	ErrLevelTrip    = errors.New("floatcell: heat overtemperature")
	ErrSpargeHold    = errors.New("floatcell: gradient hold not satisfied")
	ErrContextCanceled = errors.New("floatcell: operation canceled")
)

type DomainError struct {
	Op   string
	Code string
	Err  error
}

func (e *DomainError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Err != nil {
		return fmt.Sprintf("floatcell %s [%s]: %v", e.Op, e.Code, e.Err)
	}
	return fmt.Sprintf("floatcell %s [%s]", e.Op, e.Code)
}

func (e *DomainError) Unwrap() error { return e.Err }

func Wrap(op, code string, err error) error {
	if err == nil {
		return nil
	}
	return &DomainError{Op: op, Code: code, Err: err}
}

func Is(err, target error) bool   { return errors.Is(err, target) }
func As(err error, target any) bool { return errors.As(err, target) }
