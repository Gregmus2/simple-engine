package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
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
	prog    *graphics.Program
	Shape   *graphics.ShapeHelper
	color   graphics.Color
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
		prog:    m.Prog,
		Shape:   m.Shape,
		color:   model.Color,
	}
}

func (o *Circle) Draw(scale, offsetX, offsetY float32) error {
	pos := o.Body.GetPosition()
	o.prog.ApplyProgram(o.color)
	o.Shape.Circle((float32(pos.X)+offsetX)*scale, (float32(pos.Y)+offsetY)*scale, o.Radius)
	gl.UseProgram(0)

	return nil
}
