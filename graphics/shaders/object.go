package shaders

import (
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type ObjectShader struct {
	BaseShader
}

func NewObjectShader(program uint32, color graphics.Color) *ObjectShader {
	return &ObjectShader{BaseShader{color: color, program: program}}
}

func (s *ObjectShader) ApplyShader(projection, view, model *mgl32.Mat4) {
	s.useProgram()
	s.applyProjection(projection)
	s.applyCamera(view)
	s.applyModel(model)
	s.applyColor()

	gl.Uniform3f(gl.GetUniformLocation(s.program, gl.Str("lightColor\x00")), 1, 1, 1)
	gl.Uniform3f(gl.GetUniformLocation(s.program, gl.Str("lightPos\x00")), 0.2, 0.2, -0.1)
	gl.Uniform3f(gl.GetUniformLocation(s.program, gl.Str("viewPos\x00")), 1, 1, 1)
}
