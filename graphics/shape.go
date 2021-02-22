package graphics

import "github.com/go-gl/gl/v4.6-core/gl"

type Shape struct {
	vao, vbo *uint32
	drawFunc func()
}

func NewShape(vao, vbo *uint32, drawFunc func()) *Shape {
	return &Shape{
		vao: vao, vbo: vbo, drawFunc: drawFunc,
	}
}

func (s *Shape) Remove() {
	ClearBuffers(s.vbo, s.vao)
}

func (s *Shape) Draw() {
	gl.BindVertexArray(*s.vao)
	s.drawFunc()
}
