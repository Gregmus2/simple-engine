package common

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Drawable interface {
	Draw(scale, offsetX, offsetY float32) error
}

type Scene interface {
	Init()
	PreUpdate()
	Update()
	Drawable() *DrawableCollection
	Callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
	MouseCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey)
	ScrollCallback(w *glfw.Window, xoff, yoff float64)
	CursorPositionCallback(w *glfw.Window, xpos, ypos float64)
}

type Init interface {
	OpenGL() error
}
