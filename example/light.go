package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/objects"
	"github.com/Gregmus2/simple-engine/scenes"
)

type Light struct {
	scenes.Base
	factory *objects.ObjectFactory
	con     *graphics.PercentToPosConverter
}

func NewLight(base scenes.Base, f *objects.ObjectFactory, con *graphics.PercentToPosConverter) *Light {
	return &Light{
		Base:    base,
		factory: f,
		con:     con,
	}
}

func (l *Light) Init() {
	floor := objects.BoxModel{
		X:       l.con.X(50),
		Y:       l.con.Y(0),
		W:       l.con.W(100),
		H:       1,
		D:       l.con.D(70),
		T:       box2d.B2BodyType.B2_staticBody,
		Color:   graphics.White(),
		Density: 0,
	}
	l.DrawObjects.Put(l.factory.NewBox(floor))

	light := l.factory.NewLightBox(objects.BoxModel{
		X:       l.con.X(60),
		Y:       l.con.Y(60),
		Z:       -0.2,
		W:       l.con.W(5),
		H:       l.con.H(5),
		D:       l.con.D(5),
		T:       box2d.B2BodyType.B2_staticBody,
		Color:   graphics.White(),
		Density: 0,
	})
	l.DrawObjects.Put(light)

	box := l.factory.NewBox(objects.BoxModel{
		X:       l.con.X(50),
		Y:       l.con.Y(60),
		W:       l.con.W(5),
		H:       l.con.H(5),
		D:       l.con.D(5),
		T:       box2d.B2BodyType.B2_dynamicBody,
		Color:   graphics.Blue(),
		Density: 0,
	})
	l.DrawObjects.Put(box)

	//box := l.factory.NewBox(objects.BoxModel{
	//	X:       l.con.X(50),
	//	Y:       l.con.Y(50),
	//	W:       l.con.X(30),
	//	H:       l.con.Y(30),
	//	D:       10,
	//	T:       box2d.B2BodyType.B2_staticBody,
	//	Color:   graphics.Blue(),
	//	Density: 0,
	//})
	//l.DrawObjects.Put(box)
}
