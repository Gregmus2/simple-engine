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
	/*
		OpenGL requires that textures all have a 4-byte alignment e.g. their size is always a multiple of 4 bytes.
		Normally this won't be a problem since most textures have a width that is a multiple of 4 and/or
		use 4 bytes per pixel, but since we now only use a single byte per pixel,
		the texture can have any possible width. By setting its unpack alignment to 1
		we ensure there are no alignment issues (which could cause segmentation faults)
	*/
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)

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

func MakeTextureVertexObjects(vertices []float32, indices []uint32) (*uint32, *uint32, *uint32) {
	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)

	// copy vertices data into VBO (it needs to be bound first)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// copy indices into element buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	var stride int32 = 3*4 + 3*4 + 2*4
	var offset = 0

	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(0)
	offset += 3 * 4

	// color
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(1)
	offset += 3 * 4

	// texture position
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(2)
	offset += 2 * 4

	// unbind the VAO (safe practice, so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return &vao, &vbo, &ebo
}

func ClearBuffers(vbo, vao *uint32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.DisableVertexAttribArray(0)
	gl.DeleteBuffers(1, vbo)
	gl.DeleteVertexArrays(1, vao)
}
