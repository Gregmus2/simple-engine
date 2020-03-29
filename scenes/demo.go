package scenes

import (
	"engine/common"
	"engine/graphics"
	"engine/objects"
	"github.com/ByteArena/box2d"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Demo struct {
	drawable []common.Drawable
	world    *box2d.B2World
	cfg      *common.Config
	factory  *objects.ObjectFactory
	window   *glfw.Window
}

func NewDemo(w *box2d.B2World, cfg *common.Config, f *objects.ObjectFactory, win *glfw.Window) *Demo {
	return &Demo{
		drawable: make([]common.Drawable, 0),
		world:    w,
		cfg:      cfg,
		factory:  f,
		window:   win,
	}
}

func (d *Demo) Init() {
	b := objects.BoxModel{
		X:       0,
		Y:       5,
		H:       10,
		W:       100,
		T:       box2d.B2BodyType.B2_staticBody,
		Color:   *graphics.White(),
		Density: 0,
	}
	box := d.factory.NewBox(b)

	b.X = 5
	b.Y = 50
	b.H = 100
	b.W = 5
	box2 := d.factory.NewBox(b)

	b.X = 95
	b.Y = 50
	box3 := d.factory.NewBox(b)

	c := objects.CircleModel{
		X:       0,
		Y:       0,
		Radius:  1,
		T:       box2d.B2BodyType.B2_dynamicBody,
		Color:   *graphics.White(),
		Density: 1,
	}
	for i := 0; i < 30; i++ {
		c.X = float64(i)*1 + 50
		c.Y = float64(i)*1 + 50
		circle := d.factory.NewCircle(c)
		circle.Fixture.SetFriction(0.2)
		circle.Fixture.SetRestitution(1.0)
		d.drawable = append(d.drawable, circle)
	}
	for i := 0; i < 30; i++ {
		c.X = 50 - float64(i)*1
		c.Y = 50 - float64(i)*1
		circle := d.factory.NewCircle(c)
		circle.Fixture.SetFriction(0.2)
		circle.Fixture.SetRestitution(1.0)
		d.drawable = append(d.drawable, circle)
	}

	d.drawable = append(d.drawable, box)
	d.drawable = append(d.drawable, box2)
	d.drawable = append(d.drawable, box3)
}

func (d *Demo) Update() {

}

func (d *Demo) Drawable() []common.Drawable {
	return d.drawable
}

func (d *Demo) Callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		d.window.SetShouldClose(true)
	}
}
