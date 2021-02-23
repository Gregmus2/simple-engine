package graphics

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	view       mgl32.Mat4
	projection mgl32.Mat4
	hW, hH     float32
}

func NewCamera(cfg *common.Config) common.Camera {
	return &Camera{
		view: mgl32.LookAtV(mgl32.Vec3{0, 0, -3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0}),
		// near -> 3
		//projection: mgl32.Ortho(-1, 1, -1, 1, 0.1, 100),
		projection: mgl32.Perspective(mgl32.DegToRad(45.0), float32(cfg.Window.W)/float32(cfg.Window.H), 0.1, 10.0),
		hH:         float32(cfg.Window.H) / 2,
		hW:         float32(cfg.Window.W) / 2,
	}
}

func (c *Camera) View() mgl32.Mat4 {
	return c.view
}

func (c *Camera) UpdateView(view mgl32.Mat4) {
	c.view = view
}

func (c *Camera) Projection() mgl32.Mat4 {
	return c.projection
}

func (c *Camera) Model(x, y float32) mgl32.Mat4 {
	m := mgl32.Translate3D(x/c.hW, y/c.hH, 0)

	return m
}
