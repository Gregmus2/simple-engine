package scenes

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Base struct {
	DrawObjects *common.DrawableCollection
	Window      *glfw.Window
	Camera      *graphics.Camera
}

func NewBase(win *glfw.Window, c *graphics.Camera) Base {
	return Base{
		DrawObjects: common.NewDrawableCollection(),
		Window:      win,
		Camera:      c,
	}
}

func (b *Base) Init() {

}

func (b *Base) PreUpdate(_ int64) {

}

func (b *Base) Update(_ int64) {

}

func (b *Base) Drawable() *common.DrawableCollection {
	return b.DrawObjects
}

func (b *Base) Callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch {
	case key == glfw.KeyEscape && action == glfw.Press:
		b.Window.SetShouldClose(true)
	}
}

func (b *Base) MouseCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	b.Camera.MouseCallback(button, action)
}

func (b *Base) ScrollCallback(w *glfw.Window, xoff, yoff float64) {
	b.Camera.ScrollCallback(yoff)
}

func (b *Base) CursorPositionCallback(w *glfw.Window, xpos, ypos float64) {
	b.Camera.CursorPositionCallback(xpos, ypos)
}
