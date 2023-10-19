package texter

import (
	"github.com/mmTristan/opentsg-core/colour"
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
