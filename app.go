package engine

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/dispatchers"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"go.uber.org/dig"
)

type App struct {
	Window *glfw.Window
	World  *box2d.B2World
	GL     *graphics.OpenGL
	update *dispatchers.Update

	cfg    *common.Config
	drawer *graphics.Drawer
	scale  float32
	scene  common.Scene
	quit   bool
}

const velocityIterations = 8
const positionIterations = 2

type Params struct {
	dig.In

	Config *common.Config
	Window *glfw.Window
	OpenGL *graphics.OpenGL
	World  *box2d.B2World
	Drawer *graphics.Drawer
	Update *dispatchers.Update
}

func NewApp(params Params) (*App, error) {
	return &App{
		cfg:    params.Config,
		Window: params.Window,
		GL:     params.OpenGL,
		World:  params.World,
		drawer: params.Drawer,
		update: params.Update,
		scale:  params.Config.Graphics.Scale,
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

	app.update.Subscribe(app.scene.PreUpdate)
	if app.cfg.Physics.Enable {
		app.update.Subscribe(func(delta float64) {
			app.World.Step(delta, velocityIterations, positionIterations)
		})
	}
	app.update.Subscribe(app.scene.Update)
}

func (app *App) Loop() {
	if app.scene == nil {
		panic("scene wasn't set")
	}

	lastTime := glfw.GetTime()
	for !app.Window.ShouldClose() && !app.quit {
		currentTime := glfw.GetTime()
		dTime := currentTime - lastTime
		lastTime = currentTime

		app.OnUpdate(dTime)
		app.OnRender()
	}

	app.Destroy()
}

func (app *App) Destroy() {
	app.Window.Destroy()
	glfw.Terminate()
}

func (app *App) OnUpdate(delta float64) {
	app.update.Dispatch(delta)
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
