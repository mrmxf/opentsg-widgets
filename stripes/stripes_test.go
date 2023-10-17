package stripes

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"

	examplejson "github.com/mmTristan/opentsg-widgets/exampleJson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTemp(t *testing.T) {
	input := `
{
	"type": "builtin.ramps",
	"rampAngle": 0,
	"minimum": 0,
	"maximum": 1023,
	"depth": 10,
	"fillType": "fill",
	"stripes": {
	  "ramps": {
		"bitDepth": [
		  8,4
		],
		"fill": "gradient",
		"height": 8,
		"rampGroups": {
		  "gray1": {
			"color": "gray",
			"rampstart": 1023,
			"direction": -1
		  },
		  "gray2": {
			"color": "gray",
			"rampstart": 0,
			"direction": 1
		  },
		  "gray3": {
			"color": "gray",
			"rampstart": 302,
			"direction": 1
		  }
		}
	  }
	},
	"grid": {
	  "location": "c3:n3"
	}
  }`

	var mockJson rampJSON
	json.Unmarshal([]byte(input), &mockJson)

	testImg := image.NewNRGBA64(image.Rect(0, 0, 1024, 500))
	mockJson.Generate(testImg)

	f, _ := os.Create("example.png")
	png.Encode(f, testImg)
}

func TestNoTruncate(t *testing.T) {
	// This test suite still has issues with the draw function in golang in go 1.18
	b, _ := os.ReadFile("./testdata/angletest.json")
	var mock rampJSON
	_ = json.Unmarshal(b, &mock)
	explanation := []string{"noTruncation"}
	testF := []string{"./testdata/noTruncate.png"}

	for i, exp := range explanation {
		mock.FillType = "fill"
		mock.Angle = ""
		myImage := image.NewNRGBA64(image.Rectangle{image.Point{0, 0}, image.Point{2048, 2000}})
		examplejson.SaveExampleJson(mock, widgetType, exp)
		// Generate the ramp image
		genErr := mock.Generate(myImage)
		// Open the image to compare to
		file, _ := os.Open(testF[i])
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
		// f, _ := os.Create("./testdata/g.png")
		// png.Encode(f, myImage)

		Convey("Checking the ramps can use the \"fill\" method", t, func() {
			Convey("Generating a ramp with a 12 bit range of 4096 on an image iwth a width of 2048", func() {
				Convey("No error is returned and the ramp still runs from 0 to 4096", func() {
					So(genErr, ShouldBeNil)
					So(htest.Sum(nil), ShouldResemble, hnormal.Sum(nil))
				})
			})
		})
	}
}

func TestRightAngles(t *testing.T) {
	// This test suite still has issues with the draw function in golang in go 1.18
	b, _ := os.ReadFile("./testdata/angletest.json")
	var mock rampJSON
	_ = json.Unmarshal(b, &mock)
	explanation := []string{"flat", "90degrees", "180degrees", "270degrees"}
	angles := []string{"", "π*1/2", "π*1", "π*3/2"}
	testF := []string{"./testdata/test.png", "./testdata/test90.png", "./testdata/test180.png", "./testdata/test270.png"}

	for i, angle := range angles {
		mock.Angle = angle
		myImage := image.NewNRGBA64(image.Rectangle{image.Point{0, 0}, image.Point{4096, 2000}})
		examplejson.SaveExampleJson(mock, widgetType, explanation[i])
		// Generate the ramp image
		genErr := mock.Generate(myImage)
		// Open the image to compare to
		file, _ := os.Open(testF[i])
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
		// f, _ := os.Create(testF[i] + fmt.Sprintf("%v.png", i))
		// png.Encode(f, myImage)

		Convey("Checking the ramps are generated at 90 degree angles", t, func() {
			Convey(fmt.Sprintf("Comparing the generated ramp to %v with an angle of %v", testF[i], angle), func() {
				Convey("No error is returned and the file matches", func() {
					So(genErr, ShouldBeNil)
					So(htest.Sum(nil), ShouldResemble, hnormal.Sum(nil))
				})
			})
		})
	}
}

func TestAngles(t *testing.T) {
	// This test suite still has issues with the draw function in golang
	b, _ := os.ReadFile("./testdata/angletest.json")
	var mock rampJSON
	_ = json.Unmarshal(b, &mock) // "π*1/20"
	angles := []string{"π*1/20", "π*5/12", "π*9/10", "π*31/20"}
	explanation := []string{"9degrees", "75degrees", "162degrees", "279degrees"}
	testF := []string{"./testdata/angLinear.png", "./testdata/ang90.png", "./testdata/ang180.png", "./testdata/ang270.png"}

	for i, angle := range angles {
		mock.Angle = angle
		myImage := image.NewNRGBA64(image.Rectangle{image.Point{0, 0}, image.Point{4096, 2000}})
		examplejson.SaveExampleJson(mock, widgetType, explanation[i])
		// Generate the ramp image
		genErr := mock.Generate(myImage)
		// Open the image to compare to
		file, _ := os.Open(testF[i])
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
		Convey("Checking the ramps are generated at angles other than 90 degrees", t, func() {
			Convey(fmt.Sprintf("Comparing the generated ramp to %v with an angle of %v", testF[i], angle), func() {
				Convey("No error is returned and the file matches", func() {
					So(genErr, ShouldBeNil)
					So(htest.Sum(nil), ShouldResemble, hnormal.Sum(nil))
				})
			})
		})
	}
}

func TestInitErr(t *testing.T) {
	var rampMock rampJSON

	rampMock.Depth = 8
	rampMock.Minimum = 100

	wvTests := []int{280, 16}
	exepecErr := []string{"0121 The Max white value 280 is higher than the maximum value of 256 for a bit depth of 8",
		"0122 The black value 100 is higher than the maximum white value of 16"}

	for i := range wvTests {
		rampMock.Maximum = wvTests[i]
		initErr := rampMock.constantInit()

		// Save the file
		Convey("Checking the init errors are returned", t, func() {
			Convey(fmt.Sprintf("Using a max white colour %v ", wvTests[i]), func() {
				Convey(fmt.Sprintf("An error of %v", exepecErr[i]), func() {
					So(initErr.Error(), ShouldEqual, exepecErr[i])
				})
			})
		})
	}
}

/*
tests to add
text positions with different angles?
3 labels 4 stripes
*/
