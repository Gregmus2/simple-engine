package graphics

import (
	"github.com/Gregmus2/simple-engine/common"
)

type PosToUnitsConverter struct {
	hW float32
	hH float32
	hD float32
}

func NewPosToUnitsConverter(cfg *common.Config) *PosToUnitsConverter {
	return &PosToUnitsConverter{
		hW: float32(cfg.Window.W / 2),
		hH: float32(cfg.Window.H / 2),
		hD: 50,
	}
}

func (ph *PosToUnitsConverter) X(x float32) float32 {
	return x/ph.hW - 1
}

func (ph *PosToUnitsConverter) Y(y float32) float32 {
	return y/ph.hH - 1
}

func (ph *PosToUnitsConverter) W(w float32) float32 {
	return w / ph.hW
}

func (ph *PosToUnitsConverter) H(h float32) float32 {
	return h / ph.hH
}

func (ph *PosToUnitsConverter) D(d float32) float32 {
	return d / ph.hD
}
