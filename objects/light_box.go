package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type LightBox struct {
	Body       *box2d.B2Body
	Fixture    *box2d.B2Fixture
	w, h, d, z float32
	vbo, vao   *uint32
	prog       uint32
	shape      *graphics.ShapeHelper
	color      graphics.Color
}

func (m *ObjectFactory) NewLightBox(model BoxModel) *LightBox {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = model.T
	bodyDef.FixedRotation = false
	bodyDef.Position = box2d.MakeB2Vec2(model.X/m.Cfg.Physics.Scale, model.Y/m.Cfg.Physics.Scale)
	body := m.World.CreateBody(&bodyDef)

	shape := box2d.MakeB2PolygonShape()
	shape.SetAsBox(model.W/m.Cfg.Physics.Scale/2, model.H/m.Cfg.Physics.Scale/2)

	program := m.Prog.Light
	vbo, vao := m.Shape.LightBox(float32(model.W), float32(model.H), float32(model.D))

	return &LightBox{
		Body:    body,
		Fixture: body.CreateFixture(&shape, model.Density),
		w:       float32(model.W),
		h:       float32(model.H),
		d:       float32(model.D),
		z:       float32(model.Z),
		vbo:     vbo,
		vao:     vao,
		prog:    program,
		shape:   m.Shape,
		color:   model.Color,
	}
}

func (o *LightBox) Draw(scale float32) error {
	pos := o.Body.GetPosition()
	gl.UseProgram(o.prog)

	// near -> 3
	//projection := mgl32.Ortho(-1, 1, -1, 1, 0.1, 100)
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(1200)/float32(600), 0.1, 10.0)
	projectionUniform := gl.GetUniformLocation(o.prog, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	camera := mgl32.LookAtV(mgl32.Vec3{1, 1, 1}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cameraUniform := gl.GetUniformLocation(o.prog, gl.Str("view\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	x, y := float32(pos.X)*scale, float32(pos.Y)*scale
	println(x/600, y/300)
	model := mgl32.Translate3D(x/600, y/300, o.z).Mul4(mgl32.Scale3D(0.2, 0.2, 0.2))
	//model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(o.prog, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	gl.BindVertexArray(*o.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)

	//gl.UseProgram(0)

	return nil
}

func (o *LightBox) Die() error {
	graphics.ClearBuffers(o.vbo, o.vao)

	return nil
}
