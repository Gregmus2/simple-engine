package main

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"runtime"
	"window/common"
	"window/graphics"
	"window/objects"
	"window/physics"
	"window/scenes"
)

func main() {
	runtime.LockOSThread()

	c, err := buildContainer()
	if err != nil {
		logrus.WithError(err).Fatal("error building DI container")
	}

	if err := c.Invoke(func(app *App, demo *scenes.Demo) {
		app.SetScene(demo)
		app.Loop()
	}); err != nil {
		logrus.Fatal(err)
	}
}

func buildContainer() (*dig.Container, error) {
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
