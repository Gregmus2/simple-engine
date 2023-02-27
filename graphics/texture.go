package graphics

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/pkg/errors"
	"image"
	"image/draw"
	"log"
	"os"
)

var Textures = make(map[string]*Texture)

type Texture struct {
	texture uint32
	target  uint32 // same target as gl.BindTexture(<this param>, ...)
	texUnit uint32 // Texture unit that is currently bound to ex: gl.TEXTURE0
}

var errUnsupportedStride = errors.New("unsupported stride, only 32-bit colors supported")
var errTextureNotBound = errors.New("Texture not bound")

func LoadTextures(_ *OpenGL) {
	if common.Config.Graphics.Textures == "" {
		return
	}

	entries, err := os.ReadDir(common.Config.Graphics.Textures)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		tex, err := newTextureFromFile(common.Config.Graphics.Textures + "/" + e.Name())
		if err != nil {
			panic(err)
		}

		Textures[e.Name()] = tex
	}
}

func newTextureFromFile(file string) (*Texture, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	// Decode detexts the type of image as long as its image/<type> is imported
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return newTexture(img)
}

func newTexture(img image.Image) (*Texture, error) {
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, errUnsupportedStride
	}

	var texture uint32
	gl.GenTextures(1, &texture)

	target := uint32(gl.TEXTURE_2D)
	internalFmt := int32(gl.SRGB_ALPHA)
	format := uint32(gl.RGBA)
	width := int32(rgba.Rect.Size().X)
	height := int32(rgba.Rect.Size().Y)
	pixType := uint32(gl.UNSIGNED_BYTE)
	dataPtr := gl.Ptr(rgba.Pix)

	tex := Texture{
		texture: texture,
		target:  target,
	}

	tex.Bind(gl.TEXTURE0)
	defer tex.UnBind()

	gl.TexParameteri(tex.target, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(tex.target, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(tex.target, gl.TEXTURE_MIN_FILTER, gl.LINEAR) // minification filter
	gl.TexParameteri(tex.target, gl.TEXTURE_MAG_FILTER, gl.LINEAR) // magnification filter

	gl.TexImage2D(target, 0, internalFmt, width, height, 0, format, pixType, dataPtr)

	gl.GenerateMipmap(tex.texture)

	return &tex, nil
}

func (tex *Texture) Bind(texUnit uint32) {
	gl.ActiveTexture(texUnit)
	gl.BindTexture(tex.target, tex.texture)
	tex.texUnit = texUnit
}

func (tex *Texture) UnBind() {
	tex.texUnit = 0
	gl.BindTexture(tex.target, 0)
}
