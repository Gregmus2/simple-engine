package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/internal/common"
	"github.com/Gregmus2/simple-engine/internal/graphics"
	"github.com/Gregmus2/simple-engine/internal/graphics/shaders"
)

type BoxModel struct {
	X, Y, Z, W, H, D float64
	T                uint8
	Color            graphics.Color
	Density          float64
}

type Box struct {
	Body    *box2d.B2Body
	Fixture *box2d.B2Fixture
	shape   common.Shape
	shader  common.Shader
}

func (m *ObjectFactory) NewBox(model BoxModel) *Box {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = model.T
	bodyDef.FixedRotation = false
	bodyDef.Position = box2d.MakeB2Vec2(model.X/m.Cfg.Physics.Scale, model.Y/m.Cfg.Physics.Scale)
	body := m.World.CreateBody(&bodyDef)

	polygonShape := box2d.MakeB2PolygonShape()
	polygonShape.SetAsBox(model.W/m.Cfg.Physics.Scale/2, model.H/m.Cfg.Physics.Scale/2)

	shape := m.Shape.Box(float32(model.W), float32(model.H), float32(model.D))

	return &Box{
		Body:    body,
		Fixture: body.CreateFixture(&polygonShape, model.Density),
		shape:   shape,
		shader:  shaders.NewObjectShader(m.Prog.SimpleColor, model.Color, m.Camera, m.App, m.PositionConverter),
	}
}

func (b *Box) GetPosition() box2d.B2Vec2 {
	return b.Body.GetPosition()
}

func (b *Box) Shape() common.Shape {
	return b.shape
}

func (b *Box) Shader() common.Shader {
	return b.shader
}

func (b *Box) Die() error {
	b.shape.Remove()

	return nil
}
