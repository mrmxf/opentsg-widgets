// package saturation generates the ebu3373 saturation boxes
package saturation

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"sync"

	errhandle "github.com/mrmxf/opentsg-core/errHandle"
	"github.com/mrmxf/opentsg-core/widgethandler"
)

const (
	widgetType = "builtin.ebu3373/saturation"
)

func SatGen(canvasChan chan draw.Image, debug bool, c *context.Context, wg, wgc *sync.WaitGroup, logs *errhandle.Logger) {
	defer wg.Done()
	conf := widgethandler.GenConf[saturationJSON]{Debug: debug, Schema: schemaInit, WidgetType: widgetType}
	widgethandler.WidgetRunner(canvasChan, conf, c, logs, wgc) // Update this to pass an error which is then formatted afterwards

}

var (
	grey = color.NRGBA64{26496, 26496, 26496, 0xffff}
)

var (
	satRed0  = color.NRGBA64{793 << 6, 793 << 6, 793 << 6, 0xffff}
	satRed1  = color.NRGBA64{815 << 6, 779 << 6, 763 << 6, 0xffff}
	satRed2  = color.NRGBA64{834 << 6, 762 << 6, 730 << 6, 0xffff}
	satRed3  = color.NRGBA64{851 << 6, 742 << 6, 695 << 6, 0xffff}
	satRed4  = color.NRGBA64{866 << 6, 718 << 6, 654 << 6, 0xffff}
	satRed5  = color.NRGBA64{881 << 6, 689 << 6, 607 << 6, 0xffff}
	satRed6  = color.NRGBA64{894 << 6, 651 << 6, 549 << 6, 0xffff}
	satRed7  = color.NRGBA64{907 << 6, 600 << 6, 473 << 6, 0xffff}
	satRed8  = color.NRGBA64{919 << 6, 522 << 6, 386 << 6, 0xffff}
	satRed9  = color.NRGBA64{930 << 6, 390 << 6, 283 << 6, 0xffff}
	satRed10 = color.NRGBA64{940 << 6, 64 << 6, 64 << 6, 0xffff}
)

var (
	satGreen0  = color.NRGBA64{809 << 6, 809 << 6, 809 << 6, 0xffff}
	satGreen1  = color.NRGBA64{791 << 6, 830 << 6, 783 << 6, 0xffff}
	satGreen2  = color.NRGBA64{770 << 6, 848 << 6, 756 << 6, 0xffff}
	satGreen3  = color.NRGBA64{747 << 6, 863 << 6, 726 << 6, 0xffff}
	satGreen4  = color.NRGBA64{720 << 6, 878 << 6, 691 << 6, 0xffff}
	satGreen5  = color.NRGBA64{688 << 6, 890 << 6, 651 << 6, 0xffff}
	satGreen6  = color.NRGBA64{648 << 6, 902 << 6, 601 << 6, 0xffff}
	satGreen7  = color.NRGBA64{594 << 6, 913 << 6, 536 << 6, 0xffff}
	satGreen8  = color.NRGBA64{511 << 6, 922 << 6, 443 << 6, 0xffff}
	satGreen9  = color.NRGBA64{380 << 6, 931 << 6, 327 << 6, 0xffff}
	satGreen10 = color.NRGBA64{64 << 6, 940 << 6, 64 << 6, 0xffff}
)

var (
	satBlue0  = color.NRGBA64{579 << 6, 578 << 6, 578 << 6, 0xffff}
	satBlue1  = color.NRGBA64{552 << 6, 584 << 6, 631 << 6, 0xffff}
	satBlue2  = color.NRGBA64{520 << 6, 586 << 6, 677 << 6, 0xffff}
	satBlue3  = color.NRGBA64{484 << 6, 584 << 6, 717 << 6, 0xffff}
	satBlue4  = color.NRGBA64{445 << 6, 578 << 6, 755 << 6, 0xffff}
	satBlue5  = color.NRGBA64{403 << 6, 564 << 6, 790 << 6, 0xffff}
	satBlue6  = color.NRGBA64{360 << 6, 542 << 6, 823 << 6, 0xffff}
	satBlue7  = color.NRGBA64{312 << 6, 504 << 6, 854 << 6, 0xffff}
	satBlue8  = color.NRGBA64{260 << 6, 443 << 6, 884 << 6, 0xffff}
	satBlue9  = color.NRGBA64{197 << 6, 347 << 6, 912 << 6, 0xffff}
	satBlue10 = color.NRGBA64{64 << 6, 64 << 6, 940 << 6, 0xffff}
)

/*
set up config etc and generate the sturct
*/

func (s saturationJSON) Generate(canvas draw.Image, opt ...any) error {
	b := canvas.Bounds().Max

	// Use a map to match the keys up etc
	reds := []color.NRGBA64{satRed0, satRed1, satRed2, satRed3, satRed4, satRed5, satRed6, satRed7, satRed8, satRed9, satRed10}
	greens := []color.NRGBA64{satGreen0, satGreen1, satGreen2, satGreen3, satGreen4, satGreen5, satGreen6, satGreen7, satGreen8, satGreen9, satGreen10}
	blues := []color.NRGBA64{satBlue0, satBlue1, satBlue2, satBlue3, satBlue4, satBlue5, satBlue6, satBlue7, satBlue8, satBlue9, satBlue10}
	colours := make(map[string][]color.NRGBA64)
	cs := [][]color.NRGBA64{reds, greens, blues}
	names := []string{"red", "green", "blue"}
	for i, c := range cs {
		colours[names[i]] = c
	}
	var inputC []string
	// Assign the basic colour order if none are chosen
	if len(s.Colours) == 0 {
		inputC = names
	} else {
		inputC = s.Colours
	}

	wScale := (float64(b.X) / 2330.0)
	hScale := (float64(b.Y) / 600.0)

	height := 0.0
	// Scale the height offset to the number of colours called (max 3)
	hOff := hScale * 200 * (3.0 / float64(len(inputC)))

	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{grey}, image.Point{}, draw.Src)

	wOff := wScale * 200 // 200 is the width at 3840
	for _, cname := range inputC {
		colour := colours[cname]
		if colour == nil {
			// Blow the doors of etc
			return fmt.Errorf("TEST the colour %v is not an available colour", cname)
		}
		// Draw the narrow box
		width := wScale * 100
		draw.Draw(canvas, image.Rect(int(0), int(height), int(width), int(height+hOff)), &image.Uniform{colour[0]}, image.Point{}, draw.Over)
		// Draw the smaller boxes at a smaller fill
		for _, c := range colour[1:] {
			offx := 50 * wScale
			offy := 50 * hScale
			draw.Draw(canvas, image.Rect(int(width+offx), int(height+offy), int(width+wOff-offx), int(height+hOff-offy)), &image.Uniform{c}, image.Point{}, draw.Over)
			width += wOff
		}
		// Draw the final large box
		draw.Draw(canvas, image.Rect(int(width), int(height), b.X, int(height+hOff)), &image.Uniform{colour[len(colour)-1]}, image.Point{}, draw.Over)
		height += hOff
	}

	return nil
}
