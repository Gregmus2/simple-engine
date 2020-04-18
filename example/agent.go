package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/objects"
	"github.com/go-gl/gl/v4.5-core/gl"
	"math"
)

type Agent struct {
	circle *objects.Circle
	prog   uint32
}

func (f *ObjectFactory) NewAgent(x, y float64) *Agent {
	m := objects.CircleModel{
		X:       x,
		Y:       y,
		Radius:  5,
		T:       box2d.B2BodyType.B2_dynamicBody,
		Color:   graphics.Blue(),
		Density: 1.0,
	}
	circle := f.factory.NewCircle(m)
	circle.Body.SetFixedRotation(false)
	white := graphics.White()

	return &Agent{circle: circle, prog: f.factory.Prog.GetByColor(&white)}
}

func (a *Agent) Draw(scale float32) error {
	err := a.circle.Draw(scale)
	if err != nil {
		return err
	}

	angle := a.circle.Body.GetAngle()
	degreeAngle := angle * 180 / math.Pi
	pos := a.circle.Body.GetPosition()
	x, y := float32(pos.X)*scale, float32(pos.Y)*scale
	gl.UseProgram(a.prog)
	x2 := x + (a.circle.Radius * float32(math.Cos(degreeAngle*graphics.DoublePI/360.0)))
	y2 := y + (a.circle.Radius * float32(math.Sin(degreeAngle*graphics.DoublePI/360.0)))
	a.circle.Shape.Line(x, y, x2, y2)

	return nil
}
