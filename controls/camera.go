package controls

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type CameraControl struct {
	camera common.Camera

	sensitivity              float32
	rotation                 bool
	lastCursorX, lastCursorY float64
	yaw, pitch               float32

	xVelocity, yVelocity, zVelocity float32
	scale                           float32
	x, y, z                         float32
	xCenter, yCenter, zCenter       float32
}

func NewMouseControl(camera common.Camera, cfg *common.Config) common.CameraControl {
	return &CameraControl{
		camera:      camera,
		sensitivity: cfg.Controls.Mouse.Sensitivity / 100,
		scale:       0.3,
		pitch:       -7.6,
		yaw:         -270,
		z:           -1,
	}
}

func (c *CameraControl) MouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action) {
	if button == glfw.MouseButton2 {
		switch action {
		case glfw.Press:
			c.lastCursorX, c.lastCursorY = w.GetCursorPos()
			c.rotation = true
		case glfw.Release:
			c.rotation = false
		}
	}
}

func (c *CameraControl) MouseMove(x, y float64) {
	if c.rotation {
		c.yaw += -float32(c.lastCursorX-x) * c.sensitivity
		c.pitch += -float32(c.lastCursorY-y) * c.sensitivity
		if c.pitch > 89.0 {
			c.pitch = 89.0
		}
		if c.pitch < -89.0 {
			c.pitch = -89.0
		}

		c.updateCenter()
		c.updateView()

		c.lastCursorX, c.lastCursorY = x, y
	}
}

func (c *CameraControl) updateCenter() {
	c.xCenter = float32(math.Cos(float64(mgl32.DegToRad(c.yaw)))) * float32(math.Cos(float64(mgl32.DegToRad(c.pitch))))
	c.yCenter = -float32(math.Sin(float64(mgl32.DegToRad(c.pitch))))
	c.zCenter = float32(math.Sin(float64(mgl32.DegToRad(c.yaw)))) * float32(math.Cos(float64(mgl32.DegToRad(c.pitch))))
}

func (c *CameraControl) updateView() {
	view := mgl32.LookAtV(
		mgl32.Vec3{c.x, c.y, c.z},
		mgl32.Vec3{c.xCenter + c.x, c.yCenter + c.y, c.zCenter + c.z},
		mgl32.Vec3{0, 1, 0},
	).Mul4(mgl32.Scale3D(c.scale, c.scale, c.scale))
	c.camera.UpdateView(view)
}

func (c *CameraControl) Scroll(yOffset float64) {
	c.scale += float32(yOffset) * 0.05
	c.updateView()
}

func (c *CameraControl) Key(key glfw.Key, action glfw.Action) {
	if action == glfw.Press {
		switch key {
		case glfw.KeyA:
			c.xVelocity = -1
		case glfw.KeyD:
			c.xVelocity = 1
		case glfw.KeyW:
			c.zVelocity = 1
		case glfw.KeyS:
			c.zVelocity = -1
		case glfw.KeyLeftShift:
			c.yVelocity = -1
		case glfw.KeySpace:
			c.yVelocity = 1
		case glfw.KeyTab:
			c.rotation = !c.rotation
		}
	}

	if action == glfw.Release {
		switch key {
		case glfw.KeyD, glfw.KeyA:
			c.xVelocity = 0.0
		case glfw.KeyW, glfw.KeyS:
			c.zVelocity = 0.0
		case glfw.KeyLeftShift, glfw.KeySpace:
			c.yVelocity = 0.0
		}
	}
}

func (c *CameraControl) Update(delta float64) {
	if c.xVelocity != 0 {
		c.x += (c.xVelocity * float32(delta)) * float32(math.Cos(float64(mgl32.DegToRad(c.yaw+90)))) * float32(math.Cos(float64(mgl32.DegToRad(c.pitch))))
		c.z += (c.xVelocity * float32(delta)) * float32(math.Sin(float64(mgl32.DegToRad(c.yaw+90)))) * float32(math.Cos(float64(mgl32.DegToRad(c.pitch))))
		c.updateView()
	}
	if c.yVelocity != 0 {
		c.y += (c.yVelocity * float32(delta)) * -float32(math.Sin(float64(mgl32.DegToRad(c.pitch+90))))
		c.z += (c.yVelocity * float32(delta)) * float32(math.Sin(float64(mgl32.DegToRad(c.yaw)))) * float32(math.Cos(float64(mgl32.DegToRad(c.pitch+90))))
		c.updateView()
	}
	if c.zVelocity != 0 {
		c.x += (c.zVelocity * float32(delta)) * float32(math.Cos(float64(mgl32.DegToRad(c.yaw)))) * float32(math.Cos(float64(mgl32.DegToRad(c.pitch))))
		c.y += (c.zVelocity * float32(delta)) * -float32(math.Sin(float64(mgl32.DegToRad(c.pitch))))
		c.z += (c.zVelocity * float32(delta)) * float32(math.Sin(float64(mgl32.DegToRad(c.yaw)))) * float32(math.Cos(float64(mgl32.DegToRad(c.pitch))))
		c.updateView()
	}
}
