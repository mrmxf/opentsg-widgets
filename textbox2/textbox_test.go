package textbox2

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func TestBadStrings(t *testing.T) {

	//	mockContext := context.Background()

	base := image.NewNRGBA64(image.Rect(0, 0, 1000, 1000))
	//	text := texter.TextboxJSON{Textc: "#260498", Back: "#980609"}
	TextboxJSON{Border: "#800080", BorderSize: 5, Textc: "#260498", Back: "#980609", Text: []string{"surpise", "colours"}, Font: `https://get.fontspace.co/webfont/lgwK0/M2ZmY2VhZDMxMTNhNGE1Yzk2Y2JhZTEwNzgwOTNkN2YudHRm/halloween-clipart.ttf`}.Generate(base)

	f, _ := os.Create("testdata/A.png")
	png.Encode(f, base)

}
