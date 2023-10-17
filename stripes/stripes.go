// Package stripes generates the bitdepth ramps
package stripes

import (
	"context"
	"sort"

	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"strings"
	"sync"

	"github.com/mmTristan/opentsg-core/anglegen"
	errhandle "github.com/mmTristan/opentsg-core/errHandle"
	"github.com/mmTristan/opentsg-core/widgethandler"
)

const (
	widgetType = "builtin.ramps"
)

// RampGen generates images of ramps at any specified angle
func RampGen(canvasChan chan draw.Image, debug bool, c *context.Context, wg, wgc *sync.WaitGroup, logs *errhandle.Logger) {
	defer wg.Done()
	conf := widgethandler.GenConf[rampJSON]{Debug: debug, Schema: schemaInit, WidgetType: widgetType}
	widgethandler.WidgetRunner(canvasChan, conf, c, logs, wgc) // Update this to pass an error which is then formatted afterwards
}

func (r rampJSON) Generate(canvas draw.Image, opts ...any) error {
	// Generate the variables for ramp
	if err := r.constantInit(); err != nil {
		return err
	}

	stripeFills := groupFill(*r.Stripes)

	centre := canvas.Bounds().Max
	radian := 0.0
	if r.Angle != "" {
		angString := fmt.Sprintf("%v", r.Angle)
		var err error
		radian, err = anglegen.AngleCalc(angString)
		if err != nil {
			return err
		}
	}

	width := centre.X
	height := centre.Y
	ang := noRotation
	inverse := 0
	// Reduce all angle to be within 2pi
	if mult := (radian / (2 * math.Pi)); mult > 1 {
		radian -= (2 * math.Pi * float64(int(mult))) // Round down to below two pi
	}
	// Assign to the closest angle
	angDiff, radDiff := diff(radian, 1.571, 4.712, 3.142, 0.0)
	rads := fmt.Sprintf("%.3f", angDiff)

	if rads == "1.571" || rads == "4.712" { // Assign an angle based on the direction
		width = centre.Y // This goes down the y direction now
		height = centre.X
		if rads == "1.571" {
			ang = rotate90
		} else {
			ang = rotate270
			inverse = centre.X
		}
	} else if rads == "3.142" {
		ang = rotate180
		inverse = centre.Y
	}

	scale := 1.0
	if strings.ToLower(r.FillType) == "fill" {
		scale = float64(width) / 4096.0 // Assign based on the new width
	}
	if r.Stripes.Stripes.GroupTypes == nil {
		return fmt.Errorf("0123 No grouptypes have been supplied in the json, no stripes can be computed")
	}
	groups := r.Stripes.Stripes.GroupTypes

	// Extract the keys of the colours and organise them in alphabetical order so the ramps are repeatable
	pos := 0
	keys := make([]string, len(groups))
	for key := range groups {
		keys[pos] = key
		pos++
	}
	sort.Strings(keys)
	// Offset is the position in the image for each ramp to be generated at
	offset := 0.0
	// Fill each type of fill out

	for _, key := range keys {
		group := groups[key]
		// Generate the fill for each group
		fill := fillUpdate(stripeFills, group, r.Depth)
		fill, err := ratioToHeight(fill, height, len(groups))
		if err != nil {
			return err
		}
		for fillPos, f := range fill {
			switch f.gradient {
			case "header":
				// MAKE A Header function
				r.headerGen(canvas, fill, fillPos, width, offset, scale, ang)

			default:
				// Calculate the gradient and then fill the depth
				r.stripeGen(canvas, fill[fillPos], r.Depth, width, inverse, offset, scale, ang)

			}
			offset += f.height
		}
	}

	if radDiff > 0.001 {
		rotate(canvas, radDiff)
	}

	return nil

}

// headerGen generates the group header and divider lines
func (r rampJSON) headerGen(canvas draw.Image, fill []fill, fillPos, width int, offset, scale float64, ang string) {
	f := fill[fillPos]
	shift := 1
	if fillPos+1 < len(fill) {
		shift = int(math.Pow(2, float64(12-fill[fillPos+1].depth)))
	}
	blockc := f.colour[0]

	for i := 0; i < width; i++ {
		colourRGB, _ := assignRGBValues(blockc, 4095, uint16(r.Minimum), uint16(r.Maximum))
		// Assign the value for the depth of the bar
		for j := float64(offset); j < float64(offset+f.height); j++ {
			set(ang, canvas, colourRGB, float64(i), j)
		}
		if s := math.Mod(float64((i + 1)), float64(shift)*scale); int(s) == 0 {
			blockc = colourSwap(f.colour, blockc)
		}
	}
}

// stripeGen generates the gradient lines for the stripes
func (r rampJSON) stripeGen(canvas draw.Image, fi fill, imgDepth, width, inverse int, offset, scale float64, ang string) {
	//f := fill[fillPos]
	f := fi
	// Calculate the gradient and then fill the depth
	shift := int(math.Pow(2, float64(12-f.depth)))

	scaledShift := 1 / scale
	// Tune the scale to be valid with 4 bit colours if in the negative direction
	/*if f.direction < 0 {
		//only tune if it is below something

		// fmt.Println(f.start, f.start%256, f)
		if f.start%256 == 0 {
			f.start-- // Take away 1 to make it friendly for 4 bit
		} else if f.start%256 != 255 {
			f.start = int(validStart(float64(f.start), f.direction, 256.0)) - 1
		}
	}*/

	colourPos := float64(f.start)
	// Make sure the ramp is a suitable start point for the bit depth
	colourPos = validStart(colourPos, f.direction, float64(shift))

	for i := 0; i < width; i++ {
		// Assignrgb values as a colour
		colourRGB, _ := assignRGBValues(f.colour[0], colourPos, uint16(r.Minimum), uint16(r.Maximum))

		// Assign that colour for the depth of the stripe
		for j := float64(offset); j < float64(offset+f.height); j++ {

			set(ang, canvas, colourRGB, float64(i), j)
		}

		if s := math.Mod(float64((i + 1)), float64(shift)*scale); int(s) == 0 {
			if float64(shift)*scale < 1 {
				colourPos = colourPosShift(colourPos, scaledShift, f.direction, r.Minimum, r.Maximum)
			} else {
				colourPos = colourPosShift(colourPos, float64(shift), f.direction, r.Minimum, r.Maximum)
			}
		}
	}
	// Draw the label over the stripe if labels have been declared
	if r.Text != nil && f.label != "" {
		r.Text.labels(canvas, f.label, ang, inverse, width, offset, f.height)
	}
}

// fill creates the body of fill struct based off of the group header.
// This is designed to have the values updated for each separate group
func groupFill(body stripeHeadersJSON) []fill {
	var fills []fill

	// Make the group
	// Then alternate between
	if body.Header != nil {
		var f fill
		f.colour = body.Header.Colour
		f.gradient = "header"
		f.height = body.Header.Height
		f.depth = 12
		fills = append(fills, f)
	}
	if body.Stripes != nil {
		// assign all the values
		for i, d := range body.Stripes.Bitdepth {
			var fg fill
			fg.height = body.Stripes.Height
			fg.depth = d

			fg.gradient = body.Stripes.Fill
			// assign a a label (if possible)
			if i < len(body.Stripes.Labels) {
				fg.label = body.Stripes.Labels[i]
			}

			fills = append(fills, fg)
			// Make interstripe
			if i != len(body.Stripes.Bitdepth)-1 && body.InterStripe != nil { // Don't add on an interstripe divider if it's at the end
				var fInt fill
				fInt.height = body.InterStripe.Height
				fInt.colour = body.InterStripe.Colour
				fInt.gradient = "header"
				fInt.depth = 12
				fills = append(fills, fInt)
			}
		}
	}

	return fills
}

// fillUpdate updates all the stripe values for each group type
func fillUpdate(fills []fill, s stripesJSON, depth int) []fill {
	shift := 1
	if shift != 0 {
		shift = int(math.Pow(2, float64(12-depth)))
	}
	for i, f := range fills {
		if f.gradient != "header" {
			f.colour = []string{s.Colour}
			f.start = s.RampStart * shift
			f.direction = s.Direction
			fills[i] = f
		}
	}

	return fills
}

// generate this struct then populate the stripes
type fill struct {
	gradient  string
	colour    []string
	label     string
	height    float64
	start     int
	depth     int
	direction int
}

// ratioToheight returns the height of thr groups, the stripe and the interstripe section
func ratioToHeight(heights []fill, canvasDepth, gMult int) ([]fill, error) {
	total := 0.0
	for _, f := range heights {
		total += f.height
	}
	scale := float64(canvasDepth) / (total * float64(gMult))
	if scale < 0.9999999 { // Accounting for float errors
		return nil, fmt.Errorf("0124 The total depth of the ramps %v, is larger than the depth of the canvas %v stripeDepth", total*float64(gMult), canvasDepth)
	}

	for i, f := range heights {
		heights[i].height = f.height * scale
	}

	return heights, nil
}

func assignRGBValues(colour string, rgb float64, maxBlack, maxWhite uint16) (color.NRGBA64, error) {
	switch strings.ToLower(colour) {
	case "grey", "gray": // "black", "white",
		return color.NRGBA64{uint16(rgb) << 4, uint16(rgb) << 4, uint16(rgb) << 4, uint16(0xffff)}, nil
	case "black":
		return color.NRGBA64{maxBlack << 4, maxBlack << 4, maxBlack << 4, uint16(0xffff)}, nil
	case "white":
		return color.NRGBA64{maxWhite << 4, maxWhite << 4, maxWhite << 4, uint16(0xffff)}, nil
	case "red":
		return color.NRGBA64{uint16(rgb) << 4, 0, 0, uint16(0xffff)}, nil
	case "green":
		return color.NRGBA64{0, uint16(rgb) << 4, 0, uint16(0xffff)}, nil
	case "blue":
		return color.NRGBA64{0, 0, uint16(rgb) << 4, uint16(0xffff)}, nil
	default:
		return color.NRGBA64{0, 0, 0, 0}, fmt.Errorf("%s Non specific colour called, rgb values set at 0", colour) // Unused error
	}
}

// valid start checks the rgb value is a multiple of the shift value and that the resulting values
// will be multiples of the shift value, e.g 4bit will progress 0,256,512
func validStart(rgb float64, brightToDark int, shift float64) float64 {

	pass := false
	// If the colour isn't a constant
	if brightToDark != 0 {
		// While the colour isn't a multiple of the shift it isn't representing that bit value
		for !pass {
			// Increment in the direction of the ramp until it does
			if s := math.Mod(float64((rgb)), shift); int(s) != 0 {
				rgb += 1 * float64(brightToDark)
			} else {
				pass = true
			}
		}
	}

	return rgb
}

// use for alternating between colours in a loop across the canvas
func colourSwap(colour []string, c string) string {
	// Guard statement for arrays of length one
	if len(colour) == 1 {
		return colour[0]
	}
	// Loop through the colours to find the current one and then move onto
	// The next value in the array or start from the beginning
	for i, v := range colour {
		if c == v && i != (len(colour)-1) {
			return colour[i+1]
		} else if c == v && i == (len(colour)-1) {
			return colour[0]
		}
	}
	// In theory the code should never get here but a return statement is required
	return "black"
}

// ColourPosShift moves the colour along one colour value
// if the colour overruns the specified range then it is reset back to the opposite colour
// E.g. if it over runs black then it is reset to a white value
func colourPosShift(c, scaleShift float64, brightToDark, maxBlack, maxWhite int) float64 {
	// Shift the colour
	c += (float64(brightToDark) * scaleShift)

	// Check if it's within the colour range
	if !(maxBlack < int(c) && int(c) < maxWhite) {
		// Ramp is the range of colours available plus 1 for the ramp
		rampRange := float64(maxWhite - maxBlack + 1)
		// Reset back to the other end giving a clean start
		if int(c) > maxWhite {
			c -= rampRange
		} else if int(c) < maxBlack {
			c += rampRange
		}
	}

	// Check it is a valid multiple of the bit depth
	return validStart(c, brightToDark, scaleShift)
}

const (
	rotate180  = "rotate180"
	rotate90   = "rotate90 "
	rotate270  = "rotate270"
	noRotation = "xy"
)

// func set sets the canvas values based on the roatation without running a transformation
func set(ang string, canvas draw.Image, colourRGB color.NRGBA64, i, j float64) {

	b := canvas.Bounds().Max
	switch ang {
	case noRotation:
		canvas.Set(int(i), int(j), colourRGB)
	case rotate180:
		canvas.Set(b.X-(int(i)+1), b.Y-(int(j)+1), colourRGB)
	case rotate270:
		canvas.Set(b.X-int(j), b.Y-(int(i)+1), colourRGB)
	default:
		canvas.Set(int(j), int(i), colourRGB)
	}
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
