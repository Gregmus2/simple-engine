package objects

import (
	"engine/common"
	"engine/graphics"
	"github.com/ByteArena/box2d"
)

type ObjectFactory struct {
	cfg       *common.Config
	prog      *graphics.ProgramFactory
	shape     *graphics.ShapeHelper
	world     *box2d.B2World
	converter *graphics.PercentToPosConverter
}

func NewObjectFactory(cfg *common.Config, p *graphics.ProgramFactory, s *graphics.ShapeHelper, w *box2d.B2World, c *graphics.PercentToPosConverter) *ObjectFactory {
	return &ObjectFactory{cfg: cfg, prog: p, shape: s, world: w, converter: c}
}
