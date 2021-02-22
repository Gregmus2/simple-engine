package graphics

import "github.com/go-gl/gl/v4.6-core/gl"

type DrawFunctionsDictionary struct {
}

func NewDrawFunctionsDictionary() *DrawFunctionsDictionary {
	return &DrawFunctionsDictionary{}
}

func (d *DrawFunctionsDictionary) Box() {
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
}
