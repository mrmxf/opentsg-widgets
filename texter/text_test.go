package texter

import (
	"context"
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/mmTristan/opentsg-core/colour"
)

func TestBadStrings(t *testing.T) {

	mockContext := context.Background()

	base := image.NewNRGBA64(image.Rect(0, 0, 1000, 1000))
	TextboxJSON{font: "title", back: &colour.CNRGBA64{A: 0xffff}, textc: &colour.CNRGBA64{R: 0xffff, A: 0xffff}}.DrawString(base, &mockContext, "A long winding sentence")

	f, _ := os.Create("testdata/A.png")
	png.Encode(f, base)

	TextboxJSON{font: "title", back: &colour.CNRGBA64{A: 0xffff}, textc: &colour.CNRGBA64{R: 0xffff, A: 0xffff}, fillType: FillTypeFull}.DrawString(base, &mockContext, "A")

	fill, _ := os.Create("testdata/AFull.png")
	png.Encode(fill, base)

	TextboxJSON{font: "title", back: &colour.CNRGBA64{A: 0xffff}, textc: &colour.CNRGBA64{R: 0xffff, A: 0xffff}}.DrawStrings(base, &mockContext, []string{"The quick",
		"brown", "dog", "jumped over the lazy fox"})

	flines, _ := os.Create("testdata/lines.png")
	png.Encode(flines, base)

	TextboxJSON{font: "pixel", back: &colour.CNRGBA64{A: 0xffff}, textc: &colour.CNRGBA64{R: 0xffff, A: 0xffff},
		fillType: FillTypeFull, xAlignment: AlignmentRight, yAlignment: AlignmentBottom,
	}.DrawStrings(base, &mockContext, []string{"The quick",
		"brown", "dog", "jumped over the lazy fox"})

	flinesf, _ := os.Create("testdata/linesFull.png")
	png.Encode(flinesf, base)

}
