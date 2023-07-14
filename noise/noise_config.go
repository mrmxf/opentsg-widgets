package noise

import (
	_ "embed"

	"github.com/mrmxf/opentsg-core/config"
)

/*
// noise definitions
const wName = "noise"
const wType = "noise"
const wLibrary = "builtin"
const hooks = ""*/

type noiseJSON struct {
	Type      string       `json:"type" yaml:"type"`
	Noisetype string       `json:"noisetype" yaml:"noisetype"`
	Minimum   int          `json:"minimum" yaml:"minimum"`
	Maximum   int          `json:"maximum" yaml:"maximum"`
	GridLoc   *config.Grid `json:"grid,omitempty" yaml:"grid,omitempty"`
}

//go:embed jsonschema/noiseschema.json
var schemaInit []byte

func (n noiseJSON) Alias() string {
	return n.GridLoc.Alias
}

func (n noiseJSON) Location() string {
	return n.GridLoc.Location
}
