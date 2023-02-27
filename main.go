package engine

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/modules"
	"github.com/Gregmus2/simple-engine/scenes"
	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
	c := dig.New()

	if err := c.Invoke(common.DefineConfig); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewWindow); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewOpenGL); err != nil {
		return nil, err
	}

	if err := c.Provide(graphics.NewCamera); err != nil {
		return nil, err
	}

	if err := c.Provide(modules.NewFPS); err != nil {
		return nil, err
	}

	if err := c.Provide(NewApp); err != nil {
		return nil, err
	}

	if err := c.Provide(scenes.NewBase); err != nil {
		return nil, err
	}

	if err := c.Provide(NewInit); err != nil {
		return nil, err
	}

	if err := c.Invoke(graphics.DefinePrograms); err != nil {
		return nil, err
	}

	if err := c.Invoke(graphics.LoadTextures); err != nil {
		return nil, err
	}

	return c, nil
}
