// Package noise generates images of noise
package noise

import (
	"context"
	"fmt"
	"image/color"
	"image/draw"
	"math/rand"
	"sync"
	"time"

	errhandle "github.com/mrmxf/opentsg-cote/errHandle"
	"github.com/mrmxf/opentsg-cote/widgethandler"
)

// NGenerator generates images of noise
func NGenerator(canvasChan chan draw.Image, debug bool, c *context.Context, wg, wgc *sync.WaitGroup, logs *errhandle.Logger) {
	defer wg.Done()
	conf := widgethandler.GenConf[noiseJSON]{Debug: debug, Schema: schemaInit, WidgetType: "builtin.noise"}
	widgethandler.WidgetRunner(canvasChan, conf, c, logs, wgc) // Update this to pass an error which is then formatted afterwards
}

var randnum = randSeed

func randSeed() int64 {
	return time.Now().Unix()
}

func (n noiseJSON) Generate(canvas draw.Image, opt ...any) error {
	// Have a seed variable tht is taken out for testing purposes
	rand.Seed(randnum())

	var max int
	if n.Maximum != 0 {
		max = n.Maximum
	} else {
		// Revert to the default
		max = 4095
	}
	min := n.Minimum

	if max < min {
		return fmt.Errorf("0141 The minimum noise value %v is greater than the maximum noise value %v", min, max)
	}

	if n.Noisetype == "white noise" { // upgrade to switch statement when more types come in
		whitenoise(canvas, min, max)
	}

	return nil
}

func whitenoise(canvas draw.Image, min, max int) {
	b := canvas.Bounds().Max
	for y := 0; y < b.Y; y++ {
		for x := 0; x < b.X; x++ {
			colourPos := uint16(rand.Intn(max-min)+min) << 4
			// Fill := color.NRGBA64{colourPos << 4, colourPos << 4, colourPos << 4, uint16(0xffff)}

			canvas.Set(x, y, color.NRGBA64{colourPos, colourPos, colourPos, 0xffff})
		}
	}
}
