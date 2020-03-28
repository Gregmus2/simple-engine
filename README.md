### Usage

```go
func main() {
	runtime.LockOSThread()

	c, err := BuildContainer()
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
```