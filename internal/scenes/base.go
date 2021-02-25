package scenes

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/internal/common"
	"github.com/Gregmus2/simple-engine/internal/graphics"
	"github.com/go-gl/glfw/v3.3/glfw"
	"go.uber.org/dig"
)

type Base struct {
	DrawObjects     *common.DrawableCollection
	LightObjects    *common.LightCollection
	World           *box2d.B2World
	Cfg             *common.Config
	Window          *glfw.Window
	PositionBuilder *graphics.PositionBuilder
	camera          *graphics.Camera
}

type Params struct {
	dig.In

	World           *box2d.B2World
	Cfg             *common.Config
	Window          *glfw.Window
	PositionBuilder *graphics.PositionBuilder
	Camera          *graphics.Camera
}

func NewBase(params Params) *Base {
	return &Base{
		DrawObjects:     common.NewDrawableCollection(),
		LightObjects:    common.NewLightCollection(),
		World:           params.World,
		Cfg:             params.Cfg,
		Window:          params.Window,
		PositionBuilder: params.PositionBuilder,
		camera:          params.Camera,
	}
}

func (b *Base) Init() {

}

func (b *Base) PreUpdate(delta float64) {

}

func (b *Base) Update(delta float64) {
	b.camera.Update(delta)
}

func (b *Base) Drawable() *common.DrawableCollection {
	return b.DrawObjects
}

func (b *Base) GetLights() *common.LightCollection {
	return b.LightObjects
}

func (b *Base) KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		b.Window.SetShouldClose(true)
	}
	b.camera.Key(key, action)
}

func (b *Base) MouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	b.camera.MouseButton(w, button, action)
}

func (b *Base) MouseMoveCallback(w *glfw.Window, x, y float64) {
	b.camera.MouseMove(x, y)
}

func (b *Base) ScrollCallback(w *glfw.Window, xOffset, yOffset float64) {
	b.camera.Scroll(yOffset)
}

func (b *Base) BeginContact(contact box2d.B2ContactInterface) {

}

func (b *Base) EndContact(contact box2d.B2ContactInterface) {

}

func (b *Base) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {

}

func (b *Base) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {

}
