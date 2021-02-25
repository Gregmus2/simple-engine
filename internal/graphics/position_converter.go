package graphics

import "github.com/Gregmus2/simple-engine/internal/common"

type PositionConverter struct {
	scale, hW, hH float32
}

func NewPositionConverter(cfg *common.Config) *PositionConverter {
	return &PositionConverter{
		scale: cfg.Graphics.Scale,
		hH:    float32(cfg.Window.H) / 2,
		hW:    float32(cfg.Window.W) / 2,
	}
}

func (c *PositionConverter) Box2DToOpenGL(x, y float32) (float32, float32) {
	return x * c.scale / c.hW, y * c.scale / c.hH
}
