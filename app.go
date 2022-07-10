package engine

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sirupsen/logrus"
)

type App struct {
	Window        *glfw.Window
	GL            *graphics.OpenGL
	cfg           *common.Config
	updateActions []func()
	camera        *graphics.Camera
	scene         common.Scene
	quit          bool
}

func NewApp(cfg *common.Config, window *glfw.Window, gl *graphics.OpenGL, c *graphics.Camera, a common.UpdateActionsIn) (*App, error) {
	return &App{
		Window:        window,
		GL:            gl,
		camera:        c,
		cfg:           cfg,
		updateActions: a.Actions,
	}, nil
}

func (app *App) InitWithScene(scene common.Scene) {
	app.scene = scene
	scene.Init()
	app.initCallbacks()
}

func (app *App) initCallbacks() {
	app.Window.SetKeyCallback(app.scene.Callback)
	app.Window.SetMouseButtonCallback(app.scene.MouseCallback)
	app.Window.SetScrollCallback(app.scene.ScrollCallback)
	app.Window.SetCursorPosCallback(app.scene.CursorPositionCallback)
	app.updateActions = append([]func(){app.scene.PreUpdate}, app.updateActions...)
	app.updateActions = append(app.updateActions, app.scene.Update)
}

func (app *App) Loop() {
	if app.scene == nil {
		panic("scene isn't set")
	}

	for !app.Window.ShouldClose() {
		app.OnUpdate()
		app.OnRender()

		if app.quit {
			break
		}
	}

	app.Destroy()
}

func (app *App) Destroy() {
	app.Window.Destroy()
	glfw.Terminate()
}

func (app *App) OnUpdate() {
	for _, action := range app.updateActions {
		action()
	}
}

func (app *App) OnRender() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for d := range app.scene.Drawable().Elements {
		x, y := app.camera.Position()
		err := d.Draw(app.camera.Scale(), x, y)
		if err != nil {
			logrus.WithError(err).Fatal("draw error")
		}
	}

	glfw.PollEvents()
	app.Window.SwapBuffers()
}

func (app *App) Shutdown() {
	app.quit = true
}
