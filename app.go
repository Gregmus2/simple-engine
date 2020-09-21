package engine

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sirupsen/logrus"
)

type App struct {
	Window *glfw.Window
	World  *box2d.B2World
	GL     *graphics.OpenGL

	cfg           *common.Config
	updateActions []func()
	scale         float32
	scene         common.Scene
	quit          bool
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
		cfg:    cfg,
	}, nil
}

func (app *App) InitWithScene(scene common.Scene) {
	app.scene = scene
	app.World.SetContactListener(scene)
	scene.Init()
	app.initCallbacks()
}

func (app *App) initCallbacks() {
	app.Window.SetKeyCallback(app.scene.Callback)
	app.updateActions = make([]func(), 0)
	app.updateActions = append(app.updateActions, app.scene.PreUpdate)
	if app.cfg.Physics.Enable {
		app.updateActions = append(app.updateActions, func() {
			app.World.Step(timeStep, velocityIterations, positionIterations)
		})
	}
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
		err := d.Draw(app.scale)
		if err != nil {
			logrus.WithError(err).Fatal("draw error")
		}
	}

	glfw.PollEvents()
	app.Window.SwapBuffers()
}
