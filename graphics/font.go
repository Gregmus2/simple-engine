package graphics

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"golang.org/x/image/math/fixed"
	"image"
	"image/draw"
	"os"
)

var Font FontModel

type Character struct {
	Texture uint32
	Advance uint32 // Offset to advance to next glyph
	Size    image.Point
	Bearing image.Point // Offset from baseline to left/top of glyph
}

type FontModel struct {
	Characters map[rune]Character
}

func DefineFont(_ *OpenGL) error {
	fd, err := os.ReadFile("Arialn.ttf")
	if err != nil {
		return errors.Wrap(err, "error on reading font file")
	}

	ttf, err := truetype.Parse(fd)
	if err != nil {
		return errors.Wrap(err, "error on parsing font")
	}

	face := truetype.NewFace(ttf, &truetype.Options{
		Size: 20,
		DPI:  72,
	})
	defer face.Close()

	/*
		OpenGL requires that textures all have a 4-byte alignment e.g. their size is always a multiple of 4 bytes.
		Normally this won't be a problem since most textures have a width that is a multiple of 4 and/or
		use 4 bytes per pixel, but since we now only use a single byte per pixel,
		the texture can have any possible width. By setting its unpack alignment to 1
		we ensure there are no alignment issues (which could cause segmentation faults)
	*/
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	Font.Characters = make(map[rune]Character, 128)
	var r rune
	for r = 0; r < 128; r++ {
		dr, mask, maskp, adv, ok := face.Glyph(fixed.Point26_6{}, 'A')
		if !ok {
			return errors.New("error on getting glyph")
		}
		img := image.NewGray(dr)
		draw.DrawMask(img, dr, image.White, maskp, mask, maskp, draw.Over)

		// generate texture
		var texture uint32 = uint32(r)
		gl.GenTextures(1, &texture)
		gl.BindTexture(gl.TEXTURE_2D, texture)
		gl.TexImage2D(
			gl.TEXTURE_2D,
			0,
			gl.RED,
			int32(dr.Size().X),
			int32(dr.Size().Y),
			0,
			gl.RED,
			gl.UNSIGNED_BYTE,
			gl.Ptr(img.Pix),
		)
		// set texture options
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		// now store character for later use
		Font.Characters[r] = Character{
			Texture: texture,
			Advance: uint32(adv),
			Size:    dr.Size(),
			Bearing: image.Point{X: dr.Min.X, Y: dr.Max.Y},
		}
	}

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return nil
}
