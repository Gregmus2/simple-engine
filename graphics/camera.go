package graphics

import (
	"fmt"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Camera struct {
	x, y, scale, baseScale float64
	hW, hH                 float64
	isMoving               bool
	lastX, lastY           float64
}

func NewCamera(cfg *common.Config) *Camera {
	return &Camera{
		x:         0,
		y:         0,
		scale:     float64(cfg.Graphics.Scale),
		baseScale: float64(cfg.Graphics.Scale),
		hW:        float64(cfg.Window.W / 2),
		hH:        float64(cfg.Window.H / 2),
	}
}

func (c *Camera) Position() (float32, float32) {
	return float32((c.x + c.hW - c.hW*(c.scale/c.baseScale)) / c.scale),
		float32((c.y + c.hH - c.hH*(c.scale/c.baseScale)) / c.scale)
}

func (c *Camera) Scale() float32 {
	return float32(c.scale)
}

func (c *Camera) Move(x, y float64) {
	c.x += x
	c.y += -y
	fmt.Printf("%f %f\n", c.x, c.y)
}

func (c *Camera) Zoom(scale float32) {
	c.scale *= float64(scale)
}

func (c *Camera) MouseCallback(button glfw.MouseButton, action glfw.Action) {
	switch {
	case button == glfw.MouseButtonLeft && action == glfw.Press:
		c.isMoving = true
	case button == glfw.MouseButtonLeft && action == glfw.Release:
		c.isMoving = false
		c.lastX, c.lastY = 0, 0
	}
}

func (c *Camera) CursorPositionCallback(x, y float64) {
	if c.isMoving {
		if c.lastX == 0 && c.lastY == 0 {
			c.lastX, c.lastY = x, y
		}

		c.Move(x-c.lastX, y-c.lastY)
		c.lastX, c.lastY = x, y
	}
}

func (c *Camera) ScrollCallback(yoffset float64) {
	c.Zoom(1 + float32(yoffset)*0.1)
}
