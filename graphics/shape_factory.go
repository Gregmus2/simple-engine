package graphics

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"math"
)

const doublePI float64 = 2.0 * math.Pi

type ShapeHelper struct {
	helper *PosToUnitsConverter
}

func NewShapeFactory(h *PosToUnitsConverter) *ShapeHelper {
	return &ShapeHelper{helper: h}
}

func (f *ShapeHelper) Box(x, y, w, h float32) {
	square := f.boxVertexes(x, y, w, h)
	vao := MakeVAO(square)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}

func (f *ShapeHelper) boxVertexes(x, y, w, h float32) []float32 {
	x, y, w, h = f.helper.X(x), f.helper.Y(y), f.helper.W(w), f.helper.H(h)

	return []float32{
		x, y, 0,
		x + w, y, 0,
		x, y - h, 0,

		x, y - h, 0,
		x + w, y, 0,
		x + w, y - h, 0,
	}
}

func (f *ShapeHelper) Circle(x, y, r float32) {
	circle := f.circleVertexes(x, y, r, 360)
	vao := MakeVAO(circle)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(circle)/3))
}

func (f *ShapeHelper) circleVertexes(x, y, r float32, sides int) []float32 {
	x, y, rW, rH := f.helper.X(x), f.helper.Y(y), f.helper.W(r), f.helper.H(r)

	vertexes := make([]float32, (sides+2)*3)
	for i := 0; i < (sides+2)*3; i += 3 {
		vertexes[i] = x + (rW * float32(math.Cos(float64(i)/3*doublePI/float64(sides))))
		vertexes[i+1] = y + (rH * float32(math.Sin(float64(i)/3*doublePI/float64(sides))))
		vertexes[i+2] = 0
	}

	return vertexes
}
