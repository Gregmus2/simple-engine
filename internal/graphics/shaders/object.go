package shaders

import (
	"github.com/Gregmus2/simple-engine/internal"
	"github.com/Gregmus2/simple-engine/internal/graphics"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type ObjectShader struct {
	BaseShader
	camera *graphics.Camera
	app    *internal.App
	conv   *graphics.PositionConverter
}

func NewObjectShader(program uint32, c graphics.Color, cam *graphics.Camera, a *internal.App, conv *graphics.PositionConverter) *ObjectShader {
	return &ObjectShader{BaseShader{color: c, program: program}, cam, a, conv}
}

func (s *ObjectShader) ApplyShader(projection, view, model mgl32.Mat4) {
	s.useProgram()
	s.applyProjection(projection)
	s.applyCamera(view)
	s.applyModel(model)
	s.applyColor()

	gl.Uniform3f(gl.GetUniformLocation(s.program, gl.Str("lightColor\x00")), 1, 1, 1)
	collection := s.app.Scene().GetLights()
	for el := range collection.Elements {
		pos := el.GetPosition()
		x, y := s.conv.Box2DToOpenGL(float32(pos.X), float32(pos.Y))
		gl.Uniform3f(gl.GetUniformLocation(s.program, gl.Str("lightPos\x00")), x, y, 0)
		break
	}
	x, y, z := s.camera.GetPosition()
	gl.Uniform3f(gl.GetUniformLocation(s.program, gl.Str("viewPos\x00")), x, y, z)
}
