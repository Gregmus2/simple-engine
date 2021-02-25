package graphics

import (
	"github.com/Gregmus2/simple-engine/internal/common"
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
	return x/100*float64(ph.w) - float64(ph.w)/2
}

func (ph *PercentToPosConverter) W(w float64) float64 {
	return w / 100 * float64(ph.w)
}

func (ph *PercentToPosConverter) Y(y float64) float64 {
	return y/100*float64(ph.h) - float64(ph.h)/2
}

func (ph *PercentToPosConverter) H(h float64) float64 {
	return h / 100 * float64(ph.h)
}

func (ph *PercentToPosConverter) D(d float64) float64 {
	return d / 100 * 100
}

func (ph *PercentToPosConverter) Radius(r float32) float32 {
	return r / 100 * ph.w
}
