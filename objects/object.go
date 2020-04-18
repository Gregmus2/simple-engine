package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
)

type ObjectFactory struct {
	cfg       *common.Config
	Prog      *graphics.ProgramFactory
	shape     *graphics.ShapeHelper
	world     *box2d.B2World
	converter *graphics.PercentToPosConverter
}

func NewObjectFactory(cfg *common.Config, p *graphics.ProgramFactory, s *graphics.ShapeHelper, w *box2d.B2World, c *graphics.PercentToPosConverter) *ObjectFactory {
	return &ObjectFactory{cfg: cfg, Prog: p, shape: s, world: w, converter: c}
}
