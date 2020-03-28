package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/go-gl/gl/v4.5-core/gl"
	"window/graphics"
)

type Circle struct {
	Radius float32
	Body   *box2d.B2Body
	prog   uint32
	shape  *graphics.ShapeHelper
}

func (m *ObjectFactory) NewCircle(w *box2d.B2World, x, y float64, radius float32, color *graphics.Color) *Circle {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Position = box2d.MakeB2Vec2(x/m.cfg.Physics.Scale, y/m.cfg.Physics.Scale)
	bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	body := w.CreateBody(&bodyDef)
	shape := box2d.MakeB2CircleShape()
	shape.SetRadius(float64(radius) / m.cfg.Physics.Scale)
	fixture := body.CreateFixture(&shape, 1.0)
	fixture.SetFriction(0.2)
	fixture.SetRestitution(1.0)

	return &Circle{
		Radius: radius,
		Body:   body,
		prog:   m.prog.GetByColor(color),
		shape:  m.shape,
	}
}

func (o *Circle) Draw(scale float32) error {
	pos := o.Body.GetPosition()
	gl.UseProgram(o.prog)
	o.shape.Circle(float32(pos.X)*scale, float32(pos.Y)*scale, o.Radius)

	return nil
}
