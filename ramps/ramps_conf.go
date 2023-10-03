package ramps

type Ramp struct {
	Gradients        groupContents
	Groups           []RampProperties
	WidgetProperties control
}

type groupContents struct {
	GroupSeparator    groupSeparator
	GradientSeparator gradientSeparator
	Gradients         []Gradient
}

type textObjectJSON struct {
	TextYPosition string  `json:"textyPosition" yaml:"textyPosition"`
	TextXPosition string  `json:"textxPosition" yaml:"textxPosition"`
	TextHeight    float64 `json:"textHeight" yaml:"textHeight"`
	TextColour    string  `json:"textColor" yaml:"textColor"`
}

type RampProperties struct {
	Colour            string
	InitialPixelValue int
	Reverse           bool
}
type Gradient struct {
	Height   int
	BitDepth int
	Label    string

	// things that are added on run throughs
	startPoint int
	reverse    bool

	// Things we generate
	base   control
	colour string
}

type groupSeparator struct {
	Height int
	Colour string
}

type gradientSeparator struct {
	Colours []string
	Height  int
	// things the user does not assign
	base control
	step int
}

type control struct {
	MaxBitDepth      int
	CwRotation       string
	ObjectFitFill    bool
	PixelValueRepeat int
	TextProperties   textObjectJSON
	// These are things the user does not set
	/*
		fill function - for rotation to automatically translate the fill location
		fill - get stepsize and end goal

		step size - fill or truncate. Add a multiplier


	*/

	angleType      string
	truePixelShift float64
}
