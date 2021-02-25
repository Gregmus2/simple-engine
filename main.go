package engine

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/controls"
	"github.com/Gregmus2/simple-engine/dispatchers"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/objects"
	"github.com/Gregmus2/simple-engine/physics"
	"github.com/Gregmus2/simple-engine/scenes"
	"github.com/Gregmus2/simple-engine/utils"
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

	if err := c.Provide(controls.NewMouseControl); err != nil {
		return nil, err
	}

	if err := c.Provide(dispatchers.NewUpdate); err != nil {
		return nil, err
	}

	if err := c.Invoke(utils.NewFPS); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewCamera); err != nil {
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

	if err := c.Provide(NewApp); err != nil {
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
