package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/internal/graphics/shaders"
)

type LightBox struct {
	Box
}

func (m *ObjectFactory) NewLightBox(model BoxModel) *LightBox {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = model.T
	bodyDef.FixedRotation = false
	bodyDef.Position = box2d.MakeB2Vec2(model.X/m.Cfg.Physics.Scale, model.Y/m.Cfg.Physics.Scale)
	body := m.World.CreateBody(&bodyDef)

	polygonShape := box2d.MakeB2PolygonShape()
	polygonShape.SetAsBox(model.W/m.Cfg.Physics.Scale/2, model.H/m.Cfg.Physics.Scale/2)

	shape := m.Shape.LightBox(float32(model.W), float32(model.H), float32(model.D))

	return &LightBox{Box{
		Body:    body,
		Fixture: body.CreateFixture(&polygonShape, model.Density),
		shape:   shape,
		shader:  shaders.NewLightShader(m.Prog.Light),
	}}
}
