package shaders

import (
	"github.com/go-gl/mathgl/mgl32"
)

type LightShader struct {
	BaseShader
}

func NewLightShader(program uint32) *LightShader {
	return &LightShader{BaseShader{program: program}}
}

func (s *LightShader) ApplyShader(projection, view, model mgl32.Mat4) {
	s.useProgram()
	s.applyProjection(projection)
	s.applyCamera(view)
	s.applyModel(model)
}
