package twosi

import (
	_ "embed"

	"github.com/mmTristan/opentsg-core/colour"
	"github.com/mmTristan/opentsg-core/config"
)

/*
// Ebu3373/twosi definitions
const wName = "twosi"
const wType = "ebu3373/twosi"
const wLibrary = "builtin"
const hooks = ""*/

type twosiJSON struct {
	//	Type    string      `json:"type" yaml:"type"`
	GridLoc     config.Grid       `json:"grid,omitempty" yaml:"grid,omitempty"`
	ColourSpace colour.ColorSpace `json:"ColorSpace,omitempty" yaml:"ColorSpace,omitempty"`
}

//go:embed jsonschema/twoschema.json
var schemaInit []byte

func (t twosiJSON) Alias() string {
	return t.GridLoc.Alias
}

func (t twosiJSON) Location() string {
	return t.GridLoc.Location
}
