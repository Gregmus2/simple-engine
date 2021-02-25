package engine

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type App struct {
	Window *glfw.Window
	World  *box2d.B2World
	GL     *graphics.OpenGL

	cfg           *common.Config
	drawer        *graphics.Drawer
	updateActions []func(delta float64)
	scale         float32
	scene         common.Scene
	quit          bool
}

const velocityIterations = 8
const positionIterations = 2
const timeStep = 1.0 / 40

func NewApp(cfg *common.Config, window *glfw.Window, gl *graphics.OpenGL, world *box2d.B2World, d *graphics.Drawer) (*App, error) {
	return &App{
		Window: window,
		GL:     gl,
		World:  world,
		scale:  cfg.Graphics.Scale,
		cfg:    cfg,
		drawer: d,
	}, nil
}

func (app *App) InitWithScene(scene common.Scene) {
	app.scene = scene
	app.World.SetContactListener(scene)
	scene.Init()
	app.initCallbacks()
}

func (app *App) initCallbacks() {
	app.Window.SetKeyCallback(app.scene.KeyCallback)
	app.Window.SetMouseButtonCallback(app.scene.MouseButtonCallback)
	app.Window.SetCursorPosCallback(app.scene.MouseMoveCallback)
	app.Window.SetScrollCallback(app.scene.ScrollCallback)
	app.updateActions = make([]func(delta float64), 0)
	app.updateActions = append(app.updateActions, app.scene.PreUpdate)
	if app.cfg.Physics.Enable {
		app.updateActions = append(app.updateActions, func(delta float64) {
			app.World.Step(delta, velocityIterations, positionIterations)
		})
	}
	app.updateActions = append(app.updateActions, app.scene.Update)
}

func (app *App) Loop() {
	if app.scene == nil {
		panic("scene wasn't set")
	}

	lastFrame := glfw.GetTime()
	for !app.Window.ShouldClose() {
		currentFrame := glfw.GetTime()
		dTime := currentFrame - lastFrame
		lastFrame = currentFrame

		app.OnUpdate(dTime)
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

func (app *App) OnUpdate(delta float64) {
	for _, action := range app.updateActions {
		action(delta)
	}
}

func (app *App) OnRender() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for d := range app.scene.Drawable().Elements {
		pos := d.GetPosition()
		app.drawer.Draw(float32(pos.X), float32(pos.Y), app.scale, d.Shader(), d.Shape())
	}

	glfw.PollEvents()
	app.Window.SwapBuffers()
}

func (app *App) Shutdown() {
	app.quit = true
}
