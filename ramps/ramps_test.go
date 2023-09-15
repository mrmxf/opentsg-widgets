package ramps

import (
	"image"
	"image/png"
	"os"
	"testing"
	// examplejson "github.com/mrmxf/opentsg-widgets/exampleJson"
	// . "github.com/smartystreets/goconvey/convey"
)

func TestTemp(t *testing.T) {
	mock := ramp{stripes: []string{"red", "blue"}, l: layout{header: 5, interStripe: 3, ramp: []int{4, 3, 5}}}
	tester := image.NewNRGBA64(image.Rect(0, 0, 1000, 1000))
	firstrun(tester, mock)

	f, _ := os.Create("tester.png")
	png.Encode(f, tester)
}
