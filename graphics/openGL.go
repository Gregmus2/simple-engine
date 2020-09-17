package graphics

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/pkg/errors"
	"log"
)

type OpenGL struct{}

func NewOpenGL() (*OpenGL, error) {
	if err := gl.Init(); err != nil {
		return nil, errors.Wrap(err, "failed to initialize openGL")
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.Enable(gl.MULTISAMPLE)

	return &OpenGL{}, nil
}

// Vertex Array Object
func MakeVAO(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}
