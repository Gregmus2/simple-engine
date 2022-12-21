package graphics

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"math"
)

const DoublePI float64 = 2.0 * math.Pi

func Box(x, y, w, h float32, color Color) {
	Programs.Default.ApplyProgram(color)
	square := boxVertexes(x, y, w, h)
	vbo, vao := MakeDefaultVertexObjects(square)
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
	Programs.Default.ApplyProgram(color)
	circle := circleVertexes(x, y, r, 360)
	vbo, vao := MakeDefaultVertexObjects(circle)
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
	Programs.Default.ApplyProgram(color)
	vertexes := []float32{
		PosToUnitsX(x1), PosToUnitsY(y1), 0,
		PosToUnitsX(x2), PosToUnitsY(y2), 0,
	}

	//gl.Enable(gl.LINE_SMOOTH)

	vbo, vao := MakeDefaultVertexObjects(vertexes)
	gl.LineWidth(1.0)
	gl.DrawArrays(gl.LINES, 0, 2)
	ClearBuffers(vbo, vao)
	gl.UseProgram(0)
}

func Text(x, y float32, text string, color Color) {
	// todo replace
	scale := 20

	// activate corresponding render state
	Programs.Text.ApplyProgram(color)
	gl.Uniform1i(gl.GetUniformLocation(Programs.Text.program, gl.Str("text\x00")), 0)
	vbo, vao := MakeTextVertexObjects()
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindVertexArray(*vao)

	for _, r := range text {
		ch := Font.Characters[r]
		xpos := x + float32(ch.Bearing.X*scale)
		ypos := y - float32((ch.Size.Y-ch.Bearing.Y)*scale)
		w := ch.Size.X * scale
		h := ch.Size.Y * scale
		// update VBO for each character
		vertexes := textVertexes(xpos, ypos, float32(w), float32(h))
		// render glyph texture over quad
		gl.BindTexture(gl.TEXTURE_2D, ch.Texture)
		// update content of VBO memory
		gl.BindBuffer(gl.ARRAY_BUFFER, *vbo)
		gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(vertexes)*4, gl.Ptr(vertexes))
		gl.BindBuffer(gl.ARRAY_BUFFER, 0)
		// render quad
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
		// now advance cursors for next glyph (note that advance is number of 1/64 pixels)
		x += float32(int(ch.Advance>>6) * scale) // bitshift by 6 to get value in pixels (2^6 = 64)
	}

	ClearBuffers(vbo, vao)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.UseProgram(0)
}

func textVertexes(x, y, w, h float32) []float32 {
	x, y, w, h = PosToUnitsX(x), PosToUnitsY(y), PosToUnitsW(w), PosToUnitsH(h)

	return []float32{
		x, y + h, 0, 0,
		x, y, 0, 1,
		x + w, y, 1, 1,

		x, y + h, 0, 0,
		x + w, y, 1, 1,
		x + w, y + h, 1, 0,
	}
}
