package ramps

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"strings"

	"github.com/mrmxf/opentsg-core/anglegen"
)

const (
	rotate180  = "rotate180"
	rotate90   = "rotate90 "
	rotate270  = "rotate270"
	noRotation = "xy"
)

/*

next steps
- implement different objects for interstripes etc
- get a global value variable set up, these are all inlcuded with the children as part of the configuration

*/

func shapes() {
	/*
	   shapes cuts it into a ramp longs

	   go run all the images then draw them on top

	   how to split the patterns and the inputs

	   rampsknows the images it wants to split to

	   and the json constituants as go structs?

	   make this translateable for a checkboard pattern

	   both split into segments thn call the segment, these require shape paramters and what the fills are

	   how to utilise the _hiddden builtin -
	*/
}

func firstrun(target draw.Image, input Ramp) error {
	// calculate the whole height of each one

	rotation, err := setBase(&input.WidgetProperties)

	if err != nil {
		return err
	}
	fmt.Println(input.WidgetProperties)

	// validate teh control here
	input.StripeGroup.InterStripe.base = input.WidgetProperties

	totalHeight := input.StripeGroup.Header.Height + ((len(input.StripeGroup.Ramp) - 1) * input.StripeGroup.InterStripe.Height)
	for _, r := range input.StripeGroup.Ramp {
		totalHeight += r.Height
	}

	totalHeight *= len(input.Stripes)
	rowDimens := input.WidgetProperties.rowDimension(target.Bounds())

	groupStep := float64(rowDimens) / float64(totalHeight)

	position := 0.0
	// posPoint := image.Point{}
	for _, str := range input.Stripes {

		if input.StripeGroup.Header.Height != 0 {

			rowHeight := input.StripeGroup.Header.Height
			// draw the header
			end := int(position + groupStep*float64(rowHeight))
			rowCut := input.WidgetProperties.rowOrColumn(target.Bounds(), end, position)
			row := image.NewNRGBA64(rowCut)
			posPoint := input.WidgetProperties.positionPoint(target.Bounds().Max, end-int(position), int(position))
			hidden(target, row, posPoint, input.StripeGroup.Header)

			position += groupStep * float64(rowHeight)

		}

		for i, ramp := range input.StripeGroup.Ramp {

			end := int(position + groupStep*float64(ramp.Height))
			rowCut := input.WidgetProperties.rowOrColumn(target.Bounds(), end, position)
			rrow := image.NewNRGBA64(rowCut)

			ramp.colour = str.Colour
			ramp.startPoint = str.StartPoint
			ramp.reverse = str.Reverse

			ramp.base = input.WidgetProperties
			posPoint := input.WidgetProperties.positionPoint(target.Bounds().Max, end-int(position), int(position))
			hidden(target, rrow, posPoint, ramp)

			position += groupStep * float64(ramp.Height)
			//	posPoint = input.base.positionPoint(target.Bounds().Max, int(position))
			if i+1 < len(input.StripeGroup.Ramp) {
				interHeight := input.StripeGroup.InterStripe.Height
				// accounts for jumps in floats and ints
				end := int(position + groupStep*float64(interHeight))

				rowCut := input.WidgetProperties.rowOrColumn(target.Bounds(), end, position)
				irow := image.NewNRGBA64(rowCut)
				altCopy := input.StripeGroup.InterStripe
				altCopy.base = input.WidgetProperties
				altCopy.step = input.StripeGroup.Ramp[i+1].BitDepth
				posPoint := input.WidgetProperties.positionPoint(target.Bounds().Max, end-int(position), int(position))
				hidden(target, irow, posPoint, altCopy)

				position += groupStep * float64(interHeight)
				//	posPoint = input.base.positionPoint(target.Bounds().Max, int(position))
				//calculate segments here
			}

		}

	}

	// @TODO add the rotation algortihim
	if rotation != 0 {
		rotate(target, rotation)
	}

	return nil
}

func setBase(target *control) (float64, error) {
	radian := 0.0
	target.angleType = noRotation

	if target.Angle != "" {
		angString := fmt.Sprintf("%v", target.Angle)
		var err error
		radian, err = anglegen.AngleCalc(angString)
		if err != nil {
			return 0, err
		}
	}

	angDiff, angleOffset := diff(radian, 1.571, 4.712, 3.142, 0.0)
	rads := fmt.Sprintf("%.3f", angDiff)

	switch rads {
	case "1.571":
		target.angleType = rotate90
	case "4.712":
		target.angleType = rotate270
	case "3.142":
		target.angleType = rotate180
	}

	return angleOffset, nil
}

func diff(angle float64, targets ...float64) (target float64, diff float64) {
	diff = math.Pi * 2
	target = 0.0
	for _, t := range targets {
		calcDiff := math.Abs(angle - t)
		if calcDiff < diff {
			diff = calcDiff
			target = t
		}
	}

	return
}

func (h internalHeader) Generate(img draw.Image) {
	if h.Height == 0 {
		return
	}

	c, _ := assignRGBValues(h.Colour, 65535, 0, 65535)
	draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.Point{}, draw.Over)
}

func (a gradientSeparator) Generate(img draw.Image) {

	if a.Height == 0 {
		return
	}

	step := int(math.Pow(2, float64((a.base.GlobalBitDepth - a.step))))

	altCount := 0
	end := a.base.getLoop(img.Bounds())

	for x := 0; x <= end; x += step {
		c, _ := assignRGBValues(a.Colours[altCount%len(a.Colours)], 65535, 0, 65535)
		target := a.base.set(x, step, img.Bounds().Max)

		draw.Draw(img, target, &image.Uniform{c}, image.Point{}, draw.Over)
		altCount++
	}

}

func (s Gradient) Generate(img draw.Image) {
	shift16 := 1 << (16 - s.BitDepth)

	//set the steps relative to the max bitdepth
	step := int(math.Pow(2, float64((s.base.GlobalBitDepth - s.BitDepth))))

	// generate a start point in 16 bit
	// sanity check the start point is within the bitdepth
	startPoint := s.startPoint << (16 - s.base.GlobalBitDepth)

	overRun := startPoint % shift16

	if overRun != 0 {
		/*
					tow options line up when we can let them lineup when required
			 		or have them always lineup by shifting the value to match the lowest bit closest bitdepth
					e.g. would shift to 0?@TODO fix eveything else before tackling this

		*/
		if !s.reverse {
			startPoint += (shift16 - overRun)
		} else {
			startPoint -= overRun
		}
	}

	altCount := 0
	// have the mover position and the bounds
	end := s.base.getLoop(img.Bounds())

	for x := 0; x <= end; x += step {
		c, _ := assignRGBValues(s.colour, float64(startPoint), 0, 65535)
		target := s.base.set(x, step, img.Bounds().Max)

		// draw.Draw(img, image.Rect(x, img.Bounds().Min.Y, x+step, img.Bounds().Max.Y), &image.Uniform{c}, image.Point{}, draw.Over)

		draw.Draw(img, target, &image.Uniform{c}, image.Point{}, draw.Over)
		altCount++

		//make the colour steps 16 bit
		if !s.reverse {
			startPoint += shift16 // 1 << shift16
		} else {
			startPoint -= shift16 // 1 << shift16
		}
	}

	// generate the label if needed
	if s.Label != "" {
		s.base.TextProperties.labels(img, s.Label, s.base.angleType)
	}

	//run the labels here - use the other label code
}

type control struct {
	GlobalBitDepth int
	Angle          string
	TextProperties textObjectJSON
	// These are things the user does not set
	/*
		fill function - for rotation to automatically translate the fill location
		fill - get stepsize and end goal

		step size - fill or truncate. Add a multiplier


	*/

	angleType string
}

func (c control) getLoop(bounds image.Rectangle) (end int) {

	if c.angleType == noRotation || c.angleType == rotate180 {
		end = bounds.Max.X
	} else {
		end = bounds.Max.Y
	}
	return
}

func (c control) rowDimension(bounds image.Rectangle) (end int) {

	if c.angleType == noRotation || c.angleType == rotate180 {
		end = bounds.Max.Y
	} else {
		end = bounds.Max.X
	}
	return
}

func (c control) rowOrColumn(bounds image.Rectangle, end int, position float64) image.Rectangle {
	if c.angleType == noRotation || c.angleType == rotate180 {
		return image.Rect(0, 0, bounds.Dx(), end-int(position))
	} else {
		return image.Rect(0, 0, end-int(position), bounds.Dy())
	}

}

func (c control) positionPoint(bounds image.Point, rowSize, shift int) image.Point {

	switch c.angleType {
	case noRotation:
		return image.Point{Y: shift}
	case rotate180:
		return image.Point{Y: bounds.Y - shift - rowSize}
	case rotate90:
		return image.Point{X: bounds.X - shift - rowSize}
	default: // rotate270
		return image.Point{X: shift}
	}

}

// func set sets the canvas values based on the roatation without running a transformation
func (c control) set(position, step int, bounds image.Point) image.Rectangle {

	switch c.angleType {
	case noRotation:
		return image.Rect(position, 0, position+step, bounds.Y)
	case rotate180:
		return image.Rect(bounds.X-(position), 0, bounds.X-(position+step), bounds.Y)
	case rotate90:

		return image.Rect(0, position, bounds.X, position+step)
	case rotate270:

		return image.Rect(0, bounds.Y-(position), bounds.X, bounds.Y-(position+step))
	}
	//case rotate270:

	//		canvas.Set(b.X-int(j), b.Y-(int(i)+1), colourRGB)
	//default:
	//	canvas.Set(int(j), int(i), colourRGB)
	//}

	return image.Rectangle{}
}

type maker interface {
	Generate(img draw.Image)
}

// Defaults give the optional extras?
func hidden(base, img draw.Image, start image.Point, G maker) {

	/*
		hidden needs to be something that can be generic and useful

	*/
	G.Generate(img) //add optional parameterss?

	draw.Draw(base, img.Bounds().Add(start), img, image.Point{}, draw.Over)
}

func assignRGBValues(colour string, rgb float64, maxBlack, maxWhite uint16) (color.NRGBA64, error) {
	switch strings.ToLower(colour) {
	case "grey", "gray": // "black", "white",
		return color.NRGBA64{uint16(rgb), uint16(rgb), uint16(rgb), uint16(0xffff)}, nil
	case "black":
		return color.NRGBA64{maxBlack, maxBlack, maxBlack, uint16(0xffff)}, nil
	case "white":
		return color.NRGBA64{maxWhite, maxWhite, maxWhite, uint16(0xffff)}, nil
	case "red":
		return color.NRGBA64{uint16(rgb), 0, 0, uint16(0xffff)}, nil
	case "green":
		return color.NRGBA64{0, uint16(rgb), 0, uint16(0xffff)}, nil
	case "blue":
		return color.NRGBA64{0, 0, uint16(rgb), uint16(0xffff)}, nil
	default:
		return color.NRGBA64{0, 0, 0, 0}, fmt.Errorf("%s Non specific colour called, rgb values set at 0", colour) // Unused error
	}
}

func rotate(canvas draw.Image, radian float64) {

	// Take n as 10 for the moment
	// Math.ceil x, y and floor each one

	size := canvas.Bounds().Max
	x0, y0 := float64((size.X / 2)), float64((size.Y / 2))
	// Calculate these on initialisation
	// Use base as a method of calculating it all without changing the canvas
	base := image.NewNRGBA64(canvas.Bounds())
	draw.Draw(base, base.Bounds(), canvas, image.Point{}, draw.Src)
	N := int(10)

	rgbs := make([][]uint32, 4)
	for i := range rgbs {
		rgbs[i] = make([]uint32, 4)
	}

	val := make([]uint16, 4)

	for i := 0; i < base.Bounds().Max.X; i++ {
		for j := 0; j < base.Bounds().Max.Y; j++ {
			// Calculate the pixel location to extract from
			xp := math.Cos(-radian)*(float64(i)-x0) + math.Sin(-radian)*(float64(j)-y0) + x0
			yp := -1*math.Sin(-radian)*(float64(i)-x0) + math.Cos(-radian)*(float64(j)-y0) + y0
			_, xFrac := math.Modf(xp)
			_, yFrac := math.Modf(yp)
			x := int(xFrac * 10)
			y := int(yFrac * 10)

			xpos, ypos := int(math.Floor(xp)), int(math.Floor(yp))
			locs := [][]int{{xpos, ypos}, {xpos + 1, ypos}, {xpos, ypos + 1}, {xpos + 1, ypos + 1}}
			// Overwrite the rgb values each time instead of making a new array for each loop
			for i, loc := range locs {
				rgbs[i][0], rgbs[i][1], rgbs[i][2], rgbs[i][3] = base.At(loc[0], loc[1]).RGBA()
			}
			for k := 0; k < 4; k++ {
				val[k] = uint16((1.0 / (float64(N * N))) * float64(((N-x)*(N-y)*int(rgbs[0][k]) + x*(N-y)*int(rgbs[1][k]) + y*(N-x)*int(rgbs[2][k]) + x*y*int(rgbs[3][k]))))
			}
			// If not empty then assign the value to ignore the black background
			if val[3] != 0 {
				canvas.Set(i, j, color.NRGBA64{val[0], val[1], val[2], uint16(0xffff)})
			} else {
				canvas.Set(i, j, color.NRGBA64{0, 0, 0, uint16(0x0000)})
			}
		}
	}
}
