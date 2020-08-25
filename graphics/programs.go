package graphics

import (
	"fmt"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/sirupsen/logrus"
	"strings"
)

type ProgramFactory struct {
	programs map[string]uint32
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
    void main() {
        frag_colour = vec4(%f, %f, %f, %f);
    }
` + "\x00"

func NewProgramFactory() *ProgramFactory {
	return &ProgramFactory{programs: make(map[string]uint32)}
}

func (c *ProgramFactory) GetByColor(color Color) uint32 {
	prog, exists := c.programs[c.hash(color)]
	if exists {
		return prog
	}

	prog = c.buildProgram(color)
	c.programs[c.hash(color)] = prog

	return prog
}

func (c *ProgramFactory) hash(color Color) string {
	return fmt.Sprintf("%f %f %f %f", color.A, color.R, color.G, color.B)
}

func (c *ProgramFactory) buildProgram(color Color) uint32 {
	vertexShader, err := c.compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		logrus.WithError(err).Error("error on compile vertex shader")
		return 0
	}

	fragmentShaderSource := fmt.Sprintf(fragmentShaderTemplate, color.R, color.G, color.B, color.A)
	fragmentShader, err := c.compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		logrus.WithError(err).Error("error on compile vertex shader")
		return 0
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)

	return prog
}

func (c *ProgramFactory) compileShader(source string, shaderType uint32) (uint32, error) {
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
