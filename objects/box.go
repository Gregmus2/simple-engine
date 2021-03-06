package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type BoxModel struct {
	X, Y, W, H float64
	T          uint8
	Color      graphics.Color
	Density    float64
}

type Box struct {
	Body    *box2d.B2Body
	Fixture *box2d.B2Fixture
	w       float32
	h       float32
	prog    *graphics.Program
	shape   *graphics.ShapeHelper
	color   graphics.Color
}

func (m *ObjectFactory) NewBox(model BoxModel) *Box {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = model.T
	bodyDef.FixedRotation = false
	bodyDef.Position = box2d.MakeB2Vec2(model.X/m.Cfg.Physics.Scale, model.Y/m.Cfg.Physics.Scale)
	body := m.World.CreateBody(&bodyDef)

	shape := box2d.MakeB2PolygonShape()
	shape.SetAsBox(model.W/m.Cfg.Physics.Scale/2, model.H/m.Cfg.Physics.Scale/2)

	return &Box{
		Body:    body,
		Fixture: body.CreateFixture(&shape, model.Density),
		w:       float32(model.W),
		h:       float32(model.H),
		prog:    m.Prog,
		shape:   m.Shape,
		color:   model.Color,
	}
}

func (o *Box) Draw(scale float32) error {
	pos := o.Body.GetPosition()
	o.prog.ApplyProgram(o.color)
	o.shape.Box(float32(pos.X)*scale-o.w/2, float32(pos.Y)*scale+o.h/2, o.w, o.h)
	gl.UseProgram(0)

	return nil
}
