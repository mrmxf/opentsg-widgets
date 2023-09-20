package ramps

import (
	"image"
	"image/png"
	"os"
	"testing"

	examplejson "github.com/mrmxf/opentsg-widgets/exampleJson"
	// . "github.com/smartystreets/goconvey/convey"
)

func TestTemp(t *testing.T) {
	mock := Ramp{Stripes: []RampProperties{{Colour: "red"}, {Colour: "blue", StartPoint: 2500}, {Colour: "red", StartPoint: 3199, Reverse: true}, {Colour: "green", StartPoint: 3967, Reverse: true}},
		StripeGroup: layout{Header: internalHeader{Height: 4, Colour: "white"},
			InterStripe: alternateHeader{Colours: []string{"white", "black"}, Height: 2},
			Ramp:        []Stripe{{Height: 4, BitDepth: 12}, {Height: 3, BitDepth: 10}, {Height: 5, BitDepth: 8}, {Height: 5, BitDepth: 5}}},
		WidgetProperties: control{GlobalBitDepth: 12}}
	tester := image.NewNRGBA64(image.Rect(0, 0, 5000, 1000))
	firstrun(tester, mock)

	examplejson.SaveExampleJson(mock, "builtin.ramps2", "demo")

	f, _ := os.Create("tester.png")
	png.Encode(f, tester)

	rotates := []string{"π*1/2", "π*3/2", "π*2/2"}
	names := []string{"tester90.png", "tester270.png", "tester180.png"}

	for i, ang := range rotates {
		mockAngle := mock
		mockAngle.WidgetProperties.Angle = ang

		testerAng := image.NewNRGBA64(image.Rect(0, 0, 4000, 4000))
		firstrun(testerAng, mockAngle)

		fang, _ := os.Create(names[i])
		png.Encode(fang, testerAng)
	}

	mock.Stripes = []RampProperties{{Colour: "gray", StartPoint: 1023, Reverse: true}}
	mock.WidgetProperties = control{GlobalBitDepth: 10}
	mock.StripeGroup.Ramp = []Stripe{{Height: 20, BitDepth: 8}, {Height: 20, BitDepth: 4}}
	tester2 := image.NewNRGBA64(image.Rect(0, 0, 5000, 1000))
	firstrun(tester2, mock)

	f2, _ := os.Create("tester2.png")
	png.Encode(f2, tester2)

}
