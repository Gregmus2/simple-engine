package common

import (
	"github.com/ByteArena/box2d"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Drawable interface {
	GetPosition() box2d.B2Vec2
	Shape() Shape
	Shader() Shader
	Die() error
}

type Scene interface {
	Init()
	PreUpdate(delta float64)
	Update(delta float64)
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

type Camera interface {
	View() mgl32.Mat4
	UpdateView(view mgl32.Mat4)
	Projection() mgl32.Mat4
	Model(x, y float32) mgl32.Mat4
}

type Shader interface {
	ApplyShader(projection, view, model mgl32.Mat4)
}

type Shape interface {
	Remove()
	Draw()
}

type CameraControl interface {
	MouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action)
	MouseMove(x, y float64)
	Scroll(yOffset float64)
	Key(key glfw.Key, action glfw.Action)
	Update(delta float64)
}
