package ramps

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strings"
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

func firstrun(target draw.Image, input ramp) {
	// calculate the whole height of each one

	totalHeight := input.l.header + ((len(input.l.ramp) - 1) * input.l.interStripe)
	for _, r := range input.l.ramp {
		totalHeight += r
	}

	fmt.Println(totalHeight)

	totalHeight *= len(input.stripes)

	fmt.Println(totalHeight, target.Bounds().Dy())
	groupStep := float64(target.Bounds().Dy()) / float64(totalHeight)
	fmt.Println(groupStep)

	position := 0.0
	for _, str := range input.stripes {
		fmt.Println(position, "P")
		if input.l.header != 0 {
			// draw the header
			end := int(position + groupStep*float64(input.l.header))

			fmt.Println(int(groupStep*float64(input.l.header)), end-int(position))
			row := image.NewNRGBA64(image.Rect(0, 0, target.Bounds().Dx(), end-int(position)))
			hidden(target, row, internalHeader{"black"}, int(position))

			position += groupStep * float64(input.l.header)
		}

		for i, ramp := range input.l.ramp {

			end := int(position + groupStep*float64(ramp))

			fmt.Println(int(groupStep*float64(ramp)), end-int(position))
			rrow := image.NewNRGBA64(image.Rect(0, 0, target.Bounds().Dx(), end-int(position)))
			hidden(target, rrow, internalHeader{str}, int(position))

			position += groupStep * float64(ramp)

			if i+1 < len(input.l.ramp) {
				// accounts for jumps in floats and ints
				end := int(position + groupStep*float64(input.l.interStripe))

				fmt.Println(int(groupStep*float64(input.l.interStripe)), end-int(position))

				irow := image.NewNRGBA64(image.Rect(0, 0, target.Bounds().Dx(), end-int(position)))
				hidden(target, irow, internalHeader{"white"}, int(position))

				position += groupStep * float64(input.l.interStripe)

				//calculate segments here
			}
		}

	}
}

type internalHeader struct {
	col string
}

func (h internalHeader) Generate(img draw.Image) {
	c, _ := assignRGBValues(h.col, 4095, 0, 4095)
	draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.Point{}, draw.Over)
}

type idea struct {
	parent any
	header any // but actually a sub module that can be called
}

type ramp struct {
	l       layout
	stripes []string
}

type layout struct {
	header, interStripe int
	ramp                []int // just do the heights frst
}

type make interface {
	Generate(img draw.Image)
}

// Defaults give the optional extras?
func hidden(base, img draw.Image, G make, start int) {

	/*
		hidden needs to be something that can be generic and useful

	*/
	G.Generate(img) //add optional parameterss?

	draw.Draw(base, img.Bounds().Add(image.Point{Y: start}), img, image.Point{}, draw.Over)
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
