package engine

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sirupsen/logrus"
	"time"
)

type App struct {
	Window        *glfw.Window
	GL            *graphics.OpenGL
	updateActions []func(dt int64)
	camera        *graphics.Camera
	scene         common.Scene
	quit          bool
}

func NewApp(window *glfw.Window, gl *graphics.OpenGL, c *graphics.Camera, a common.UpdateActionsIn, scene common.Scene) (*App, error) {
	return &App{
		Window:        window,
		GL:            gl,
		camera:        c,
		updateActions: a.Actions,
		scene:         scene,
	}, nil
}

func (app *App) initCallbacks() {
	app.Window.SetKeyCallback(app.scene.Callback)
	app.Window.SetMouseButtonCallback(app.scene.MouseCallback)
	app.Window.SetScrollCallback(app.scene.ScrollCallback)
	app.Window.SetCursorPosCallback(app.scene.CursorPositionCallback)
	app.updateActions = append([]func(dt int64){app.scene.PreUpdate}, app.updateActions...)
	app.updateActions = append(app.updateActions, app.scene.Update)
}

func (app *App) Loop() {
	if app.scene == nil {
		panic("scene isn't set")
	}

	app.scene.Init()
	app.initCallbacks()

	targetDT := int64(1000 / common.Config.Graphics.FPS)
	t := time.Now()
	for !app.Window.ShouldClose() {
		dt := time.Now().Sub(t).Milliseconds()
		if dt < targetDT {
			time.Sleep(time.Millisecond * time.Duration(targetDT-dt))
			dt = time.Now().Sub(t).Milliseconds()
		}
		t = time.Now()
		app.OnUpdate(dt)
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

func (app *App) OnUpdate(dt int64) {
	for _, action := range app.updateActions {
		action(dt)
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
