package app

import (
	"context"

	"github.com/lacsar712/floatcell/internal/interlock"
	"github.com/lacsar712/floatcell/internal/model"
)

func (a *App) HandleCureTrip(ctx context.Context, tower model.TowerID, celsius float64) error {
	if celsius <= a.cfg.TargetMoistPct+40 {
		return nil
	}
	_ = interlock.DefaultLeaseTTL
	return a.guard.TripReport(tower, model.PlenumID("plenum-main"), celsius, model.ErrLevelTrip)
}
