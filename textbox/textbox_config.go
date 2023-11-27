package textbox

import (
	"github.com/mmTristan/opentsg-core/colour"
	"github.com/mmTristan/opentsg-core/config"
)

type TextboxJSON struct {
	// Type       string       `json:"type" yaml:"type"`
	Text []string `json:"text" yaml:"text"`

	GridLoc     *config.Grid      `json:"grid" yaml:"grid"`
	ColourSpace colour.ColorSpace `json:"ColorSpace,omitempty" yaml:"ColorSpace,omitempty"`
	Border      string            `json:"bordercolor" yaml:"bordercolor"`
	BorderSize  float64           `json:"bordersize" yaml:"bordersize"`
	Font        string            `json:"font" yaml:"font"`

	Back       string `json:"backgroundcolor" yaml:"backgroundcolor"`
	Textc      string `json:"textcolor" yaml:"textcolor"`
	FillType   string `json:"fillType" yaml:"fillType"`
	XAlignment string `json:"xAlignment" yaml:"xAlignment"`
	YAlignment string `json:"yAlignment" yaml:"yAlignment"`
}

var textBoxSchema = []byte(`{
	"$schema": "https://json-schema.org/draft/2020-12/schema",
	"$id": "https://example.com/product.schema.json",
	"title": "Allow anything through for tests",
	"description": "An empty schema to allow custom structs to run through",
	"type": "object"
	}`)

func (tb TextboxJSON) Alias() string {
	return tb.GridLoc.Alias
}

func (tb TextboxJSON) Location() string {
	return tb.GridLoc.Location
}
