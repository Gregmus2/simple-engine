Engine to create small simulations.

It is providing 
* Graphics (opengl-v4.6 + glfw-v3.3) https://github.com/go-gl/gl
* Physics (box2d) https://github.com/ByteArena/box2d

![](example.gif)

### Requirements

```build-essential libgl1-mesa-dev xorg-dev```

### Usage

Run only with `go run .` It won't work correctly with IDE runner

Main file example:
```go
func main() {
	runtime.LockOSThread()

	c, err := engine.BuildContainer()
	if err != nil {
		logrus.WithError(err).Fatal("error building DI container")
	}

	err = buildContainer(c)
	if err != nil {
		logrus.WithError(err).Fatal("error building DI container")
	}

	if err := c.Invoke(func(app *engine.App, demo *scenes.Demo) {
		app.InitWithScene(demo)
		app.Loop()
	}); err != nil {
		logrus.Fatal(err)
	}
}

func buildContainer(c *dig.Container) error {
	if err := c.Provide(NewBar); err != nil {
		return err
	}

	if err := c.Provide(NewFoo); err != nil {
		return err
	}

	return nil
}

```
* With your buildContainer method you can provide additional structures into your application
  (actually into dig service container (go.uber.org/dig))
* You should implement Scene interface to replace scenes.Demo scene.
You can use `scenes.Base` as embedded structure
* You can use `objects.ObjectFactory` structure as embedded in your factory to create new objects.
* Also you need to implement `common.Init` interface, at least with empty content

You should look into `example` directory to figure it out better 