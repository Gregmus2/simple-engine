package scenes

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Base struct {
	DrawObjects *common.DrawableCollection
	World       *box2d.B2World
	Cfg         *common.Config
	Window      *glfw.Window
	Camera      *graphics.Camera
}

func NewBase(w *box2d.B2World, cfg *common.Config, win *glfw.Window, c *graphics.Camera) Base {
	return Base{
		DrawObjects: common.NewDrawableCollection(),
		World:       w,
		Cfg:         cfg,
		Window:      win,
		Camera:      c,
	}
}

func (b *Base) Init() {

}

func (b *Base) PreUpdate() {
	b.Camera.Update()
}

func (b *Base) Update() {

}

func (b *Base) Drawable() *common.DrawableCollection {
	return b.DrawObjects
}

func (b *Base) Callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch {
	case key == glfw.KeyEscape && action == glfw.Press:
		b.Window.SetShouldClose(true)
	case key == glfw.KeyW && action == glfw.Press:
		b.Camera.Moving(0, -1)
	case key == glfw.KeyS && action == glfw.Press:
		b.Camera.Moving(0, 1)
	case key == glfw.KeyA && action == glfw.Press:
		b.Camera.Moving(1, 0)
	case key == glfw.KeyD && action == glfw.Press:
		b.Camera.Moving(-1, 0)
	case key == glfw.KeyW && action == glfw.Release:
		b.Camera.StopMoving(0, -1)
	case key == glfw.KeyS && action == glfw.Release:
		b.Camera.StopMoving(0, 1)
	case key == glfw.KeyA && action == glfw.Release:
		b.Camera.StopMoving(1, 0)
	case key == glfw.KeyD && action == glfw.Release:
		b.Camera.StopMoving(-1, 0)
	case key == glfw.KeyQ && action == glfw.Press:
		b.Camera.Zoom(0.5)
	case key == glfw.KeyE && action == glfw.Press:
		b.Camera.Zoom(2)
	}
}

func (b *Base) BeginContact(contact box2d.B2ContactInterface) {

}

func (b *Base) EndContact(contact box2d.B2ContactInterface) {

}

func (b *Base) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {

}

func (b *Base) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {

}
