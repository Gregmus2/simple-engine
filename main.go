package engine

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/objects"
	"github.com/Gregmus2/simple-engine/physics"
	"github.com/Gregmus2/simple-engine/scenes"
	"go.uber.org/dig"
)

func BuildContainer(cfgFile string) (*dig.Container, error) {
	c := dig.New()

	if err := c.Provide(common.NewConfig); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewWindow); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewProgramFactory); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewPosToUnitsConverter); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewPercentToPosConverter); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewShapeFactory); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewOpenGL); err != nil {
		return nil, err
	}

	if err := c.Provide(physics.NewWorld); err != nil {
		return nil, err
	}

	if err := c.Provide(objects.NewObjectFactory); err != nil {
		return nil, err
	}

	if err := c.Provide(NewApp); err != nil {
		return nil, err
	}

	if err := c.Provide(scenes.NewDemo); err != nil {
		return nil, err
	}

	return c, nil
}
