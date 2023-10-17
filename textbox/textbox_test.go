package textbox

import (
	"context"
	"crypto/sha256"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"

	examplejson "github.com/mmTristan/opentsg-widgets/exampleJson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestZoneGenAngle(t *testing.T) {
	mockContext := context.Background()

	var mockTB TextboxJSON

	mockTB.Textc = "#C2A649"
	mockTB.Border = "#C2A649"
	mockTB.Back = "#ffffff"
	mockTB.BorderSize = 0.02666
	stringsToCheck := [][]string{{"sample text"}, {"sample", "text"}}
	original := []string{"./testdata/singleline.png", "./testdata/multiline.png"}
	explanation := []string{"singleline", "multiline"}

	for i, str := range stringsToCheck {

		myImage := image.NewNRGBA64(image.Rectangle{image.Point{0, 0}, image.Point{1024, 240}})
		mockTB.Text = str
		genErr := mockTB.Generate(myImage, &mockContext)
		examplejson.SaveExampleJson(mockTB, widgetType, explanation[i])
		file, _ := os.Open(original[i])
		// Decode to get the colour values
		baseVals, _ := png.Decode(file)

		// Assign the colour to the correct type of image NGRBA64 and replace the colour values
		readImage := image.NewNRGBA64(baseVals.Bounds())
		draw.Draw(readImage, readImage.Bounds(), baseVals, image.Point{0, 0}, draw.Over)

		// Make a hash of the pixels of each image
		hnormal := sha256.New()
		htest := sha256.New()
		hnormal.Write(readImage.Pix)
		htest.Write(myImage.Pix)

		//	f, _ := os.Create("./testdata/" + fmt.Sprintf("%v", i) + ".png")
		// Png.Encode(f, myImage)
		// Save the file
		Convey("Checking that strings are generated", t, func() {
			Convey(fmt.Sprintf("Generating an image with the following strings: %v ", str), func() {
				Convey("No error is returned and the file matches exactly", func() {
					So(genErr, ShouldBeNil)
					So(htest.Sum(nil), ShouldResemble, hnormal.Sum(nil))
				})
			})
		})
	}
}

/*
` {
	"text": ["4k-test00-white_noise"],
	"textfile": "./MavenPro-Bold.ttf",
	"savefile": "./newpub/noize.png",
	"bordercolour": "#C2A649",
	"textcolour": "#C2A649",
	"backgroundcolour": "#ffffff",
	"dimension": {
		"w": 2560,
		"h": 240
	}`*/
