package scenes

import (
	"github.com/ByteArena/box2d"
	"window/common"
	"window/graphics"
	"window/objects"
)

type Demo struct {
	drawable []common.Drawable
	world    *box2d.B2World
	cfg      *common.Config
	factory  *objects.ObjectFactory
}

func NewDemo(w *box2d.B2World, cfg *common.Config, f *objects.ObjectFactory) *Demo {
	return &Demo{
		drawable: make([]common.Drawable, 0),
		world:    w,
		cfg:      cfg,
		factory:  f,
	}
}

func (d *Demo) Init() {
	box := d.factory.NewBox(d.world, float64(d.cfg.Window.W/2), 20, 20, float32(d.cfg.Window.W), box2d.B2BodyType.B2_staticBody, graphics.Green())
	box2 := d.factory.NewBox(d.world, 5, float64(d.cfg.Window.Center.Y), float32(d.cfg.Window.H), 10, box2d.B2BodyType.B2_staticBody, graphics.White())
	box3 := d.factory.NewBox(d.world, float64(d.cfg.Window.W-5), float64(d.cfg.Window.Center.Y), float32(d.cfg.Window.H), 10, box2d.B2BodyType.B2_staticBody, graphics.White())
	for i := 0; i < 30; i++ {
		circle := d.factory.NewCircle(d.world, float64(int(d.cfg.Window.Center.X)-i*10), float64(int(d.cfg.Window.Center.Y)-i*10), 10, graphics.White())
		d.drawable = append(d.drawable, circle)
	}
	for i := 0; i < 30; i++ {
		circle := d.factory.NewCircle(d.world, float64(int(d.cfg.Window.Center.X)+i*10), float64(int(d.cfg.Window.Center.Y)-i*10), 10, graphics.White())
		d.drawable = append(d.drawable, circle)
	}
	//circle := d.factory.NewCircle(d.world, 0, 0, 10, graphics.White())
	//d.drawable = append(d.drawable, circle)
	d.drawable = append(d.drawable, box)
	d.drawable = append(d.drawable, box2)
	d.drawable = append(d.drawable, box3)
}

func (d *Demo) Update() {

}

func (d *Demo) Drawable() []common.Drawable {
	return d.drawable
}
