package textbox2

import (
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/mmTristan/opentsg-widgets/texter"
)

func TestBadStrings(t *testing.T) {

	//	mockContext := context.Background()

	base := image.NewNRGBA64(image.Rect(0, 0, 1000, 1000))
	text := texter.TextboxJSON{Textc: "#260498", Back: "#980609"}
	TextboxJSON{Border: "#800080", BorderSize: 5, TextProperties: text, Text: []string{"surpise", "colours"}}.Generate(base)

	f, _ := os.Create("testdata/A.png")
	png.Encode(f, base)

}
