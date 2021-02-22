package graphics

import (
	"github.com/Gregmus2/simple-engine/common"
)

type Drawer struct {
	camera common.Camera
}

func NewDrawer(camera common.Camera) *Drawer {
	return &Drawer{camera: camera}
}

func (d *Drawer) Draw(x, y, scale float32, shader common.Shader, shape common.Shape) {
	x, y = x*scale, y*scale
	projection, view, model := d.camera.Projection(), d.camera.View(), d.camera.Model(x, y)
	shader.ApplyShader(projection, view, model)

	shape.Draw()
}
