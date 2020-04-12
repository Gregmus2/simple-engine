package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.5-core/gl"
)

type BoxModel struct {
	X, Y    float64
	H, W    float32
	T       uint8
	Color   graphics.Color
	Density float64
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
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = model.T
	bodyDef.FixedRotation = false
	bodyDef.Position = box2d.MakeB2Vec2(model.X/m.cfg.Physics.Scale, model.Y/m.cfg.Physics.Scale)
	body := m.world.CreateBody(&bodyDef)

	shape := box2d.MakeB2PolygonShape()
	shape.SetAsBox(float64(model.W)/m.cfg.Physics.Scale/2, float64(model.H)/m.cfg.Physics.Scale/2)

	return &Box{
		Body:    body,
		Fixture: body.CreateFixture(&shape, model.Density),
		w:       model.W,
		h:       model.H,
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
