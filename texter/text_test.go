package texter

import (
	"context"
	"image"
	"image/png"
	"os"
	"testing"
)

func TestBadStrings(t *testing.T) {

	mockContext := context.Background()

	base := image.NewNRGBA64(image.Rect(0, 0, 1000, 1000))
	TextboxJSON{Font: "title", Back: "#000000", Textc: "#ffffff"}.DrawString(base, &mockContext, "A long winding sentence")

	f, _ := os.Create("testdata/A.png")
	png.Encode(f, base)

	TextboxJSON{Font: "title", Back: "#000000", Textc: "#ffffff", FillType: FillTypeFull}.DrawString(base, &mockContext, "A")

	fill, _ := os.Create("testdata/AFull.png")
	png.Encode(fill, base)

	TextboxJSON{Font: "title", Back: "#000000", Textc: "#ffffff"}.DrawStrings(base, &mockContext, []string{"The quick",
		"brown", "dog", "jumped over the lazy fox"})

	flines, _ := os.Create("testdata/lines.png")
	png.Encode(flines, base)

	TextboxJSON{Font: "pixel", Back: "#000000", Textc: "#ffffff",
		FillType: FillTypeFull, XAlignment: AlignmentRight, YAlignment: AlignmentBottom,
	}.DrawStrings(base, &mockContext, []string{"The quick",
		"brown", "dog", "jumped over the lazy fox"})

	flinesf, _ := os.Create("testdata/linesFull.png")
	png.Encode(flinesf, base)

}
