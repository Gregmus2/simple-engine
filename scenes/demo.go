package scenes

import (
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/objects"
)

type Demo struct {
	Base
	factory *objects.ObjectFactory
	con     *graphics.PercentToPosConverter
}

func NewDemo(base *Base, f *objects.ObjectFactory, con *graphics.PercentToPosConverter) *Demo {
	return &Demo{
		Base:    *base,
		factory: f,
		con:     con,
	}
}

func (d *Demo) Init() {
	//b := objects.BoxModel{
	//	X:       d.con.X(50),
	//	Y:       d.con.Y(5),
	//	W:       d.con.X(100),
	//	H:       d.con.Y(1),
	//	T:       box2d.B2BodyType.B2_staticBody,
	//	Color:   graphics.White(),
	//	Density: 0,
	//}
	//box := d.factory.NewBox(b)
	//
	//b.X = d.con.X(5)
	//b.Y = d.con.Y(50)
	//b.H = d.con.Y(100)
	//b.W = d.con.X(1)
	//box2 := d.factory.NewBox(b)
	//
	//b.X = d.con.X(95)
	//b.Y = d.con.Y(50)
	//box3 := d.factory.NewBox(b)
	//
	//c := objects.CircleModel{
	//	X:       0,
	//	Y:       0,
	//	Radius:  d.con.Radius(1),
	//	T:       box2d.B2BodyType.B2_dynamicBody,
	//	Color:   graphics.White(),
	//	Density: 1,
	//}
	//for i := 0; i < 30; i++ {
	//	c.X = float64(i) + d.con.X(50)
	//	c.Y = float64(i) + d.con.Y(50)
	//	circle := d.factory.NewCircle(c)
	//	circle.Fixture.SetFriction(0.2)
	//	circle.Fixture.SetRestitution(1.0)
	//	d.DrawObjects.Put(circle)
	//}
	//for i := 0; i < 30; i++ {
	//	c.X = d.con.X(50) - float64(i)
	//	c.Y = d.con.Y(50) - float64(i)
	//	circle := d.factory.NewCircle(c)
	//	circle.Fixture.SetFriction(0.2)
	//	circle.Fixture.SetRestitution(1.0)
	//	d.DrawObjects.Put(circle)
	//}
	//
	//d.DrawObjects.Put(box)
	//d.DrawObjects.Put(box2)
	//d.DrawObjects.Put(box3)
}
