package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/objects"
)

type Food struct {
	objects.Circle
}

func (f *ObjectFactory) NewFood(x, y float64) *Food {
	m := objects.CircleModel{
		X:       x,
		Y:       y,
		Radius:  2,
		T:       box2d.B2BodyType.B2_dynamicBody,
		Color:   graphics.Yellow(),
		Density: 1.0,
	}
	circle := f.NewCircle(m)

	return &Food{Circle: *circle}
}
