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

	"github.com/mmTristan/opentsg-core/colour"
	examplejson "github.com/mmTristan/opentsg-widgets/exampleJson"
	"github.com/mmTristan/opentsg-widgets/text"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLines(t *testing.T) {
	mockContext := context.Background()

	mockTB := TextboxJSON{
		Textc: "#C2A649", Border: "#f0f0f0", Back: "#ffffff", BorderSize: 20,
		YAlignment: text.AlignmentMiddle}
	stringsToCheck := [][]string{{"sample text"}, {"sample", "text"}}
	original := []string{"./testdata/singleline.png", "./testdata/multiline.png"}
	explanation := []string{"singleline", "multiline"}

	for i, str := range stringsToCheck {

		myImage := colour.NewNRGBA64(colour.ColorSpace{}, image.Rectangle{image.Point{0, 0}, image.Point{1024, 240}})
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
		htest.Write(myImage.Pix())

		//f, _ := os.Create("./testdata/" + fmt.Sprintf("%v", i) + ".png")
		//colour.PngEncode(f, myImage)
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
func TestFontImport(t *testing.T) {

	//	mockContext := context.Background()

	base := image.NewNRGBA64(image.Rect(0, 0, 1000, 1000))
	//	text := texter.TextboxJSON{Textc: "#260498", Back: "#980609"}
	TextboxJSON{Border: "#800080", BorderSize: 5, Textc: "#260498", Back: "#980609", Text: []string{"The quick",
		"brown dog jumped", "over the lazy gray fox"}, Font: `https://get.fontspace.co/webfont/lgwK0/M2ZmY2VhZDMxMTNhNGE1Yzk2Y2JhZTEwNzgwOTNkN2YudHRm/halloween-clipart.ttf`}.Generate(base)

	f, _ := os.Create("testdata/A.png")
	png.Encode(f, base)

}
