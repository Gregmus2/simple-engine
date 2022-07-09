package graphics

import "github.com/Gregmus2/simple-engine/common"

type Camera struct {
	x, y, scale      float32
	movingX, movingY int
	speed            float32
}

func NewCamera(cfg *common.Config) *Camera {
	return &Camera{
		x:     0,
		y:     0,
		scale: 1,
		speed: cfg.Camera.Speed,
	}
}

func (c *Camera) Position() (float32, float32) {
	return c.x, c.y
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
