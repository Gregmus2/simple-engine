package controls

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type MouseControl struct {
	camera common.Camera

	sensitivity float32
	rotation    bool
	x, y        float64
	yaw, pitch  float32
	scale       float32
}

func NewMouseControl(camera common.Camera, cfg *common.Config) common.MouseController {
	return &MouseControl{camera: camera, sensitivity: cfg.Controls.Mouse.Sensitivity / 100}
}

func (c *MouseControl) MouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if button == glfw.MouseButton2 {
		switch action {
		case glfw.Press:
			c.x, c.y = w.GetCursorPos()
			c.rotation = true
		case glfw.Release:
			c.rotation = false
		}
	}
}

func (c *MouseControl) MouseMoveCallback(w *glfw.Window, x, y float64) {
	if c.rotation {
		view := c.camera.View()
		if x > 0 || y > 0 {
			c.yaw += -float32(c.x-x) * c.sensitivity
			c.pitch += -float32(c.y-y) * c.sensitivity

			view = mgl32.LookAtV(
				mgl32.Vec3{
					float32(math.Cos(float64(mgl32.DegToRad(c.yaw)))) * float32(math.Cos(float64(mgl32.DegToRad(c.pitch)))),
					float32(math.Sin(float64(mgl32.DegToRad(c.pitch)))),
					float32(math.Sin(float64(mgl32.DegToRad(c.yaw)))) * float32(math.Cos(float64(mgl32.DegToRad(c.pitch)))),
				},
				mgl32.Vec3{0, 0, 0},
				mgl32.Vec3{0, 1, 0},
			).Mul4(mgl32.Scale3D(0.3, 0.3, 0.3))
		}

		c.camera.UpdateView(view)
		c.x, c.y = x, y
	}
}

func (c *MouseControl) Update(delta float64) {

}
