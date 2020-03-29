package engine

import (
	"engine/common"
	"engine/graphics"
	"engine/objects"
	"engine/physics"
	"engine/scenes"
	"go.uber.org/dig"
)

func BuildContainer(cfgFile string) (*dig.Container, error) {
	c := dig.New()

	if err := c.Provide(common.NewConfig(cfgFile)); err != nil {
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
