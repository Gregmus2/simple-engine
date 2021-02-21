package graphics

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"math"
)

const DoublePI float64 = 2.0 * math.Pi

type ShapeHelper struct {
	helper *PosToUnitsConverter
}

func NewShapeFactory(h *PosToUnitsConverter) *ShapeHelper {
	return &ShapeHelper{helper: h}
}

func (s *ShapeHelper) Box(w, h, d float32, vert *uint32) (*uint32, *uint32) {
	box := s.boxVertexes(w, h, d)
	vbo := MakeVBO(box)
	vao := MakeVAO(vert)

	return vbo, vao
}

func (s *ShapeHelper) boxVertexes(w, h, d float32) []float32 {
	// width, height, depth
	w, h, d = s.helper.W(w), s.helper.H(h), s.helper.D(d)
	// half width, half height, half depth
	hW, hH, hD := w/2, h/2, d/2
	// left, right, top, bottom, near, far
	l, r, t, b, n, f := -hW, hW, hH, -hH, hD, -hD

	return []float32{
		l, t, f, 0, 0,
		r, t, f, 1, 0,
		r, b, f, 1, 1,
		r, b, f, 1, 1,
		l, b, f, 0, 1,
		l, t, f, 0, 0,

		l, t, n, 0, 0,
		r, t, n, 1, 0,
		r, b, n, 1, 1,
		r, b, n, 1, 1,
		l, b, n, 0, 1,
		l, t, n, 0, 0,

		r, b, n, 1, 0,
		r, b, f, 1, 1,
		r, t, f, 0, 1,
		r, t, f, 0, 1,
		r, t, n, 0, 0,
		r, b, n, 1, 0,

		l, b, n, 1, 0,
		l, b, f, 1, 1,
		l, t, f, 0, 1,
		l, t, f, 0, 1,
		l, t, n, 0, 0,
		l, b, n, 1, 0,

		l, t, f, 0, 1,
		r, t, f, 1, 1,
		r, t, n, 1, 0,
		r, t, n, 1, 0,
		l, t, n, 0, 0,
		l, t, f, 0, 1,

		l, b, f, 0, 1,
		r, b, f, 1, 1,
		r, b, n, 1, 0,
		r, b, n, 1, 0,
		l, b, n, 0, 0,
		l, b, f, 0, 1,
	}
}

func (s *ShapeHelper) Circle(x, y, r float32) {
	circle := s.circleVertexes(x, y, r, 360)
	vbo := MakeVBO(circle)
	vao := MakeVAO(vbo)
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(circle)/3))
	ClearBuffers(vbo, vao)
}

func (s *ShapeHelper) circleVertexes(x, y, r float32, sides int) []float32 {
	x, y, rW, rH := s.helper.X(x), s.helper.Y(y), s.helper.W(r), s.helper.H(r)

	vertexes := make([]float32, (sides+2)*3)
	for i := 0; i < (sides+2)*3; i += 3 {
		vertexes[i] = x + (rW * float32(math.Cos(float64(i)/3*DoublePI/float64(sides))))
		vertexes[i+1] = y + (rH * float32(math.Sin(float64(i)/3*DoublePI/float64(sides))))
		vertexes[i+2] = 0
	}

	return vertexes
}

func (s *ShapeHelper) Line(x1, y1, x2, y2 float32) {
	vertexes := []float32{
		s.helper.X(x1), s.helper.Y(y1), 0,
		s.helper.X(x2), s.helper.Y(y2), 0,
	}

	//gl.Enable(gl.LINE_SMOOTH)

	vbo := MakeVBO(vertexes)
	vao := MakeVAO(vbo)
	gl.LineWidth(1.0)
	gl.DrawArrays(gl.LINES, 0, 2)
	ClearBuffers(vbo, vao)
}
