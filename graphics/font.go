package graphics

import (
	"fmt"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"os"
)

const highChar = 127
const lowChar = 27
const fontDPI = 72

var Fonts = make(map[string]*Font)

type Character struct {
	textureID uint32 // ID handle of the glyph texture
	width     int    //glyph width
	height    int    //glyph height
	advance   int    //glyph advance
	bearingH  int    //glyph bearing horizontal
	bearingV  int    //glyph bearing vertical
}

type Font struct {
	Characters   map[rune]*Character
	trueTypeFont *truetype.Font
	scale        float32
	vao          uint32
	vbo          uint32
	program      uint32
	texture      uint32 // Holds the glyph texture id.
}

func GetFont(font string, scale float32) *Font {
	key := fmt.Sprintf("%s%f", font, scale)
	if f, ok := Fonts[key]; ok {
		return f
	}

	path := common.Config.Graphics.Font[font]
	f, err := createFont(path, scale)
	if err != nil {
		panic(errors.Wrap(err, "error on creating font"))
	}
	Fonts[key] = f

	return f
}

func createFont(path string, scale float32) (*Font, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "error on reading font file")
	}

	trueTypeFont, err := truetype.Parse(file)
	if err != nil {
		return nil, errors.Wrap(err, "error on parsing font")
	}

	fontModel := &Font{
		Characters:   make(map[rune]*Character, highChar-lowChar),
		trueTypeFont: trueTypeFont,
		scale:        scale,
		program:      Programs.Text.program,
	}

	err = fontModel.GenerateGlyphs()
	if err != nil {
		panic(err)
	}

	gl.UseProgram(Programs.Text.program)
	resUniform := gl.GetUniformLocation(Programs.Text.program, gl.Str("resolution\x00"))
	gl.Uniform2f(resUniform, float32(common.Config.Window.W), float32(common.Config.Window.H))

	gl.GenVertexArrays(1, &fontModel.vao)
	gl.GenBuffers(1, &fontModel.vbo)
	gl.BindVertexArray(fontModel.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, fontModel.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 6*4*4, nil, gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(fontModel.program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(0))
	defer gl.DisableVertexAttribArray(vertAttrib)

	texCoordAttrib := uint32(gl.GetAttribLocation(fontModel.program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(2*4))
	defer gl.DisableVertexAttribArray(texCoordAttrib)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return fontModel, nil
}

func (f *Font) GenerateGlyphs() error {
	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(f.trueTypeFont)
	c.SetFontSize(float64(f.scale))
	c.SetHinting(font.HintingFull)

	ttfFace := truetype.NewFace(f.trueTypeFont, &truetype.Options{
		Size:    float64(f.scale),
		DPI:     fontDPI,
		Hinting: font.HintingFull,
	})

	// make each glyph
	var ch rune
	for ch = lowChar; ch <= highChar; ch++ {
		gBnd, gAdv, ok := ttfFace.GlyphBounds(ch)
		if ok != true {
			return fmt.Errorf("trueTypeFont face glyphBounds error")
		}

		gh := int32((gBnd.Max.Y - gBnd.Min.Y) >> 6)
		gw := int32((gBnd.Max.X - gBnd.Min.X) >> 6)

		// if glyph has no dimensions set to a max value
		if gw == 0 || gh == 0 {
			gBnd = f.trueTypeFont.Bounds(fixed.Int26_6(f.scale))
			gw = int32((gBnd.Max.X - gBnd.Min.X) >> 6)
			gh = int32((gBnd.Max.Y - gBnd.Min.Y) >> 6)

			// above can sometimes yield 0 for font smaller than 48pt, 1 is minimum
			if gw == 0 || gh == 0 {
				gw = 1
				gh = 1
			}
		}

		// The glyph's ascent and descent equal -bounds.Min.Y and +bounds.Max.Y.
		gAscent := int(-gBnd.Min.Y) >> 6
		gDescent := int(gBnd.Max.Y) >> 6

		char := &Character{
			width:    int(gw),
			height:   int(gh),
			advance:  int(gAdv),
			bearingV: gDescent,
			bearingH: int(gBnd.Min.X) >> 6,
		}

		//create image to draw glyph
		fg, bg := image.White, image.Black
		rect := image.Rect(0, 0, int(gw), int(gh))
		rgba := image.NewRGBA(rect)
		draw.Draw(rgba, rgba.Bounds(), bg, image.Point{}, draw.Src)

		//set the glyph dot
		px := 0 - (int(gBnd.Min.X) >> 6)
		py := (gAscent)
		pt := freetype.Pt(px, py)

		// Draw the text from mask to image
		c.SetClip(rgba.Bounds())
		c.SetDst(rgba)
		c.SetSrc(fg)
		_, err := c.DrawString(string(ch), pt)
		if err != nil {
			return err
		}

		// Generate texture
		var texture uint32
		gl.GenTextures(1, &texture)
		gl.BindTexture(gl.TEXTURE_2D, texture)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Dx()), int32(rgba.Rect.Dy()), 0,
			gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

		char.textureID = texture

		//add char to fontChar list
		f.Characters[ch] = char
	}

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return nil
}
