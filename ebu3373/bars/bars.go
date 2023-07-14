// Package bars generates the ebu colour bars
package bars

import (
	"context"
	"image"
	"image/color"
	"image/draw"
	"math"
	"sync"

	errhandle "github.com/mrmxf/opentsg-core/errHandle"
	"github.com/mrmxf/opentsg-core/widgethandler"
)

func BarGen(canvasChan chan draw.Image, debug bool, c *context.Context, wg, wgc *sync.WaitGroup, logs *errhandle.Logger) {
	defer wg.Done()
	conf := widgethandler.GenConf[barJSON]{Debug: debug, Schema: schemaInit, WidgetType: "builtin.ebu3373/bars"}
	widgethandler.WidgetRunner(canvasChan, conf, c, logs, wgc) // Update this to pass an error which is then formatted afterwards

}

// these colours should all be relative to BT 2100 (rec 2020)
// they were given as 10 bit but moved to 16 bit for running with go
var (
	white   = color.NRGBA64{60160, 60160, 60160, 0xffff}
	yellow  = color.NRGBA64{60160, 60160, 4096, 0xffff}
	cyan    = color.NRGBA64{4096, 60160, 60160, 0xffff}
	green   = color.NRGBA64{4096, 60160, 4096, 0xffff}
	magenta = color.NRGBA64{60160, 4096, 60160, 0xffff}
	red     = color.NRGBA64{60160, 4096, 4096, 0xffff}
	blue    = color.NRGBA64{4096, 4096, 60160, 0xffff}
	grey    = color.NRGBA64{26496, 26496, 26496, 0xffff}
)

var (
	white40   = color.NRGBA64{46144, 46144, 46144, 0xffff}
	yellow40  = color.NRGBA64{46144, 46144, 4096, 0xffff}
	cyan40    = color.NRGBA64{4096, 46144, 46144, 0xffff}
	green40   = color.NRGBA64{4096, 46144, 4096, 0xffff}
	magenta40 = color.NRGBA64{46144, 4096, 46144, 0xffff}
	red40     = color.NRGBA64{46144, 4096, 4096, 0xffff}
	blue40    = color.NRGBA64{4096, 4096, 46144, 0xffff}
)
var (
	dLWhite   = color.NRGBA64{38528, 38528, 38528, 0xffff}
	dLYellow  = color.NRGBA64{594 << 6, 601 << 6, 246 << 6, 0xffff}
	dLCyan    = color.NRGBA64{408 << 6, 591 << 6, 601 << 6, 0xffff}
	dLGreen   = color.NRGBA64{388 << 6, 589 << 6, 232 << 6, 0xffff}
	dLMagenta = color.NRGBA64{534 << 6, 227 << 6, 595 << 6, 0xffff}
	dLRed     = color.NRGBA64{522 << 6, 216 << 6, 138 << 6, 0xffff}
	dLBlue    = color.NRGBA64{187 << 6, 127 << 6, 602 << 6, 0xffff}
)

var (
	sLWhite   = color.NRGBA64{39552, 39552, 39552, 0xffff}
	sLYellow  = color.NRGBA64{610 << 6, 616 << 6, 253 << 6, 0xffff}
	sLCyan    = color.NRGBA64{422 << 6, 605 << 6, 615 << 6, 0xffff}
	sLGreen   = color.NRGBA64{400 << 6, 603 << 6, 238 << 6, 0xffff}
	sLMagenta = color.NRGBA64{541 << 6, 230 << 6, 601 << 6, 0xffff}
	sLRed     = color.NRGBA64{527 << 6, 218 << 6, 139 << 6, 0xffff}
	sLBlue    = color.NRGBA64{186 << 6, 126 << 6, 598 << 6, 0xffff}
)

func (bar barJSON) Generate(canvas draw.Image, opt ...any) error {
	b := canvas.Bounds().Max

	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{grey}, image.Point{}, draw.Src)
	wScale := (float64(b.X) / 3840.0)
	barWidth := wScale * 412

	heights := []float64{200, 560, 200, 200}
	cbars := [][]color.NRGBA64{
		{white, yellow, cyan, green, magenta, red, blue},
		{white40, yellow40, cyan40, green40, magenta40, red40, blue40},
		{dLWhite, dLYellow, dLCyan, dLGreen, dLMagenta, dLRed, dLBlue},
		{sLWhite, sLYellow, sLCyan, sLGreen, sLMagenta, sLRed, sLBlue},
	}
	hOff := 0.0
	hScale := (float64(b.Y) / 1160.0)
	for i, h := range heights {

		barHeight := hScale * h

		boxHeight := hOff + barHeight

		//accounting for rounding errors at the end of the bars
		if math.Abs(float64(b.Y)-boxHeight) < 0.002 {
			boxHeight = float64(b.Y)
		}

		off := wScale * 480
		for _, c := range cbars[i] {
			area := image.Rect(int(off), int(hOff), int(off+barWidth), int(boxHeight))
			draw.Draw(canvas, area, &image.Uniform{c}, image.Point{}, draw.Over)
			off += barWidth
		}

		hOff = boxHeight
	}

	return nil
}
