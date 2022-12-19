package graphics

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sirupsen/logrus"
	"strings"
)

var Program *ProgramManager

func DefineProgram(_ *OpenGL, _ *glfw.Window) {
	Program = &ProgramManager{}
	Program.generateProgram()
}

type ProgramManager struct {
	program uint32
}

const vertexShaderSource string = `
    #version 410
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

const fragmentShaderTemplate string = `
    #version 410
    out vec4 frag_colour;
	uniform vec3 color;
    void main() {
        frag_colour = vec4(color, 1.0);
    }
` + "\x00"

func (c *ProgramManager) generateProgram() {
	vertexShader, err := c.compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		logrus.WithError(err).Error("error on compile vertex shader")
		return
	}

	fragmentShader, err := c.compileShader(fragmentShaderTemplate, gl.FRAGMENT_SHADER)
	if err != nil {
		logrus.WithError(err).Error("error on compile vertex shader")
		return
	}

	c.program = gl.CreateProgram()
	gl.AttachShader(c.program, vertexShader)
	gl.AttachShader(c.program, fragmentShader)
	gl.LinkProgram(c.program)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
}

func (c *ProgramManager) ApplyProgram(color Color) {
	gl.UseProgram(c.program)
	gl.Uniform3f(gl.GetUniformLocation(c.program, gl.Str("color\x00")), color.R, color.G, color.B)
}

func (c *ProgramManager) compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
