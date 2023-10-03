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
	mock := Ramp{Groups: []RampProperties{{Colour: "green", InitialPixelValue: 960}, {Colour: "gray", InitialPixelValue: 960}},
		Gradients: groupContents{GroupSeparator: groupSeparator{Height: 0, Colour: "white"},
			GradientSeparator: gradientSeparator{Colours: []string{"white", "black", "red", "blue"}, Height: 1},
			Gradients:         []Gradient{{Height: 5, BitDepth: 4, Label: "ブルースのテスト列"}, {Height: 5, BitDepth: 6, Label: "ब्रूस परीक्षण पंक्ति"}, {Height: 5, BitDepth: 8, Label: "8b"}, {Height: 5, BitDepth: 10, Label: "10b"}}},
		WidgetProperties: control{MaxBitDepth: 10, TextProperties: textObjectJSON{TextColour: "#345AB6", TextHeight: 70}}}
	tester := image.NewNRGBA64(image.Rect(0, 0, 1024, 1000)) //960))
	firstrun(tester, mock)

	examplejson.SaveExampleJson(mock, "builtin.ramps2", "demo")

	f, _ := os.Create("tester.png")
	png.Encode(f, tester)

	rotates := []string{"π*1/2", "π*3/2", "π*2/2", "π*5/8"}
	names := []string{"tester90.png", "tester270.png", "tester180.png", "testerwonk.png"}

	for i, ang := range rotates {
		mockAngle := mock
		mockAngle.WidgetProperties.CwRotation = ang
		// mock.WidgetProperties.TextProperties = textObjectJSON{TextColour: "#F32399"}

		testerAng := image.NewNRGBA64(image.Rect(0, 0, 4000, 4000))
		firstrun(testerAng, mockAngle)

		fang, _ := os.Create(names[i])
		png.Encode(fang, testerAng)
	}

	mock.Groups = []RampProperties{{Colour: "gray", InitialPixelValue: 1023, Reverse: true}}
	mock.WidgetProperties = control{MaxBitDepth: 10, TextProperties: textObjectJSON{TextColour: "#F32399"}}
	mock.Gradients.Gradients = []Gradient{{Height: 20, BitDepth: 8}, {Height: 20, BitDepth: 4}}
	tester2 := image.NewNRGBA64(image.Rect(0, 0, 5000, 1000))
	firstrun(tester2, mock)

	f2, _ := os.Create("tester2.png")
	png.Encode(f2, tester2)

	mock = Ramp{Groups: []RampProperties{{Colour: "green", InitialPixelValue: 960}, {Colour: "gray", InitialPixelValue: 960}},
		Gradients: groupContents{GroupSeparator: groupSeparator{Height: 0, Colour: "white"},
			GradientSeparator: gradientSeparator{Colours: []string{"white", "black", "red", "blue"}, Height: 1},
			Gradients:         []Gradient{{Height: 5, BitDepth: 4, Label: "4b"}, {Height: 5, BitDepth: 6, Label: "6b"}, {Height: 5, BitDepth: 8, Label: "8b"}, {Height: 5, BitDepth: 10, Label: "10b"}}},
		WidgetProperties: control{MaxBitDepth: 10, TextProperties: textObjectJSON{TextColour: "#345AB6", TextHeight: 70}}}
	Squeezer := image.NewNRGBA64(image.Rect(0, 0, 5000, 1000)) //960))
	mock.WidgetProperties.ObjectFitFill = true
	firstrun(Squeezer, mock)
	fsqueeze, _ := os.Create("testerSqu.png")
	png.Encode(fsqueeze, Squeezer)

	twoer := image.NewNRGBA64(image.Rect(0, 0, 5000, 1000)) //960))
	mock.WidgetProperties.ObjectFitFill = false
	mock.WidgetProperties.PixelValueRepeat = 2
	firstrun(twoer, mock)
	fstwo, _ := os.Create("testerTwo.png")
	png.Encode(fstwo, twoer)

}
