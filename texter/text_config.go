package texter

import (
	"github.com/mmTristan/opentsg-core/colour"
	"github.com/mmTristan/opentsg-core/colourgen"
)

type TextboxJSON struct {
	// Type       string       `json:"type" yaml:"type"`
	Font        string            `json:"font" yaml:"font"`
	ColourSpace colour.ColorSpace `json:"ColorSpace,omitempty" yaml:"ColorSpace,omitempty"`
	Back        string            `json:"backgroundcolor" yaml:"backgroundcolor"`
	Textc       string            `json:"textcolor" yaml:"textcolor"`
	FillType    string
	XAlignment  string
	YAlignment  string
}

type TextboxJSON2 struct {
	// Type       string       `json:"type" yaml:"type"`
	Font        string            `json:"font" yaml:"font"`
	ColourSpace colour.ColorSpace `json:"ColorSpace,omitempty" yaml:"ColorSpace,omitempty"`
	Back        *colour.CNRGBA64  `json:"backgroundcolor" yaml:"backgroundcolor"`
	Textc       *colour.CNRGBA64  `json:"textcolor" yaml:"textcolor"`
	FillType    string
	XAlignment  string
	YAlignment  string
}

func NewTextboxer(ColourSpace colour.ColorSpace, options ...func(*TextboxJSON2)) *TextboxJSON2 {
	txt := &TextboxJSON2{ColourSpace: ColourSpace}
	for _, opt := range options {
		opt(txt)
	}
	return txt
}

func WithFill(fill string) func(t *TextboxJSON2) {

	return func(t *TextboxJSON2) {
		t.FillType = fill
	}
}

func WithFont(font string) func(t *TextboxJSON2) {

	return func(t *TextboxJSON2) {
		t.Font = font
	}
}

func WithTextColourString(colour string) func(t *TextboxJSON2) {

	return func(t *TextboxJSON2) {
		c := colourgen.HexToColour(colour, t.ColourSpace)
		t.Textc = c
	}
}

func WithTextColour(colour *colour.CNRGBA64) func(t *TextboxJSON2) {

	return func(t *TextboxJSON2) {

		t.Textc = colour
	}
}

func WithBackgroundColourString(colour string) func(t *TextboxJSON2) {

	return func(t *TextboxJSON2) {
		c := colourgen.HexToColour(colour, t.ColourSpace)
		t.Back = c
	}
}

func WithBackgroundColour(colour *colour.CNRGBA64) func(t *TextboxJSON2) {

	return func(t *TextboxJSON2) {

		t.Back = colour
	}
}

func WithXAlignment(x string) func(t *TextboxJSON2) {
	return func(t *TextboxJSON2) {
		t.XAlignment = x
	}
}

func WithYAlignment(y string) func(t *TextboxJSON2) {
	return func(t *TextboxJSON2) {
		t.XAlignment = y
	}
}
