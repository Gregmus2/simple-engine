package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
)

type ObjectFactory struct {
	Cfg       *common.Config
	Prog      *graphics.ProgramFactory
	Shape     *graphics.ShapeHelper
	World     *box2d.B2World
	Converter *graphics.PercentToPosConverter
}

func NewObjectFactory(cfg *common.Config, p *graphics.ProgramFactory, s *graphics.ShapeHelper, w *box2d.B2World, c *graphics.PercentToPosConverter) *ObjectFactory {
	return &ObjectFactory{Cfg: cfg, Prog: p, Shape: s, World: w, Converter: c}
}
