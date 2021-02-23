package scenes

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Base struct {
	DrawObjects  *common.DrawableCollection
	World        *box2d.B2World
	Cfg          *common.Config
	Window       *glfw.Window
	Camera       common.Camera
	mouseControl common.MouseController
}

func NewBase(w *box2d.B2World, cfg *common.Config, win *glfw.Window, c common.Camera, m common.MouseController) Base {
	return Base{
		DrawObjects:  common.NewDrawableCollection(),
		World:        w,
		Cfg:          cfg,
		Window:       win,
		Camera:       c,
		mouseControl: m,
	}
}

func (b *Base) Init() {

}

func (b *Base) PreUpdate(delta float64) {

}

func (b *Base) Update(delta float64) {
	b.mouseControl.Update(delta)
}

func (b *Base) Drawable() *common.DrawableCollection {
	return b.DrawObjects
}

func (b *Base) KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		b.Window.SetShouldClose(true)
	}
}

func (b *Base) MouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	b.mouseControl.MouseButtonCallback(w, button, action, mods)
}

func (b *Base) MouseMoveCallback(w *glfw.Window, x, y float64) {
	b.mouseControl.MouseMoveCallback(w, x, y)
}

func (b *Base) BeginContact(contact box2d.B2ContactInterface) {

}

func (b *Base) EndContact(contact box2d.B2ContactInterface) {

}

func (b *Base) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {

}

func (b *Base) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {

}
