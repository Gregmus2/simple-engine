package objects

import (
	"engine/graphics"
	"github.com/ByteArena/box2d"
	"github.com/go-gl/gl/v4.5-core/gl"
)

type CircleModel struct {
	X, Y    float64
	Radius  float32
	T       uint8
	Color   graphics.Color
	Density float64
}

type Circle struct {
	Radius  float32
	Body    *box2d.B2Body
	Fixture *box2d.B2Fixture
	prog    uint32
	shape   *graphics.ShapeHelper
}

func (m *ObjectFactory) NewCircle(model CircleModel) *Circle {
	model.X, model.Y, model.Radius = m.converter.X(model.X), m.converter.Y(model.Y), m.converter.Radius(model.Radius)

	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Position = box2d.MakeB2Vec2(model.X/m.cfg.Physics.Scale, model.Y/m.cfg.Physics.Scale)
	bodyDef.Type = model.T
	body := m.world.CreateBody(&bodyDef)

	shape := box2d.MakeB2CircleShape()
	shape.SetRadius(float64(model.Radius) / m.cfg.Physics.Scale)

	return &Circle{
		Radius:  model.Radius,
		Body:    body,
		Fixture: body.CreateFixture(&shape, model.Density),
		prog:    m.prog.GetByColor(&model.Color),
		shape:   m.shape,
	}
}

func (o *Circle) Draw(scale float32) error {
	pos := o.Body.GetPosition()
	gl.UseProgram(o.prog)
	o.shape.Circle(float32(pos.X)*scale, float32(pos.Y)*scale, o.Radius)

	return nil
}
