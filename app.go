package engine

import (
	"engine/common"
	"engine/graphics"
	"github.com/ByteArena/box2d"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sirupsen/logrus"
)

type App struct {
	Window *glfw.Window
	World  *box2d.B2World
	GL     *graphics.OpenGL

	scale float32
	scene common.Scene
	quit  bool
}

const velocityIterations = 8
const positionIterations = 2
const timeStep = 1.0 / 60

func NewApp(cfg *common.Config, window *glfw.Window, gl *graphics.OpenGL, world *box2d.B2World) (*App, error) {
	return &App{
		Window: window,
		GL:     gl,
		World:  world,
		scale:  cfg.Graphics.Scale,
	}, nil
}

func (app *App) SetScene(scene common.Scene) {
	app.scene = scene
	scene.Init()
	app.Window.SetKeyCallback(scene.Callback)
}

func (app *App) Loop() {
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
	app.World.Step(timeStep, velocityIterations, positionIterations)
	app.scene.Update()
}

func (app *App) OnRender() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, d := range app.scene.Drawable() {
		err := d.Draw(app.scale)
		if err != nil {
			logrus.WithError(err).Fatal("draw error")
		}
	}

	glfw.PollEvents()
	app.Window.SwapBuffers()
}
