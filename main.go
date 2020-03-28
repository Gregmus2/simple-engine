package engine

import (
	"engine/common"
	"engine/graphics"
	"engine/objects"
	"engine/physics"
	"engine/scenes"
	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
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

	if err := c.Provide(graphics.NewPositionHelper); err != nil {
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
