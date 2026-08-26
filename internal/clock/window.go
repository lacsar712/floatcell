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
	return ProcessWindowOpen(w.clk, anchor, w.duration)
}

func (w *AerateWindow) Require(anchor time.Time) error {
	if ProcessWindowOpen(w.clk, anchor, w.duration) {
		return nil
	}
	if ProcessWindowClosed(w.clk, anchor, w.duration) {
		return model.ErrSpargeHold
	}
	return model.ErrSpargeHold
}
