package textbox

import (
	_ "embed"

	"github.com/mrmxf/opentsg-cote/config"
)

/*
// textbox definitions
const wName = "textbox"
const wType = "textbox"
const wLibrary = "builtin"
const hooks = ""*/

type TextboxJSON struct {
	Type       string       `json:"type" yaml:"type"`
	Text       []string     `json:"text" yaml:"text"`
	Font       string       `json:"font" yaml:"font"`
	GridLoc    *config.Grid `json:"grid" yaml:"grid"`
	Back       string       `json:"backgroundcolor" yaml:"backgroundcolor"`
	Border     string       `json:"bordercolor" yaml:"bordercolor"`
	Textc      string       `json:"textcolor" yaml:"textcolor"`
	BorderSize float64      `json:"bordersize" yaml:"bordersize"`
}

//go:embed jsonschema/textboxschema.json
var textBoxSchema []byte

func (tb TextboxJSON) Alias() string {
	return tb.GridLoc.Alias
}

func (tb TextboxJSON) Location() string {
	return tb.GridLoc.Location
}
