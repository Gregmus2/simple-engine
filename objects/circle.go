package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/graphics"
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
	Shape   *graphics.ShapeHelper
}

func (m *ObjectFactory) NewCircle(model CircleModel) *Circle {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Position = box2d.MakeB2Vec2(model.X/m.Cfg.Physics.Scale, model.Y/m.Cfg.Physics.Scale)
	bodyDef.Type = model.T
	body := m.World.CreateBody(&bodyDef)

	shape := box2d.MakeB2CircleShape()
	shape.SetRadius(float64(model.Radius) / m.Cfg.Physics.Scale)

	return &Circle{
		Radius:  model.Radius,
		Body:    body,
		Fixture: body.CreateFixture(&shape, model.Density),
		prog:    m.Prog.GetByColor(&model.Color),
		Shape:   m.Shape,
	}
}

func (o *Circle) Draw(scale float32) error {
	pos := o.Body.GetPosition()
	gl.UseProgram(o.prog)
	o.Shape.Circle(float32(pos.X)*scale, float32(pos.Y)*scale, o.Radius)

	return nil
}
