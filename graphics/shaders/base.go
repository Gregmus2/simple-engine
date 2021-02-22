package shaders

import (
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type BaseShader struct {
	program uint32
	color   graphics.Color
}

func (s *BaseShader) useProgram() {
	gl.UseProgram(s.program)
}

func (s *BaseShader) applyColor() {
	gl.Uniform3f(gl.GetUniformLocation(s.program, gl.Str("objectColor\x00")), s.color.R, s.color.G, s.color.B)
}

func (s *BaseShader) applyProjection(projection *mgl32.Mat4) {
	projectionUniform := gl.GetUniformLocation(s.program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])
}

func (s *BaseShader) applyCamera(view *mgl32.Mat4) {
	cameraUniform := gl.GetUniformLocation(s.program, gl.Str("view\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &view[0])
}

func (s *BaseShader) applyModel(model *mgl32.Mat4) {
	modelUniform := gl.GetUniformLocation(s.program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
}
