package graphics

import (
	"github.com/Gregmus2/simple-engine/common"
)

type PercentToPosConverter struct {
	w float32
	h float32
}

func NewPercentToPosConverter(cfg *common.Config) *PercentToPosConverter {
	return &PercentToPosConverter{
		w: float32(cfg.Window.W),
		h: float32(cfg.Window.H),
	}
}

func (ph *PercentToPosConverter) X(x float64) float64 {
	return x / 100 * float64(ph.w)
}

func (ph *PercentToPosConverter) Y(y float64) float64 {
	return y / 100 * float64(ph.h)
}

func (ph *PercentToPosConverter) Radius(r float32) float32 {
	return r / 100 * ph.w
}
