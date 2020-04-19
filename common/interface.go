package common

import (
	"github.com/ByteArena/box2d"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Drawable interface {
	Draw(scale float32) error
}

type Scene interface {
	Init()
	PreUpdate()
	Update()
	Drawable() []Drawable
	Callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
	box2d.B2ContactListenerInterface
}
