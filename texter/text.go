package texter

import (
	"context"
	"fmt"
	"image"
	"image/draw"
	"math"
	"os"

	_ "embed"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/mmTristan/opentsg-core/colourgen"
	"github.com/mrmxf/opentsg-core/config/core"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	FillTypeFull    = "full"
	FillTypeRelaxed = "relaxed"

	AlignmentLeft   = "left"
	AlignmentRight  = "right"
	AlignmentMiddle = "middle"
	AlignmentTop    = "top"
	AlignmentBottom = "bottom"

	FontTitle  = "title"
	FontBody   = "body"
	FontPixel  = "pixel"
	FontHeader = "header"
)

//go:embed MavenPro-Bold.ttf
var Title []byte

//go:embed MavenPro-Regular.ttf
var Header []byte

//go:embed Marvel-Regular.ttf
var Body []byte

//go:embed PixeloidSans.ttf
var Pixel []byte

/*

text is a struct that has a function that takes an image. And writes text of a certain colour on.

Background s not considered?

*/

// tsgContext is only required for when you are looking from the font from anywhere.
// a simple context.Background can be used if you're using an inbuilt font like text.
func (t TextboxJSON) DrawString(canvas draw.Image, tsgContext *context.Context, label string) error {
	return t.DrawStrings(canvas, tsgContext, []string{label})
}

func (t TextboxJSON) DrawStrings(canvas draw.Image, tsgContext *context.Context, labels []string) error {

	if t.Back != "" { //only draw the back if it's required
		backgroundcolor := colourgen.HexToColour(t.Back, t.ColourSpace)
		if backgroundcolor.A != 0 {
			draw.Draw(canvas, canvas.Bounds(), &image.Uniform{backgroundcolor}, image.Point{}, draw.Src)
		}
	}

	fontByte := FontSelector(tsgContext, t.Font)

	fontain, err := freetype.ParseFont(fontByte)
	if err != nil {
		return fmt.Errorf("0101 %v", err)
	}

	lines := len(labels)

	// scale the text to which ever dimension is smaller
	height := /*4 / 3.0 * (6.0 / 8.0)* */ (float64(canvas.Bounds().Max.Y) / float64(lines))
	width := /*4 / 3.0 * /*(6.0 / 8.0) * */ (float64(canvas.Bounds().Max.X))
	if width < height {
		height = width
	}

	opt := truetype.Options{Size: height, SubPixelsY: 8, Hinting: 2}
	myFace := truetype.NewFace(fontain, &opt)

	textCB := colourgen.HexToColour(t.Textc, t.ColourSpace)

	bounds := canvas.Bounds().Max
	bounds.Y /= lines

	// @TODO give th fonts a uniform option
	for i, label := range labels {

		var labelBox fixed.Rectangle26_6
		switch t.FillType {
		case FillTypeFull:
			myFace, labelBox = fullFill(bounds, myFace, fontain, height, label)
		default:
			myFace, labelBox = relaxedFill(bounds, myFace, fontain, height, label)
		}
		/*labelBox, _ := font.BoundString(myFace, label)
		textAreaX := float64(canvas.Bounds().Max.X)
		textAreaY := float64(canvas.Bounds().Max.Y) // Both side
		big := true

		// scale the text down to fix the box
		for big {

			thresholdX := float64(labelBox.Max.X.Round() + labelBox.Min.X.Round())
			thresholdY := float64(labelBox.Max.Y.Round() + labelBox.Min.Y.Round())
			fmt.Println(thresholdX, thresholdY, labelBox, label, height)
			// Compare the text width to the width of the text box
			if (thresholdX > textAreaX) || (thresholdY > textAreaY) {

				height *= 0.9
				opt = truetype.Options{Size: height, SubPixelsY: 8, Hinting: 2}
				myFace = truetype.NewFace(fontain, &opt)
				labelBox, _ = font.BoundString(myFace, label)

			} else {
				big = false
			}
		} */

		xOff := xPos(canvas, labelBox, t.XAlignment)
		yOff := yPos(canvas, labelBox, t.YAlignment, float64(lines), i)

		point := fixed.Point26_6{X: fixed.Int26_6(xOff * 64), Y: fixed.Int26_6(yOff * 64)}
		//	fmt.Println(xOff, point.X.Round())
		//	myFace := truetype.NewFace(fontain, &opt)
		d := &font.Drawer{
			Dst:  canvas,
			Src:  image.NewUniform(textCB),
			Face: myFace,
			Dot:  point,
		}
		d.DrawString(label)
	}

	return nil
}

func fullFill(area image.Point, sizeFont font.Face, fontain *truetype.Font, height float64, label string) (font.Face, fixed.Rectangle26_6) {
	labelBox, adv := font.BoundString(sizeFont, label)
	textAreaX := float64(area.X)
	textAreaY := float64(area.Y) // Both side
	big := true
	prevFont := sizeFont
	prevBox := labelBox
	// chnage the font when the initial bit is already too big
	if adv.Round() > int(textAreaX) || math.Abs(float64(labelBox.Max.Y.Round()))+math.Abs(float64(labelBox.Min.Y.Round())) > float64(textAreaY) {
		return relaxedFill(area, sizeFont, fontain, height, label)
	}

	// scale the text down to fix the box
	for big {
		// the base is always 0
		thresholdX := float64(labelBox.Max.X.Round()) //+ labelBox.Min.X.Round())
		thresholdY := math.Abs(float64(labelBox.Max.Y.Round())) + math.Abs(float64(labelBox.Min.Y.Round()))
		//fmt.Println(thresholdX, thresholdY, labelBox, label, height, textAreaX, textAreaY)
		// Compare the text width to the width of the text box
		if (thresholdX < textAreaX) && (thresholdY < textAreaY) {

			height *= 1.1
			opt := truetype.Options{Size: height, SubPixelsY: 8, Hinting: 2}
			prevFont = sizeFont
			prevBox = labelBox

			sizeFont = truetype.NewFace(fontain, &opt)
			//var adv fixed.Int26_6
			labelBox, _ = font.BoundString(sizeFont, label)
			//fmt.Println(adv.Round())

		} else {
			big = false
		}
	}
	//fmt.Println(prevBox)
	return prevFont, prevBox
}

func relaxedFill(area image.Point, sizeFont font.Face, fontain *truetype.Font, height float64, label string) (font.Face, fixed.Rectangle26_6) {
	labelBox, _ := font.BoundString(sizeFont, label)
	textAreaX := float64(area.X)
	textAreaY := float64(area.Y) // Both side
	big := true

	// scale the text down to fix the box
	for big {

		thresholdX := float64(labelBox.Max.X.Round()) // + labelBox.Min.X.Round())
		thresholdY := float64(labelBox.Max.Y.Round() - labelBox.Min.Y.Round())
		fmt.Println(thresholdX, thresholdY, labelBox, label, height)
		// Compare the text width to the width of the text box
		if (thresholdX > textAreaX) || (thresholdY > textAreaY) {

			height *= 0.9
			opt := truetype.Options{Size: height, SubPixelsY: 8, Hinting: 2}
			sizeFont = truetype.NewFace(fontain, &opt)
			labelBox, _ = font.BoundString(sizeFont, label)

		} else {
			big = false
		}
	}

	return sizeFont, labelBox
}

// place in the middle
func xPos(canvas image.Image, rect fixed.Rectangle26_6, position string) int {
	textWidth := rect.Max.X.Round() - rect.Min.X.Round()
	// textWidth := rect.Max.X.Ceil() - rect.Min.X.Ceil()
	// account for the minimum is where the text is started to be drawn

	switch position {
	case AlignmentLeft:
		return 0 - rect.Min.X.Round()
	case AlignmentRight:
		return canvas.Bounds().Max.X - textWidth
	default:
		return ((((canvas.Bounds().Max.X) - textWidth) / 2) - rect.Min.X.Round())
	}

}

// ypos calculates the yposition for the text
func yPos(canvas image.Image, rect fixed.Rectangle26_6, position string, lines float64, count int) int {
	yOffset := (rect.Max.Y.Round()) - (rect.Min.Y.Round())
	mid := (float64(canvas.Bounds().Max.Y) + float64(yOffset)) / (2.0 * lines)

	switch position {
	case AlignmentBottom:
		// fmt.Println((canvas.Bounds().Max.Y*count)/int(lines) - yOffset)
		return (canvas.Bounds().Max.Y*(count+1))/int(lines) - rect.Max.Y.Round()
	case AlignmentTop:
		return (canvas.Bounds().Max.Y*count)/int(lines) - rect.Min.Y.Round()
	default:
		fmt.Println(int(mid)+(canvas.Bounds().Max.Y*count)/int(lines), "Ypos", yOffset)
		return int(mid) + (canvas.Bounds().Max.Y*count)/int(lines)
	}

}

// font selector enumerates through the different sources of http,
// local files,
// then predetermined embedded fonts and returns the font based on the input string.
func FontSelector(c *context.Context, fontLocation string) []byte {

	font, err := core.GetWebBytes(c, fontLocation)

	if err == nil {
		return font
	}

	font, err = os.ReadFile(fontLocation)
	if err == nil {
		return font
	}

	switch fontLocation {
	case "title":
		return Title
	case "body":
		return Body
	case "pixel":
		return Pixel
	default:
		return Header
	}
}