package framecount

import (
	_ "embed"

	"github.com/mmTristan/opentsg-core/config"
)

/*
// framecount definitions
const wName = "frame count"
const wType = "framecounter"
const wLibrary = "builtin"
const hooks = "[]"*/

type frameJSON struct {
	Type         string       `json:"type" yaml:"type"`
	FrameCounter bool         `json:"framecounter,omitempty" yaml:"framecounter,omitempty"`
	Imgpos       interface{}  `json:"gridPosition" yaml:"gridPosition"`
	TextColour   string       `json:"textcolor" yaml:"textcolor"`
	BackColour   string       `json:"backgroundcolor" yaml:"backgroundcolor"`
	Font         string       `json:"font" yaml:"font"`
	FontSize     float64      `json:"fontsize" yaml:"fontsize"`
	GridLoc      *config.Grid `json:"grid,omitempty" yaml:"grid,omitempty"`
	//	DesignScale  string       `json:"designScale" yaml:"designScale"`
	// This is added in for metadata purposes
	FrameNumber int `json:"framenumber"`
}

// start the count at -1 as it is incremented before being returned
var framecount = -1

//go:embed jsonschema/framecounter.json
var frameSchema []byte

func (f *frameJSON) getFrames() bool {
	if f.FrameCounter {
		framecount++
	}

	return f.FrameCounter
}

func framePos() int { return framecount }

func (f frameJSON) Alias() string {
	return f.GridLoc.Alias
}

func (f frameJSON) Location() string {
	return f.GridLoc.Location
}
