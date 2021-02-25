package objects

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/internal"
	"github.com/Gregmus2/simple-engine/internal/common"
	"github.com/Gregmus2/simple-engine/internal/graphics"
	"go.uber.org/dig"
)

type ObjectFactory struct {
	Cfg               *common.Config
	Prog              *graphics.Programs
	Shape             *graphics.ShapeFactory
	World             *box2d.B2World
	Converter         *graphics.PercentToPosConverter
	Camera            *graphics.Camera
	App               *internal.App
	PositionConverter *graphics.PositionConverter
}

type Params struct {
	dig.In

	Cfg               *common.Config
	Prog              *graphics.Programs
	Shape             *graphics.ShapeFactory
	World             *box2d.B2World
	Converter         *graphics.PercentToPosConverter
	Camera            *graphics.Camera
	App               *internal.App
	PositionConverter *graphics.PositionConverter
}

func NewObjectFactory(params Params) *ObjectFactory {
	return &ObjectFactory{
		Cfg:               params.Cfg,
		Prog:              params.Prog,
		Shape:             params.Shape,
		World:             params.World,
		Converter:         params.Converter,
		Camera:            params.Camera,
		App:               params.App,
		PositionConverter: params.PositionConverter,
	}
}
