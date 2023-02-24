package graphics

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pkg/errors"
	"strings"
)

var Programs = struct {
	Default *Program
	Text    *Program
}{}

func DefinePrograms(_ *OpenGL, _ *glfw.Window) {
	Programs = struct {
		Default *Program
		Text    *Program
	}{Default: &Program{}, Text: &Program{}}
	Programs.Default.generateProgram(defaultVertexShaderSource, defaultFragmentShaderTemplate)
	Programs.Text.generateProgram(textVertexShaderSource, textFragmentShaderTemplate)
}

type Program struct {
	program uint32
}

func (p *Program) generateProgram(vertexShaderSource, fragmentShaderTemplate string) {
	vertexShader, err := p.compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(errors.Wrap(err, "error on compile vertex shader"))
	}

	fragmentShader, err := p.compileShader(fragmentShaderTemplate, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(errors.Wrap(err, "error on compile fragment shader"))
	}

	p.program = gl.CreateProgram()
	gl.AttachShader(p.program, vertexShader)
	gl.AttachShader(p.program, fragmentShader)
	gl.LinkProgram(p.program)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
}

func (p *Program) ApplyProgram(color Color) {
	gl.UseProgram(p.program)
	gl.Uniform3f(gl.GetUniformLocation(p.program, gl.Str("color\x00")), color.R, color.G, color.B)
}

func (p *Program) compileShader(source string, shaderType uint32) (uint32, error) {
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
