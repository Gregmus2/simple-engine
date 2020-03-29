package objects

import (
	"engine/graphics"
	"github.com/ByteArena/box2d"
	"github.com/go-gl/gl/v4.5-core/gl"
)

type BoxModel struct {
	X, Y, H, W float64
	T          uint8
	Color      graphics.Color
	Density    float64
}

type Box struct {
	Body    *box2d.B2Body
	Fixture *box2d.B2Fixture
	w       float32
	h       float32
	prog    uint32
	shape   *graphics.ShapeHelper
}

func (m *ObjectFactory) NewBox(model BoxModel) *Box {
	model.X, model.Y = m.converter.X(model.X), m.converter.Y(model.Y)
	model.W, model.H = m.converter.X(model.W), m.converter.Y(model.H)

	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = model.T
	bodyDef.FixedRotation = false
	bodyDef.Position = box2d.MakeB2Vec2(model.X/m.cfg.Physics.Scale, model.Y/m.cfg.Physics.Scale)
	body := m.world.CreateBody(&bodyDef)

	shape := box2d.MakeB2PolygonShape()
	shape.SetAsBox(model.W/m.cfg.Physics.Scale/2, model.H/m.cfg.Physics.Scale/2)

	return &Box{
		Body:    body,
		Fixture: body.CreateFixture(&shape, model.Density),
		w:       float32(model.W),
		h:       float32(model.H),
		prog:    m.prog.GetByColor(&model.Color),
		shape:   m.shape,
	}
}

func (o *Box) Draw(scale float32) error {
	pos := o.Body.GetPosition()
	gl.UseProgram(o.prog)
	o.shape.Box(float32(pos.X)*scale-o.w/2, float32(pos.Y)*scale+o.h/2, o.w, o.h)

	return nil
}
