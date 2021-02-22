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
	PreUpdate()
	Update()
	Drawable() *DrawableCollection
	Callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
	box2d.B2ContactListenerInterface
}

type Init interface {
	OpenGL() error
}

type Camera interface {
	View() *mgl32.Mat4
	Projection() *mgl32.Mat4
	Model(x, y float32) *mgl32.Mat4
}

type Shader interface {
	ApplyShader(projection, view, model *mgl32.Mat4)
}

type Shape interface {
	Remove()
	Draw()
}
