package graphics

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	_ "image/png"
	"math"
	"unsafe"
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

func Circle(x, y, r float32, color Color) {
	Programs.Default.ApplyProgram(color)
	circle := circleVertexes(x, y, r, 360)
	vbo, vao := MakeDefaultVertexObjects(circle)
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(circle)/3))
	ClearBuffers(vbo, vao)
	gl.UseProgram(0)
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

func Text(x, y, scale float32, text string, font *Font, color Color) {
	indices := []rune(text)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.UseProgram(font.program)
	gl.Uniform4f(gl.GetUniformLocation(font.program, gl.Str("textColor\x00")), color.R, color.G, color.B, color.A)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindVertexArray(font.vao)

	// Iterate through all characters in string
	for _, r := range indices {
		ch, ok := font.Characters[r]

		//load missing runes in batches of 32
		if !ok {
			panic("glyph wasn't generated for this character")
		}

		//skip runes that are not in font chacter range
		if !ok {
			panic(fmt.Errorf("char %s is not supported", string(r)))
		}

		gl.BindTexture(gl.TEXTURE_2D, ch.textureID)
		gl.BindBuffer(gl.ARRAY_BUFFER, font.vbo)

		vertices := textVertices(x, y, scale, ch)

		gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(vertices)*4, gl.Ptr(vertices))
		gl.DrawArrays(gl.TRIANGLES, 0, 16)

		gl.BindBuffer(gl.ARRAY_BUFFER, 0)
		// Now advance cursors for next glyph (note that advance is number of 1/64 pixels)
		x += float32(ch.advance>>6) * scale // Bitshift by 6 to get value in pixels (2^6 = 64 (divide amount of 1/64th pixels by 64 to get amount of pixels))

	}

	//clear opengl textures and programs
	gl.BindVertexArray(0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.UseProgram(0)
	gl.Disable(gl.BLEND)
}

func DrawTexture(x, y, w, h float32, texture string) {
	vertices := textureVertexes(x, y, w, h)
	indices := []uint32{
		// rectangle
		0, 1, 2, // top triangle
		0, 2, 3, // bottom triangle
	}

	vao, vbo, ebo := MakeTextureVertexObjects(vertices, indices)

	tex := Textures[texture]
	Programs.Texture.ApplyProgram(White())
	tex.Bind(gl.TEXTURE0)
	Programs.Texture.UniformTexture(tex.texUnit)

	gl.BindVertexArray(*vao)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))

	tex.UnBind()

	ClearBuffers(vbo, vao)
	gl.DeleteBuffers(1, ebo)
}

func textureVertexes(x, y, w, h float32) []float32 {
	x, y, w, h = PosToUnitsX(x), PosToUnitsY(y), PosToUnitsW(w), PosToUnitsH(h)

	return []float32{
		// top left
		x, y, 0.0, // position
		1.0, 0.0, 0.0, // Color
		1.0, 0.0, // texture coordinates

		// top right
		x + w, y, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0,

		// bottom right
		x + w, y - h, 0.0,
		0.0, 0.0, 1.0,
		0.0, 1.0,

		// bottom left
		x, y - h, 0.0,
		1.0, 1.0, 1.0,
		1.0, 1.0,
	}
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

// calculate position and size
func textVertices(x, y, scale float32, ch *Character) []float32 {
	xPos := x + float32(ch.bearingH)*scale
	yPos := y - float32(ch.height-ch.bearingV)*scale
	w := float32(ch.width) * scale
	h := float32(ch.height) * scale

	return []float32{
		xPos + w, yPos, 1.0, 0.0,
		xPos, yPos, 0.0, 0.0,
		xPos, yPos + h, 0.0, 1.0,

		xPos, yPos + h, 0.0, 1.0,
		xPos + w, yPos + h, 1.0, 1.0,
		xPos + w, yPos, 1.0, 0.0,
	}
}
