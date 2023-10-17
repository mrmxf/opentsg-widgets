package fourcolour

import (
	"github.com/mmTristan/opentsg-core/config"
)

type fourJSON struct {
	Colourpallette []string     `json:"colors" yaml:"colors"`
	GridLoc        *config.Grid `json:"grid,omitempty" yaml:"grid,omitempty"`
}

var schemaInit = []byte(`{
	"$schema": "https://json-schema.org/draft/2020-12/schema",
	"$id": "https://example.com/product.schema.json",
	"title": "Allow anything through for tests",
	"description": "An empty schema to allow custom structs to run through",
	"type": "object"
	}`)

func (f fourJSON) Alias() string {
	return f.GridLoc.Alias
}

func (f fourJSON) Location() string {
	return f.GridLoc.Location
}
