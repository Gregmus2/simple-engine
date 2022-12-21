package graphics

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/pkg/errors"
	"log"
	"unsafe"
)

type OpenGL struct{}

func NewOpenGL(init common.Init) (*OpenGL, error) {
	if err := gl.Init(); err != nil {
		return nil, errors.Wrap(err, "failed to initialize openGL")
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.Enable(gl.MULTISAMPLE)
	// to render text, which is semi-transparent textures
	gl.Enable(gl.BLEND)
	// Factor is equal to 1−alpha of the source color vector
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	if err := init.OpenGL(); err != nil {
		return nil, err
	}

	if common.Config.Graphics.Debug {
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

func MakeDefaultVertexObjects(points []float32) (*uint32, *uint32) {
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

	return &vbo, &vao
}

func MakeTextVertexObjects() (*uint32, *uint32) {
	var vao, vbo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*6*4, nil, gl.DYNAMIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 4*4, nil)

	return &vbo, &vao
}

func ClearBuffers(vbo, vao *uint32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.DisableVertexAttribArray(0)
	gl.DeleteBuffers(1, vbo)
	gl.DeleteVertexArrays(1, vao)
}
