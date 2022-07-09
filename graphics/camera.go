package graphics

import "github.com/Gregmus2/simple-engine/common"

type Camera struct {
	x, y, scale, baseScale float32
	movingX, movingY       int
	speed                  float32
	hW, hH                 float32
}

func NewCamera(cfg *common.Config) *Camera {
	return &Camera{
		x:         0,
		y:         0,
		scale:     cfg.Graphics.Scale,
		baseScale: cfg.Graphics.Scale,
		speed:     cfg.Camera.Speed,
		hW:        float32(cfg.Window.W / 2),
		hH:        float32(cfg.Window.H / 2),
	}
}

func (c *Camera) Position() (float32, float32) {
	return (c.x + c.hW - c.hW*(c.scale/c.baseScale)) / c.scale,
		(c.y + c.hH - c.hH*(c.scale/c.baseScale)) / c.scale
}

func (c *Camera) Scale() float32 {
	return c.scale
}

func (c *Camera) Move(x, y int) {
	c.x += float32(x) * c.speed
	c.y += float32(y) * c.speed
}

func (c *Camera) Zoom(scale float32) {
	c.scale *= scale
}

func (c *Camera) Moving(x, y int) {
	c.movingX += x
	c.movingY += y
}

func (c *Camera) StopMoving(x, y int) {
	c.movingX -= x
	c.movingY -= y
}

func (c *Camera) Update() {
	c.Move(c.movingX, c.movingY)
}
