package graphics

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"math"
)

const DoublePI float64 = 2.0 * math.Pi

func Box(x, y, w, h float32, color Color) {
	Program.ApplyProgram(color)
	square := boxVertexes(x, y, w, h)
	vbo := MakeVBO(square)
	vao := MakeVAO(vbo)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
	ClearBuffers(vbo, vao)
	gl.UseProgram(0)
}

func boxVertexes(x, y, w, h float32) []float32 {
	x, y, w, h = PosToUnitsX(x), PosToUnitsY(y), PosToUnitsW(w), PosToUnitsH(h)

	return []float32{
		x, y, 0,
		x + w, y, 0,
		x, y - h, 0,

		x, y - h, 0,
		x + w, y, 0,
		x + w, y - h, 0,
	}
}

func Circle(x, y, r float32, color Color) {
	Program.ApplyProgram(color)
	circle := circleVertexes(x, y, r, 360)
	vbo := MakeVBO(circle)
	vao := MakeVAO(vbo)
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(circle)/3))
	ClearBuffers(vbo, vao)
	gl.UseProgram(0)
}

func circleVertexes(x, y, r float32, sides int) []float32 {
	x, y, rW, rH := PosToUnitsX(x), PosToUnitsY(y), PosToUnitsW(r), PosToUnitsH(r)

	vertexes := make([]float32, (sides+2)*3)
	for i := 0; i < (sides+2)*3; i += 3 {
		vertexes[i] = x + (rW * float32(math.Cos(float64(i)/3*DoublePI/float64(sides))))
		vertexes[i+1] = y + (rH * float32(math.Sin(float64(i)/3*DoublePI/float64(sides))))
		vertexes[i+2] = 0
	}

	return vertexes
}

func Line(x1, y1, x2, y2 float32, color Color) {
	Program.ApplyProgram(color)
	vertexes := []float32{
		PosToUnitsX(x1), PosToUnitsY(y1), 0,
		PosToUnitsX(x2), PosToUnitsY(y2), 0,
	}

	//gl.Enable(gl.LINE_SMOOTH)

	vbo := MakeVBO(vertexes)
	vao := MakeVAO(vbo)
	gl.LineWidth(1.0)
	gl.DrawArrays(gl.LINES, 0, 2)
	ClearBuffers(vbo, vao)
	gl.UseProgram(0)
}
