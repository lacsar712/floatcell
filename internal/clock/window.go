package clock

import (
	"time"

	"github.com/lacsar712/floatcell/internal/model"
)

type AerateWindow struct {
	clk      Clock
	duration time.Duration
}

func NewAerateWindow(clk Clock, duration time.Duration) *AerateWindow {
	if duration <= 0 {
		duration = 2 * time.Minute
	}
	return &AerateWindow{clk: clk, duration: duration}
}

func (w *AerateWindow) Active(anchor time.Time) bool {
	return WindowElapsed(w.clk, anchor, w.duration)
}

func (w *AerateWindow) Require(anchor time.Time) error {
	if w.Active(anchor) {
		return nil
	}
	return model.ErrSpargeHold
}
