// Package framecount adds a framecounter to a user specified location
package framecount

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/mrmxf/opentsg-core/colourgen"
	errhandle "github.com/mrmxf/opentsg-core/errHandle"
	"github.com/mrmxf/opentsg-core/widgethandler"
	"github.com/mrmxf/opentsg-widgets/textbox"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func CountGen(canvasChan chan draw.Image, debug bool, c *context.Context, wg, wgc *sync.WaitGroup, logs *errhandle.Logger) {
	defer wg.Done()
	opts := []any{c}
	conf := widgethandler.GenConf[frameJSON]{Debug: debug, Schema: frameSchema, WidgetType: "builtin.framecounter", ExtraOpt: opts}
	widgethandler.WidgetRunner(canvasChan, conf, c, logs, wgc) // Update this to pass an error which is then formatted afterwards
}

var pos = framePos

func (f frameJSON) Helper(key string, c *context.Context) {
	// Update the frame number add amend context with it
	f.FrameNumber = pos()
	fc := make(map[string]frameJSON)
	fc[key] = f

	// Widgethandler.Put(fc, c)

}

func (f frameJSON) Generate(canvas draw.Image, extraOpts ...any) error {

	b := canvas.Bounds().Max
	if !f.getFrames() {
		return nil
	}

	if f.Font == "" {
		f.Font = "pixel"
	}

	if len(extraOpts) != 1 {
		return fmt.Errorf("0154 framecount configuration error")
	}

	c, ok := extraOpts[0].(*context.Context)
	if !ok {
		return fmt.Errorf("0155 configuration error when assiging framecount context")
	}

	fontByte := textbox.FontSelector(c, f.Font)
	fontain, err := freetype.ParseFont(fontByte)

	if err != nil {
		return fmt.Errorf("0152 %v", err)
	}
	// Size of the text in pixels to font
	f.FontSize = (float64(b.Y) * (f.FontSize / 100)) / 0.75 // Convert from pixels to points
	if b.Y > b.X {
		f.FontSize *= (float64(b.X) / float64(b.Y)) // Scale the font size for narrow grids
	}

	opt := truetype.Options{Size: f.FontSize, SubPixelsY: 8, Hinting: 2}
	myFace := truetype.NewFace(fontain, &opt)

	// MyFont.Advance
	mes, err := intTo4(pos())
	if err != nil {
		return err
	}
	// Get the width of 0
	width, _ := myFace.GlyphAdvance('0')
	height := (width.Ceil()) * len(mes)
	// Keep it square with +1 for tolerance
	square := image.Point{height + 1, height + 1}

	frame := image.NewNRGBA64(image.Rect(0, 0, square.X, square.Y))
	background := userColour(f.BackColour, color.NRGBA64{uint16(195) << 8, uint16(195) << 8, uint16(195) << 8, uint16(195) << 8})
	// Generate a semi transparent grey background
	for i := 0; i < frame.Bounds().Max.Y; i++ {
		for j := 0; j < frame.Bounds().Max.X; j++ {
			frame.SetNRGBA64(j, i, background)
		}
	}

	text := userColour(f.TextColour, color.NRGBA64{0, 0, 0, 65535})
	yOff := (float64(square.Y) / 29) * 5 // This constant is to place the y at the text at the center of the square for each height
	point := fixed.Point26_6{X: fixed.Int26_6(1 * 64), Y: fixed.Int26_6(((float64(height) / 2) + yOff) * 64)}
	d := &font.Drawer{
		Dst:  frame,
		Src:  image.NewUniform(image.NewUniform(text)),
		Face: myFace,
		Dot:  point,
	}

	d.DrawString(mes)

	fb := frame.Bounds().Max
	// If pos not given then draw it here

	var x, y int
	switch imgpos := f.Imgpos.(type) {
	case map[string]interface{}:
		x, y = userPos(imgpos, b, fb)
	default:
		x, y = 0, 0

	}

	if x > (b.X - fb.X) {
		return fmt.Errorf("_0153 the x position %v is greater than the x boundary of %v with frame width of %v", x, canvas.Bounds().Max.X, fb.X)
	} else if y > b.Y-fb.Y {
		return fmt.Errorf("_0153 the y position %v is greater than the y boundary of %v with frame height of %v", y, canvas.Bounds().Max.Y, fb.Y)
	}

	// Corner := image.Point{-1 * (canvas.Bounds().Max.X - height - 1), -1 * (canvas.Bounds().Max.Y - height - 1)}
	draw.Draw(canvas, canvas.Bounds(), frame, image.Point{-x, -y}, draw.Over)

	return nil
}

func userColour(input string, defaultC color.NRGBA64) color.NRGBA64 {
	var gen color.NRGBA64

	if input == "" {
		gen = defaultC
	} else {
		inter := colourgen.HexToColour(input)
		gen = colourgen.ConvertNRGBA64(inter)
	}

	return gen
}

func intTo4(num int) (string, error) {
	s := strconv.Itoa(num)
	if len(s) > 4 {
		return "", fmt.Errorf("frame Count greater then 9999")
	}

	buf0 := strings.Repeat("0", 4-len(s))

	s = buf0 + s

	return s, nil
}

func userPos(location map[string]interface{}, canSize, frameSize image.Point) (int, int) {
	if location["alias"] != nil {
		// Process as simple location
		// The minus one is inluded to compensate for canvas startnig at 0
		switch location["alias"].(string) {
		case "bottom left":
			return 0, canSize.Y - frameSize.Y - 1
		case "bottom right":
			return canSize.X - frameSize.X - 1, canSize.Y - frameSize.Y - 1
		case "top right":
			return canSize.X - frameSize.X - 1, 0
		default:
			return 0, 0
		}
	} else {
		var x, y int
		if mid := location["x"]; mid != nil { // Make a percentage of the canvas
			var percent float64
			switch val := mid.(type) {
			case float64:
				percent = val
			case int:
				percent = float64(val)
			}
			x = int(math.Floor(percent * (float64(canSize.X) / 100)))
		}
		if mid := location["y"]; mid != nil {
			var percent float64
			switch val := mid.(type) {
			case float64:
				percent = val
			case int:
				percent = float64(val)
			}
			y = int(math.Floor(percent * (float64(canSize.Y) / 100)))
		}

		return x, y
	}
}
