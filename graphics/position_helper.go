package graphics

import (
	"window/common"
)

type PositionHelper struct {
	hW float32
	hH float32
}

func NewPositionHelper(cfg *common.Config) *PositionHelper {
	return &PositionHelper{
		hW: float32(cfg.Window.W / 2),
		hH: float32(cfg.Window.H / 2),
	}
}

// convert x axis screen pos to openGL pos
func (ph *PositionHelper) CalcX(x float32) float32 {
	return x/ph.hW - 1
}

// convert y axis screen pos to openGL pos
func (ph *PositionHelper) CalcY(y float32) float32 {
	return y/ph.hH - 1
}

// convert width on openGL values
func (ph *PositionHelper) CalcW(w float32) float32 {
	return w / ph.hW
}

// convert height on openGL values
func (ph *PositionHelper) CalcH(h float32) float32 {
	return h / ph.hH
}
