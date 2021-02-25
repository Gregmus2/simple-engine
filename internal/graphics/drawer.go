package graphics

import (
	"github.com/Gregmus2/simple-engine/internal/common"
)

type Drawer struct {
	builder *PositionBuilder
}

func NewDrawer(b *PositionBuilder) *Drawer {
	return &Drawer{builder: b}
}

func (d *Drawer) Draw(x, y float32, shader common.Shader, shape common.Shape) {
	projection, view, model := d.builder.Projection(), d.builder.View(), d.builder.Model(x, y)
	shader.ApplyShader(projection, view, model)

	shape.Draw()
}
