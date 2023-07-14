// package nearblack generates the ebu3373 nearblack bar
package nearblack

import (
	"context"
	"image"
	"image/color"
	"image/draw"
	"sync"

	errhandle "github.com/mrmxf/opentsg-core/errHandle"
	"github.com/mrmxf/opentsg-core/widgethandler"
)

func NBGenerate(canvasChan chan draw.Image, debug bool, c *context.Context, wg, wgc *sync.WaitGroup, logs *errhandle.Logger) {
	defer wg.Done()
	conf := widgethandler.GenConf[nearblackJSON]{Debug: debug, Schema: schemaInit, WidgetType: "builtin.ebu3373/nearblack"}
	widgethandler.WidgetRunner(canvasChan, conf, c, logs, wgc)

}

var (
	neg4  = color.NRGBA64{2048, 2048, 2048, 0xffff}
	neg2  = color.NRGBA64{3072, 3072, 3072, 0xffff}
	neg1  = color.NRGBA64{3584, 3584, 3584, 0xffff}
	black = color.NRGBA64{4096, 4096, 4096, 0xffff}
	pos1  = color.NRGBA64{4608, 4608, 4608, 0xffff}
	pos2  = color.NRGBA64{5120, 5120, 5120, 0xffff}
	pos4  = color.NRGBA64{6144, 6144, 6144, 0xffff}

	grey = color.NRGBA64{26496, 26496, 26496, 0xffff}
)

func (nb nearblackJSON) Generate(canvas draw.Image, opt ...any) error {

	b := canvas.Bounds().Max
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{grey}, image.Point{}, draw.Src)
	// Scale everything so it fits the shape of the canvas
	wScale := (float64(b.X) / 3840.0)
	startPoint := wScale * 480
	off := wScale * 206

	order := []color.NRGBA64{neg4, neg2, neg1, pos1, pos2, pos4}
	area := image.Rect(int(startPoint), 0, int(startPoint+off*2), b.Y)
	draw.Draw(canvas, area, &image.Uniform{black}, image.Point{}, draw.Src)
	startPoint += off * 2
	for _, c := range order {
		// alternate through the colours
		draw.Draw(canvas, image.Rect(int(startPoint), 0, int(startPoint+off), b.Y), &image.Uniform{c}, image.Point{}, draw.Src)
		startPoint += off
		// append with the 0% black
		draw.Draw(canvas, image.Rect(int(startPoint), 0, int(startPoint+off), b.Y), &image.Uniform{black}, image.Point{}, draw.Src)
		startPoint += off
	}

	return nil
}
