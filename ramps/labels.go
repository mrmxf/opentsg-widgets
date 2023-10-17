package ramps

import (
	"image"
	"image/draw"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/mmTristan/opentsg-core/colourgen"
	"github.com/mmTristan/opentsg-widgets/textbox"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func xPos(cWidth int, userFont font.Face, labelText, position string) int {

	boxw, _ := font.BoundString(userFont, labelText)
	width := boxw.Max.X.Ceil()
	switch strings.ToLower(position) {
	case "center":
		return (((cWidth) - width) / 2)
	case "right":
		return cWidth - width
	default: // Doesn't need a case as the schema only allows three inputs
		return 0
	}

}

func yPos(font font.Face, position string, stripeHeight int) int {
	// The two is for rounding errors
	yOffset := (font.Metrics().Ascent - font.Metrics().Descent - fixed.Int26_6(2)).Ceil()
	// YOffset := font.Ascent - font.Descent
	switch strings.ToLower(position) {
	case "top":
		return yOffset
	case "middle":
		mid := (stripeHeight + yOffset) / 2

		return mid
	default: // Doesn't need a case as the schema only allows three inputs
		return stripeHeight
	}
}

// labels places the label on the stripe based on the angle of the stripe, the text does not change angle
func (t *textObjectJSON) labels(target draw.Image, label, angle string) {

	var canvas draw.Image

	// get the flat row
	switch angle {
	case rotate180, noRotation:
		canvas = image.NewNRGBA64(target.Bounds())
	case rotate270, rotate90:
		canvas = image.NewNRGBA64(image.Rect(0, 0, target.Bounds().Dy(), target.Bounds().Dx()))
	}

	gradientBounds := canvas.Bounds().Max

	var stripeH int
	stripeH = gradientBounds.Y
	/*
		switch angle {
		case rotate180, noRotation:
			stripeH = gradientBounds.Y
		case rotate270, rotate90:
			stripeH = gradientBounds.X
		}*/
	lFont := fontGen(t.TextHeight, stripeH)

	// @TODO update so it always draws the same thing then put over the image as a overlay
	// so things are transposed we always draw the rows so they fit
	// 180 is inverse x and y
	// 90 is swap x and y
	// 270 is inverse and swap x and y

	col := colourgen.HexToColour(t.TextColour)
	b := canvas.Bounds().Max

	xpos := xPos(b.X, lFont, label, t.TextXPosition)
	ypos := yPos(lFont, t.TextYPosition, b.Y)

	point := fixed.Point26_6{X: fixed.Int26_6(xpos * 64), Y: fixed.Int26_6(ypos * 64)}

	/*
		// Assign the point based on the rotation to ensure the label lines up with the bar
		var point fixed.Point26_6
		switch angle {
		case rotate180, noRotation:
			xpos := xPos(b.X, lFont, label, t.TextXPosition)
			ypos := yPos(lFont, t.TextYPosition, stripeH)
			// else do not change the y
			point = fixed.Point26_6{X: fixed.Int26_6(xpos * 64), Y: fixed.Int26_6(ypos * 64)}
		case rotate270, rotate90:
			xpos := xPos(b.X, lFont, label, t.TextXPosition)
			ypos := yPos(lFont, t.TextYPosition, b.Y)

			point = fixed.Point26_6{X: fixed.Int26_6(xpos * 64), Y: fixed.Int26_6(ypos * 64)}
		}*/

	d := &font.Drawer{
		Dst:  canvas,
		Src:  image.NewUniform(col),
		Face: lFont,
		Dot:  point,
	}

	d.DrawString(label)

	// rotate the text and transpose it on
	// @TODO figure out how to make this more efficent

	var intermediate draw.Image
	intermediate = image.NewNRGBA64(target.Bounds())
	switch angle {
	case rotate90:
		for x := 0; x <= b.X; x++ {
			for y := 0; y <= b.Y; y++ {
				c := canvas.At(x, y)
				intermediate.Set(b.Y-y, x, c)
			}
		}
	case rotate270:
		for x := 0; x <= b.X; x++ {
			for y := 0; y <= b.Y; y++ {
				c := canvas.At(x, y)
				intermediate.Set(y, b.X-x, c)
			}
		}

	case rotate180:

		for x := 0; x <= b.X; x++ {
			for y := 0; y <= b.Y; y++ {
				c := canvas.At(x, y)
				intermediate.Set(b.X-x, b.Y-y, c)
			}
		}

	default:
		intermediate = canvas

	}
	// add the label
	draw.Draw(target, target.Bounds(), intermediate, image.Point{}, draw.Over)
}

// Fontgen returns a font a percentage height of the input
func fontGen(pixelPercent float64, height int) font.Face {
	var face font.Face
	var textHeight float64
	// Font is now out of a 100
	pixelheight := int((pixelPercent / 100.0) * float64(height))

	if pixelheight != 0 && pixelheight < height {
		textHeight = float64(pixelheight) / 0.7767
	} else {
		// Default height
		textHeight = float64(height) * 0.2 // * 0.005365
	}

	if textHeight < minTextHeightFont && !(minTextHeightPix > height) {
		textHeight = minTextHeightFont
	} else if height < minTextHeightPix {
		textHeight = 0 // And don't draw the text height
	}

	fontain, _ := truetype.Parse(textbox.Pixel)
	opt := truetype.Options{Size: textHeight}
	face = truetype.NewFace(fontain, &opt)

	return face
}

const ( // These are for the tiny font as it doesn't utilise all of the pixel space
	minTextHeightPix  = 7
	minTextHeightFont = 9
)
