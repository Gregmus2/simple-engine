package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/objects"
)

type Ground struct {
	objects.Box
}

func (f *ObjectFactory) NewGround(x, y, h, w float64, color graphics.Color) {
	b := objects.BoxModel{
		X:       x,
		Y:       y,
		H:       h,
		W:       w,
		T:       box2d.B2BodyType.B2_staticBody,
		Color:   color,
		Density: 0,
	}
	f.factory.NewBox(b)
}
