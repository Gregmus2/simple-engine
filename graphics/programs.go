package graphics

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/sirupsen/logrus"
	"strings"
)

type Programs struct {
	SimpleColor uint32
}

const vertexShaderSource string = `
    #version 410
	layout (location = 0) in vec3 vert;
    uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;
    void main() {
        gl_Position = projection * view * model * vec4(vert, 1.0);
    }
` + "\x00"

const fragmentShaderTemplate string = `
    #version 410
	out vec4 frag_colour;
    uniform vec3 objectColor;
	uniform vec3 lightColor;
    void main() {
        frag_colour = vec4(objectColor * lightColor, 1.0);
    }
` + "\x00"

// todo need to compile all colors at the beginning
func NewPrograms() *Programs {
	p := &Programs{}
	p.SimpleColor = p.buildProgram()

	return p
}

func (c *Programs) buildProgram() uint32 {
	vertexShader, err := c.compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		logrus.WithError(err).Error("error on compile vertex shader")
		return 0
	}

	fragmentShader, err := c.compileShader(fragmentShaderTemplate, gl.FRAGMENT_SHADER)
	if err != nil {
		logrus.WithError(err).Error("error on compile vertex shader")
		return 0
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return prog
}

func (c *Programs) compileShader(source string, shaderType uint32) (uint32, error) {
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
