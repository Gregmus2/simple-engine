package graphics

import (
	"github.com/Gregmus2/simple-engine/internal/common"
	"github.com/go-gl/mathgl/mgl32"
)

type PositionBuilder struct {
	conv *PositionConverter

	view       mgl32.Mat4
	projection mgl32.Mat4
}

func NewPositionBuilder(cfg *common.Config, c *PositionConverter) *PositionBuilder {
	return &PositionBuilder{
		conv: c,

		view:       mgl32.LookAtV(mgl32.Vec3{0, 0, -3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0}),
		projection: mgl32.Perspective(mgl32.DegToRad(45.0), float32(cfg.Window.W)/float32(cfg.Window.H), 0.1, 10.0),
	}
}

func (c *PositionBuilder) View() mgl32.Mat4 {
	return c.view
}

func (c *PositionBuilder) UpdateView(view mgl32.Mat4) {
	c.view = view
}

func (c *PositionBuilder) Projection() mgl32.Mat4 {
	return c.projection
}

func (c *PositionBuilder) Model(x, y float32) mgl32.Mat4 {
	x, y = c.conv.Box2DToOpenGL(x, y)
	m := mgl32.Translate3D(x, y, 0)

	return m
}
