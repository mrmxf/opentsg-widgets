package bars

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
	// . "github.com/smartystreets/goconvey/convey"
)

func TestTemp(t *testing.T) {

	fills := [][]bars{{{width: d, color: gray40}, {width: f, color: white75}, {width: c, color: yellow75}, {width: c, color: cyan75}, {width: c, color: green75}, {width: c, color: mag75}, {width: c, color: red75}, {width: f, color: blue75}, {width: d, color: gray40}},
		{{width: d, color: cyan100}, {width: f, color: white100}, {width: 5*c + f, color: white75}, {width: d, color: blue100}},
		{{width: d, color: yellow100}, {width: f, color: black0}, {width: 5 * c, fill: yRamp}, {width: f, color: white100}, {width: d, color: red100}},
		{{width: d, color: gray15}, {width: k, fill: superBlack}, {width: g, fill: superWhite}, {width: h, color: black0}, {width: (69 / 1920.0), color: black2Neg}, {width: (68 / 1920.0), color: black0}, {width: (69 / 1920.0), color: black2Pos}, {width: (68 / 1920.0), color: black0}, {width: (69 / 1920.0), color: black4Pos}, {width: c, color: black0}, {width: d, color: gray15}}}

	HD := image.NewNRGBA64(image.Rect(0, 0, 1920, 1080))
	bar(HD, fills)

	fe, _ := os.Create("test.png")
	png.Encode(fe, HD)

	fmt.Println(color.YCbCrToRGB(219, 16, 138))

	fillsSD := [][]bars{{{width: d, color: gray40SD}, {width: f, color: white75SD}, {width: c, color: yellow75SD}, {width: c, color: cyan75SD}, {width: c, color: green75SD}, {width: c, color: mag75SD}, {width: c, color: red75SD}, {width: f, color: blue75SD}, {width: d, color: gray40SD}},
		{{width: d, color: cyan100SD}, {width: f, color: white100SD}, {width: 5*c + f, color: white75SD}, {width: d, color: blue100SD}},
		{{width: d, color: yellow100SD}, {width: f, color: black0SD}, {width: 5 * c, fill: yRampSD}, {width: f, color: white100SD}, {width: d, color: red100SD}},
		{{width: d, color: gray15SD}, {width: k, fill: superBlackSD}, {width: g, fill: superWhiteSD}, {width: h, color: black0SD}, {width: (69 / 1920.0), color: black2NegSD}, {width: (68 / 1920.0), color: black0SD}, {width: (69 / 1920.0), color: black2PosSD}, {width: (68 / 1920.0), color: black0SD}, {width: (69 / 1920.0), color: black4PosSD}, {width: c, color: black0SD}, {width: d, color: gray15SD}}}

	SD := image.NewNRGBA(image.Rect(0, 0, 720, 480))
	bar(SD, fillsSD)

	fmt.Println(SD.At(234, 300))

	fsd, _ := os.Create("testSD.png")
	png.Encode(fsd, SD)

}
