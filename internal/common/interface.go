package common

import (
	"github.com/ByteArena/box2d"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Scene interface {
	Init()
	PreUpdate(delta float64)
	Update(delta float64)
	GetLights() *LightCollection
	Drawable() *DrawableCollection
	KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
	MouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey)
	MouseMoveCallback(w *glfw.Window, x, y float64)
	ScrollCallback(w *glfw.Window, xOffset, yOffset float64)
	box2d.B2ContactListenerInterface
}

type Init interface {
	OpenGL() error
}

type Shader interface {
	ApplyShader(projection, view, model mgl32.Mat4)
}

type Shape interface {
	Remove()
	Draw()
}

type Drawable interface {
	GetPosition() box2d.B2Vec2
	Shape() Shape
	Shader() Shader
	Die() error
}

type Light interface {
	GetPosition() box2d.B2Vec2
}
