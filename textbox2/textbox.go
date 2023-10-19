package textbox2

import (
	"context"
	"fmt"
	"image"
	"image/draw"

	"github.com/mmTristan/opentsg-core/colour"
	"github.com/mmTristan/opentsg-core/colourgen"
	"github.com/mmTristan/opentsg-widgets/texter"
	"github.com/mrmxf/opentsg-core/config"
)

/*
textbox 2 has the border on the outside
then the textbox in the middle

*/

type TextboxJSON struct {
	// Type       string       `json:"type" yaml:"type"`
	Text           []string `json:"text" yaml:"text"`
	TextProperties texter.TextboxJSON
	GridLoc        *config.Grid      `json:"grid" yaml:"grid"`
	ColourSpace    colour.ColorSpace `json:"ColorSpace,omitempty" yaml:"ColorSpace,omitempty"`
	Border         string            `json:"bordercolor" yaml:"bordercolor"`
	BorderSize     float64           `json:"bordersize" yaml:"bordersize"`
}

func (tb TextboxJSON) Generate(canvas draw.Image, opts ...any) {
	// calculate the border here
	/*
	 take the percentage of whatevers thinner?


	*/
	bounds := canvas.Bounds().Max
	width, height := (float64(bounds.X)*tb.BorderSize)/100, (float64(bounds.Y)*tb.BorderSize)/100

	borderwidth := int(height)
	if width < height {
		borderwidth = int(width)
	}

	borderColour := colourgen.HexToColour(tb.Border, tb.ColourSpace)

	draw.Draw(canvas, image.Rect(0, 0, borderwidth, canvas.Bounds().Max.Y), &image.Uniform{borderColour}, image.Point{}, draw.Src)
	draw.Draw(canvas, image.Rect(0, 0, canvas.Bounds().Max.X, borderwidth), &image.Uniform{borderColour}, image.Point{}, draw.Src)
	draw.Draw(canvas, image.Rect(canvas.Bounds().Max.X-borderwidth, 0, canvas.Bounds().Max.X, canvas.Bounds().Max.Y), &image.Uniform{borderColour}, image.Point{}, draw.Src)
	draw.Draw(canvas, image.Rect(0, canvas.Bounds().Max.Y-borderwidth, canvas.Bounds().Max.X, canvas.Bounds().Max.Y), &image.Uniform{borderColour}, image.Point{}, draw.Src)
	//get the colour
	c := colour.NewNRGBA64(tb.ColourSpace, image.Rect(0, 0, canvas.Bounds().Max.X-borderwidth*2, canvas.Bounds().Max.Y-borderwidth*2))
	cb := context.Background()
	tb.TextProperties.DrawStrings(c, &cb, tb.Text)
	fmt.Println(c.Bounds())
	draw.Draw(canvas, image.Rect(borderwidth, borderwidth, canvas.Bounds().Max.X-borderwidth, canvas.Bounds().Max.Y-borderwidth), c, image.Point{}, draw.Src)
	// colourgen.HexToColour(tb.Border, tb.ColourSpace)
	/*

		text section - generate a seperate box
		run that as the text section

	*/
}
