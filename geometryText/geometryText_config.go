package geometrytext

import (
	"github.com/mrmxf/opentsg-core/colour"
	"github.com/mrmxf/opentsg-core/config"
)

type geomTextJSON struct {
	TextColour  string            `json:"textColor" yaml:"textColor"`
	GridLoc     *config.Grid      `json:"grid,omitempty" yaml:"grid,omitempty"`
	ColourSpace colour.ColorSpace `json:"colorSpace,omitempty" yaml:"colorSpace,omitempty"`
}

var schemaInit = []byte(`{
	"$schema": "https://json-schema.org/draft/2020-12/schema",
	"$id": "https://example.com/product.schema.json",
	"title": "Allow anything through for tests",
	"description": "An empty schema to allow custom structs to run through",
	"type": "object"
	}`)

func (f geomTextJSON) Alias() string {
	return f.GridLoc.Alias
}

func (f geomTextJSON) Location() string {
	return f.GridLoc.Location
}
