package textbox2

import (
	"github.com/mmTristan/opentsg-core/colour"
	"github.com/mmTristan/opentsg-widgets/texter"
	"github.com/mrmxf/opentsg-core/config"
)

/*
textbox 2 has the border on the outside
then the textbox in the middle

*/

type TextboxJSON struct {
	// Type       string       `json:"type" yaml:"type"`
	Text           []string `json:"text" yaml:"text"`
	TextProperties texter.TextboxJSON
	GridLoc        *config.Grid      `json:"grid" yaml:"grid"`
	ColourSpace    colour.ColorSpace `json:"ColorSpace,omitempty" yaml:"ColorSpace,omitempty"`
	Border         string            `json:"bordercolor" yaml:"bordercolor"`
	BorderSize     float64           `json:"bordersize" yaml:"bordersize"`
}
