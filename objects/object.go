package objects

import (
	"window/common"
	"window/graphics"
)

type ObjectFactory struct {
	cfg   *common.Config
	prog  *graphics.ProgramFactory
	shape *graphics.ShapeHelper
}

func NewObjectFactory(cfg *common.Config, p *graphics.ProgramFactory, s *graphics.ShapeHelper) *ObjectFactory {
	return &ObjectFactory{cfg: cfg, prog: p, shape: s}
}
