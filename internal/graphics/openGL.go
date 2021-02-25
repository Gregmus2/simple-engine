package graphics

import (
	"github.com/Gregmus2/simple-engine/internal/common"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/pkg/errors"
	"log"
	"unsafe"
)

type OpenGL struct{}

func NewOpenGL(cfg *common.Config, init common.Init) (*OpenGL, error) {
	if err := gl.Init(); err != nil {
		return nil, errors.Wrap(err, "failed to initialize openGL")
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.Enable(gl.MULTISAMPLE)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	err := init.OpenGL()
	if err != nil {
		return nil, err
	}

	if cfg.Graphics.Debug {
		gl.Enable(gl.DEBUG_OUTPUT)
		gl.Enable(gl.DEBUG_OUTPUT_SYNCHRONOUS)
		debugMessage := func(
			source uint32,
			gltype uint32,
			id uint32,
			severity uint32,
			length int32,
			message string,
			userParam unsafe.Pointer) {
			log.Println(source, "|", gltype, "|", severity, "|", message, "|", userParam)
		}
		gl.DebugMessageCallback(debugMessage, nil)
		gl.DebugMessageControl(gl.DONT_CARE, gl.DONT_CARE, gl.DONT_CARE, 0, nil, true)
		// todo print GPU memory
	}

	return &OpenGL{}, nil
}

func MakeVBO(points []float32) *uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	return &vbo
}

func MakeVAO(light bool) *uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	if !light {
		gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(3*4))
		gl.EnableVertexAttribArray(1)
	}

	return &vao
}

func ClearBuffers(vbo, vao *uint32) {
	gl.DisableVertexAttribArray(0)
	gl.DeleteBuffers(1, vbo)
	gl.DeleteVertexArrays(1, vao)
}
