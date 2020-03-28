package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/go-gl/gl/v4.5-core/gl"
	"window/graphics"
)

type Box struct {
	Body  *box2d.B2Body
	w     float32
	h     float32
	prog  uint32
	shape *graphics.ShapeHelper
}

func (m *ObjectFactory) NewBox(world *box2d.B2World, x, y float64, h, w float32, t uint8, color *graphics.Color) *Box {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = t
	bodyDef.FixedRotation = false
	bodyDef.Position = box2d.MakeB2Vec2(x/m.cfg.Physics.Scale, y/m.cfg.Physics.Scale)
	body := world.CreateBody(&bodyDef)

	shape := box2d.MakeB2PolygonShape()
	shape.SetAsBox(float64(w)/m.cfg.Physics.Scale/2, float64(h)/m.cfg.Physics.Scale/2)
	d := 0.0
	if t == box2d.B2BodyType.B2_dynamicBody {
		d = 1
	}
	body.CreateFixture(&shape, d)

	return &Box{
		Body:  body,
		w:     w,
		h:     h,
		prog:  m.prog.GetByColor(color),
		shape: m.shape,
	}
}

func (o *Box) Draw(scale float32) error {
	pos := o.Body.GetPosition()
	gl.UseProgram(o.prog)
	o.shape.Box(float32(pos.X)*scale-o.w/2, float32(pos.Y)*scale+o.h/2, o.w, o.h)

	return nil
}
