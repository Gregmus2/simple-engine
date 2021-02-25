package engine

import (
	"github.com/Gregmus2/simple-engine/internal"
	"github.com/Gregmus2/simple-engine/internal/common"
	"github.com/Gregmus2/simple-engine/internal/dispatchers"
	"github.com/Gregmus2/simple-engine/internal/graphics"
	"github.com/Gregmus2/simple-engine/internal/objects"
	"github.com/Gregmus2/simple-engine/internal/physics"
	"github.com/Gregmus2/simple-engine/internal/scenes"
	"github.com/Gregmus2/simple-engine/internal/utils"
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

	if err := c.Provide(graphics.NewPrograms); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewPosToUnitsConverter); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewPercentToPosConverter); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewDrawFunctionsDictionary); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewCamera); err != nil {
		return nil, err
	}

	if err := c.Provide(dispatchers.NewUpdate); err != nil {
		return nil, err
	}

	if err := c.Invoke(utils.NewFPS); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewPositionConverter); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewPositionBuilder); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewDrawer); err != nil {
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

	if err := c.Provide(internal.NewApp); err != nil {
		return nil, err
	}

	if err := c.Provide(scenes.NewBase); err != nil {
		return nil, err
	}

	if err := c.Provide(scenes.NewDemo); err != nil {
		return nil, err
	}

	return c, nil
}
