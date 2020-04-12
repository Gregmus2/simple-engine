package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
)

type ObjectFactory struct {
	cfg   *common.Config
	prog  *graphics.ProgramFactory
	shape *graphics.ShapeHelper
	world *box2d.B2World
}

func NewObjectFactory(cfg *common.Config, p *graphics.ProgramFactory, s *graphics.ShapeHelper, w *box2d.B2World) *ObjectFactory {
	return &ObjectFactory{cfg: cfg, prog: p, shape: s, world: w}
}
