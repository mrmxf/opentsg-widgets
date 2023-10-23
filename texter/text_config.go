package texter

import (
	"github.com/mmTristan/opentsg-core/colour"
	"github.com/mmTristan/opentsg-core/colourgen"
)

/*
type TextboxJSON struct {
	// Type       string       `json:"type" yaml:"type"`
	Font        string            `json:"font" yaml:"font"`
	ColourSpace colour.ColorSpace `json:"ColorSpace,omitempty" yaml:"ColorSpace,omitempty"`
	Back        string            `json:"backgroundcolor" yaml:"backgroundcolor"`
	Textc       string            `json:"textcolor" yaml:"textcolor"`
	FillType    string
	XAlignment  string
	YAlignment  string
}*/

type TextboxJSON struct {
	// Type       string       `json:"type" yaml:"type"`
	font        string            // `json:"font" yaml:"font"`
	colourSpace colour.ColorSpace //`json:"ColorSpace,omitempty" yaml:"ColorSpace,omitempty"`
	back        *colour.CNRGBA64  //`json:"backgroundcolor" yaml:"backgroundcolor"`
	textc       *colour.CNRGBA64  //`json:"textcolor" yaml:"textcolor"`
	fillType    string
	xAlignment  string
	yAlignment  string
}

// NewText
// the order in which the withs are specifired are the orderin which the executed
func NewTextboxer(ColourSpace colour.ColorSpace, options ...func(*TextboxJSON)) *TextboxJSON {
	txt := &TextboxJSON{colourSpace: ColourSpace}
	for _, opt := range options {
		opt(txt)
	}
	return txt
}

// WithFill sets the fill type
func WithFill(fill string) func(t *TextboxJSON) {

	return func(t *TextboxJSON) {
		t.fillType = fill
	}
}

func WithFont(font string) func(t *TextboxJSON) {

	return func(t *TextboxJSON) {
		t.font = font
	}
}

func WithTextColourString(colour string) func(t *TextboxJSON) {

	return func(t *TextboxJSON) {
		c := colourgen.HexToColour(colour, t.colourSpace)
		t.textc = c
	}
}

func WithTextColour(colour *colour.CNRGBA64) func(t *TextboxJSON) {

	return func(t *TextboxJSON) {

		t.textc = colour
	}
}

func WithBackgroundColourString(colour string) func(t *TextboxJSON) {

	return func(t *TextboxJSON) {
		c := colourgen.HexToColour(colour, t.colourSpace)
		t.back = c
	}
}

func WithBackgroundColour(colour *colour.CNRGBA64) func(t *TextboxJSON) {

	return func(t *TextboxJSON) {

		t.back = colour
	}
}

func WithXAlignment(x string) func(t *TextboxJSON) {
	return func(t *TextboxJSON) {
		t.xAlignment = x
	}
}

func WithYAlignment(y string) func(t *TextboxJSON) {
	return func(t *TextboxJSON) {
		t.yAlignment = y
	}
}
