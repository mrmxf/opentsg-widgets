// package luma generates the ebu3373 luma bar
package luma

import (
	"context"
	"image/color"
	"image/draw"
	"math"
	"sync"

	errhandle "github.com/mrmxf/opentsg-cote/errHandle"
	"github.com/mrmxf/opentsg-cote/widgethandler"
)

func Generate(canvasChan chan draw.Image, debug bool, c *context.Context, wg, wgc *sync.WaitGroup, logs *errhandle.Logger) {
	defer wg.Done()
	conf := widgethandler.GenConf[lumaJSON]{Debug: debug, Schema: schemaInit, WidgetType: "builtin.ebu3373/luma"}
	widgethandler.WidgetRunner(canvasChan, conf, c, logs, wgc)

}

func (l lumaJSON) Generate(canvas draw.Image, opt ...any) error {

	b := canvas.Bounds().Max
	// Set so the image scales at any size
	// Will need a caveat for widths < 1015 that the values will increase by more than one 10 bit pixel at a time
	wScale3 := 3.0 * (float64(b.X) / 3840)
	// Ceil the block width for floats and the extra pixel then goes to the start
	// This is following the design of the charts
	blockWidth := int(math.Ceil((float64(b.X) - wScale3*1015.0) / 2.0))

	for x := 0; x < b.X; x++ {
		var setColour color.NRGBA64
		// Check the x position and set the relevant colour
		switch {
		case x < blockWidth:
			setColour = color.NRGBA64{4096, 4096, 4096, 0xffff}
		case x >= (blockWidth + int(math.Ceil(wScale3*1015.0))):
			setColour = color.NRGBA64{46144, 46144, 46144, 0xffff}
		case x >= blockWidth && x < (blockWidth+int(math.Ceil(wScale3*1015.0))):
			// Calculate the changer per pixel and add to the base off 4
			pixVal := (float32(x-blockWidth) / float32(wScale3)) + 4.0
			// Floor the value and assign it as a 16 bit value
			// Aces.RGBA128{uint16(pixVal) << 6, uint16(pixVal) << 6, uint16(pixVal) << 6, 0xffff}
			setColour = color.NRGBA64{uint16(pixVal) << 6, uint16(pixVal) << 6, uint16(pixVal) << 6, 0xffff}
		}

		// Set for the same colour for the depth of the ramp
		for y := 0; y < b.Y; y++ {
			canvas.Set(x, y, setColour)
		}
	}

	return nil
}
