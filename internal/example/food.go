package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/internal/graphics"
	"github.com/Gregmus2/simple-engine/internal/objects"
)

type Food struct {
	objects.Circle
}

func (f *ObjectFactory) NewFood(x, y float64) *Food {
	m := objects.CircleModel{
		X:       x,
		Y:       y,
		Radius:  2,
		T:       box2d.B2BodyType.B2_staticBody,
		Color:   graphics.Yellow(),
		Density: 1.0,
	}
	circle := f.NewCircle(m)
	circle.Fixture.SetUserData([]Tag{FoodTag})

	return &Food{Circle: *circle}
}
