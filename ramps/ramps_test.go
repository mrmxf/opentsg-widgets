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
	mock := Ramp{Stripes: []RampProperties{{Colour: "green", StartPoint: 960}, {Colour: "gray", StartPoint: 960}},
		StripeGroup: layout{Header: internalHeader{Height: 0, Colour: "white"},
			InterStripe: alternateHeader{Colours: []string{"white", "black", "red", "blue"}, Height: 1},
			Ramp:        []Stripe{{Height: 5, BitDepth: 4, Label: "4b"}, {Height: 5, BitDepth: 6, Label: "6b"}, {Height: 5, BitDepth: 8, Label: "8b"}, {Height: 5, BitDepth: 10, Label: "10b"}}},
		WidgetProperties: control{GlobalBitDepth: 10, TextProperties: textObjectJSON{TextColour: "#345AB6", TextHeight: 70}}}
	tester := image.NewNRGBA64(image.Rect(0, 0, 1024, 1000)) //960))
	firstrun(tester, mock)

	examplejson.SaveExampleJson(mock, "builtin.ramps2", "demo")

	f, _ := os.Create("tester.png")
	png.Encode(f, tester)

	rotates := []string{"π*1/2", "π*3/2", "π*2/2"}
	names := []string{"tester90.png", "tester270.png", "tester180.png"}

	for i, ang := range rotates {
		mockAngle := mock
		mockAngle.WidgetProperties.Angle = ang
		// mock.WidgetProperties.TextProperties = textObjectJSON{TextColour: "#F32399"}

		testerAng := image.NewNRGBA64(image.Rect(0, 0, 4000, 4000))
		firstrun(testerAng, mockAngle)

		fang, _ := os.Create(names[i])
		png.Encode(fang, testerAng)
	}

	mock.Stripes = []RampProperties{{Colour: "gray", StartPoint: 1023, Reverse: true}}
	mock.WidgetProperties = control{GlobalBitDepth: 10, TextProperties: textObjectJSON{TextColour: "#F32399"}}
	mock.StripeGroup.Ramp = []Stripe{{Height: 20, BitDepth: 8}, {Height: 20, BitDepth: 4}}
	tester2 := image.NewNRGBA64(image.Rect(0, 0, 5000, 1000))
	firstrun(tester2, mock)

	f2, _ := os.Create("tester2.png")
	png.Encode(f2, tester2)

}
