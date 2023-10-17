package stripes

import (
	"fmt"
	"math"

	_ "embed"

	"github.com/mmTristan/opentsg-core/config"
)

/*
// zoneplate definitions
const wName = "stripes"
const wType = "ramps"
const wLibrary = "builtin"
const hooks = ""*/

type rampJSON struct {
	//	Type     string            `json:"type" yaml:"type"`
	Angle    string             `json:"rampAngle" yaml:"rampAngle"`
	Minimum  int                `json:"minimum" yaml:"minimum"`
	Maximum  int                `json:"maximum" yaml:"maximum"`
	Depth    int                `json:"depth" yaml:"depth"`
	FillType string             `json:"fillType" yaml:"fillType"`
	Stripes  *stripeHeadersJSON `json:"stripes,omitempty" yaml:"stripes,omitempty"`
	GridLoc  *config.Grid       `json:"grid,omitempty" yaml:"grid,omitempty"`
	Text     *textObjectJSON    `json:"text,omitempty" yaml:"text,omitempty"`
}

type textObjectJSON struct {
	TextYPosition string  `json:"textyPosition" yaml:"textyPosition"`
	TextXPosition string  `json:"textxPosition" yaml:"textxPosition"`
	TextHeight    float64 `json:"textHeight" yaml:"textHeight"`
	TextColour    string  `json:"textColor" yaml:"textColor"`
}

type stripeHeadersJSON struct {
	Header      *dividerJSON    `json:"groupHeader,omitempty" yaml:"groupHeader,omitempty"`
	Stripes     *stripesObjJSON `json:"ramps,omitempty" yaml:"ramps,omitempty"`
	InterStripe *dividerJSON    `json:"interStripes,omitempty" yaml:"interStripes,omitempty"`
}

type dividerJSON struct {
	Colour []string `json:"color,omitempty" yaml:"color,omitempty"`
	Height float64  `json:"height,omitempty" yaml:"height,omitempty"`
}

type stripesObjJSON struct {
	Bitdepth []int    `json:"bitDepth,omitempty" yaml:"bitDepth,omitempty"`
	Height   float64  `json:"height,omitempty" yaml:"height,omitempty"`
	Labels   []string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Fill     string   `json:"fill,omitempty" yaml:"fill,omitempty"`
	// GroupTypes *[]stripesToDraw         `json:"rampGroups2,omitempty"`
	GroupTypes map[string]stripesJSON `json:"rampGroups,omitempty" yaml:"rampGroups,omitempty"`
}

type stripesJSON struct {
	Colour    string `json:"color,omitempty" yaml:"color,omitempty"`
	RampStart int    `json:"rampstart,omitempty" yaml:"rampstart,omitempty"`
	Direction int    `json:"direction,omitempty" yaml:"direction,omitempty"`
}

//go:embed jsonschema/stripeschema.json
var schemaInit []byte

func (r *rampJSON) constantInit() error {
	if r.Depth == 0 {
		r.Depth = 12
	}

	if r.Maximum > int(math.Pow(2, float64(r.Depth))) {
		return fmt.Errorf("0121 The Max white value %v is higher than the maximum value of %v for a bit depth of %v", r.Maximum, math.Pow(2, float64(r.Depth)), r.Depth)
	}
	if r.Minimum > r.Maximum {
		return fmt.Errorf("0122 The black value %v is higher than the maximum white value of %v", r.Minimum, r.Maximum)
	}
	// Convert all constants to 12 bit as all calualtions are done with a 12 bit scale
	shift := int(math.Pow(2, float64(12-r.Depth)))
	if shift != 1 {
		r.Minimum = (shift * r.Minimum)
		r.Maximum = (shift * r.Maximum) + shift - 1 // This brings it to the effective 12 bit range where 12 bits are missed off
		// Rampstart = (shift * r.Rampstart) // Work this into the start point
	}

	return nil
}

func (r rampJSON) Alias() string {
	return r.GridLoc.Alias
}

func (r rampJSON) Location() string {
	return r.GridLoc.Location
}
