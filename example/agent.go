package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/objects"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/gregmus2/nnga"
	"math"
)

type Agent struct {
	circle *objects.Circle
	prog   uint32
	person *nnga.Person
	cursor *common.Pos

	targetPos box2d.B2Vec2
	distance  float64
}

func (f *ObjectFactory) NewAgent(x, y float64, p *nnga.Person) *Agent {
	m := objects.CircleModel{
		X:       x,
		Y:       y,
		Radius:  5,
		T:       box2d.B2BodyType.B2_dynamicBody,
		Color:   graphics.Blue(),
		Density: 1.0,
	}
	circle := f.NewCircle(m)
	circle.Body.SetFixedRotation(false)
	circle.Body.SetLinearDamping(10.0)
	white := graphics.White()

	angle := circle.Body.GetAngle()
	pos := circle.Body.GetPosition()
	x1, y1 := float32(pos.X)*f.Cfg.Graphics.Scale, float32(pos.Y)*f.Cfg.Graphics.Scale
	x2 := x1 + (circle.Radius * float32(math.Cos(angle)))
	y2 := y1 + (circle.Radius * float32(math.Sin(angle)))

	return &Agent{
		circle: circle,
		prog:   f.Prog.GetByColor(&white),
		person: p,
		cursor: &common.Pos{X: x2, Y: y2},
	}
}

func (a *Agent) Draw(scale float32) error {
	err := a.circle.Draw(scale)
	if err != nil {
		return err
	}

	angle := a.circle.Body.GetAngle()
	pos := a.circle.Body.GetPosition()
	x, y := float32(pos.X)*scale, float32(pos.Y)*scale
	gl.UseProgram(a.prog)
	x2 := x + (a.circle.Radius * float32(math.Cos(angle)))
	y2 := y + (a.circle.Radius * float32(math.Sin(angle)))
	a.circle.Shape.Line(x, y, x2, y2)

	return nil
}
