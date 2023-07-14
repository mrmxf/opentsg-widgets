// Package textbox is used for adding textboxes onto the testcard
package textbox

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"os"
	"sync"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/mrmxf/opentsg-core/colourgen"
	"github.com/mrmxf/opentsg-core/config/core"
	errhandle "github.com/mrmxf/opentsg-core/errHandle"
	"github.com/mrmxf/opentsg-core/widgethandler"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	_ "embed"
)

//go:embed MavenPro-Bold.ttf
var Title []byte

//go:embed MavenPro-Regular.ttf
var Header []byte

//go:embed Marvel-Regular.ttf
var Body []byte

//go:embed PixeloidSans.ttf
var Pixel []byte

// TextBoxGen generates text boxes on a given image based on config values
func TBGenerate(canvasChan chan draw.Image, debug bool, c *context.Context, wg, wgc *sync.WaitGroup, logs *errhandle.Logger) {
	defer wg.Done()
	opts := []any{c}
	conf := widgethandler.GenConf[TextboxJSON]{Debug: debug, Schema: textBoxSchema, WidgetType: "builtin.textbox", ExtraOpt: opts}
	widgethandler.WidgetRunner(canvasChan, conf, c, logs, wgc) // Update this to pass an error which is then formatted afterwards
}

func (tb TextboxJSON) Generate(canvas draw.Image, opts ...any) error {

	// FontFile := conf.GetFont() // "./MavenPro-Bold.ttf" //"Marvel-Bold.ttf"
	// FontBytes, err := os.ReadFile(fontFile) // Insert font types to skip this bit
	if len(opts) != 1 {
		return fmt.Errorf("0102 textbox configuration error")
	}

	c, ok := opts[0].(*context.Context)
	if !ok {
		return fmt.Errorf("0103 configuration error when assiging textbox context")
	}

	fontByte := FontSelector(c, tb.Font)

	fontain, err := freetype.ParseFont(fontByte)
	if err != nil {
		return fmt.Errorf("0101 %v", err)
	}

	back := colourgen.HexToColour(tb.Back)
	backC := colourgen.ConvertNRGBA64(back)
	border := colourgen.HexToColour(tb.Border)
	borderC := colourgen.ConvertNRGBA64(border)

	borderSize := tb.BorderSize
	if borderC.A != 0 && backC.A != 0 { // only fill if there's colours to add
		borderBox(borderC, backC, borderSize, canvas) // generate a background filled box
	}

	labels := tb.Text
	lines := len(labels)

	height := (6.0 / 8.0) * (float64(canvas.Bounds().Max.Y) / float64(lines))
	width := (6.0 / 8.0) * (float64(canvas.Bounds().Max.X) / float64(lines))
	if width < height {
		height = width
	}
	opt := truetype.Options{Size: height, SubPixelsY: 8, Hinting: 2}
	myFace := truetype.NewFace(fontain, &opt)

	textCBase := colourgen.HexToColour(tb.Textc)
	textC := colourgen.ConvertNRGBA64(textCBase)

	for i, label := range labels {
		labelBox, _ := font.BoundString(myFace, label)
		textAreaX := float64(canvas.Bounds().Max.X) - float64(canvas.Bounds().Max.X)*borderSize*2 // 2 because the border is from
		textAreaY := float64(canvas.Bounds().Max.Y) - float64(canvas.Bounds().Max.Y)*borderSize*2 // Both sides
		big := true

		for big {

			thresholdX := float64(labelBox.Max.X.Round() + labelBox.Min.X.Round())
			thresholdY := float64(labelBox.Max.Y.Round() + labelBox.Min.Y.Round())
			// Comparre the text width to the width of the text box
			if (thresholdX > textAreaX) || (thresholdY > textAreaY) {

				height *= 0.9
				opt = truetype.Options{Size: height, SubPixelsY: 8, Hinting: 2}
				myFace = truetype.NewFace(fontain, &opt)
				labelBox, _ = font.BoundString(myFace, label)

			} else {
				big = false
			}
		}
		xOff := xPos(canvas, labelBox)
		yOff := yPos(canvas, labelBox, float64(lines), i)

		point := fixed.Point26_6{X: fixed.Int26_6(xOff * 64), Y: fixed.Int26_6(yOff * 64)}

		//	myFace := truetype.NewFace(fontain, &opt)
		d := &font.Drawer{
			Dst:  canvas,
			Src:  image.NewUniform(textC),
			Face: myFace,
			Dot:  point,
		}
		d.DrawString(label)
	}

	return nil
}

func borderBox(border, background color.Color, borderPercent float64, box draw.Image) {

	// Fill the whole thing as a border
	draw.Draw(box, box.Bounds(), &image.Uniform{border}, image.Point{}, draw.Src)

	bounds := box.Bounds().Max
	width, height := bounds.X, bounds.Y
	// This will be flexible at later dates

	borderw := int(math.Ceil(borderPercent * float64(height)))

	fills := image.Rect(0, 0, width-borderw*2, height-borderw*2)
	fill := image.NewNRGBA64(fills)
	draw.Draw(fill, fill.Bounds(), &image.Uniform{background}, image.Point{}, draw.Src)

	// Combine the elements
	//	draw.Draw(box, box.Bounds(), fill, image.Point{-borderw, -borderw}, draw.Src)
	//	fmt.Println(borderw)
	draw.Draw(box, image.Rect(borderw, borderw, box.Bounds().Max.X-borderw, box.Bounds().Max.Y-borderw), fill, image.Point{}, draw.Src)

}

// place in the middle
func xPos(canvas image.Image, rect fixed.Rectangle26_6) int {
	textWidth := rect.Max.X.Round() - rect.Min.X.Round()

	return (((canvas.Bounds().Max.X) - textWidth) / 2)

}

func yPos(canvas image.Image, rect fixed.Rectangle26_6, lines float64, count int) int {
	yOffset := (rect.Max.Y.Round()) - (rect.Min.Y.Round())
	mid := (float64(canvas.Bounds().Max.Y) + float64(yOffset)*0.8) / (2.0 * lines)

	return int(mid) + (canvas.Bounds().Max.Y*count)/int(lines)

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
