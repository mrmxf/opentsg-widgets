package ramps

type Ramp struct {
	StripeGroup      layout
	Stripes          []RampProperties
	WidgetProperties control
}

type layout struct {
	Header      internalHeader
	InterStripe gradientSeparator
	Ramp        []Gradient // just do the heights frst
}

type textObjectJSON struct {
	TextYPosition string  `json:"textyPosition" yaml:"textyPosition"`
	TextXPosition string  `json:"textxPosition" yaml:"textxPosition"`
	TextHeight    float64 `json:"textHeight" yaml:"textHeight"`
	TextColour    string  `json:"textColor" yaml:"textColor"`
}

type RampProperties struct {
	Colour     string
	StartPoint int
	Reverse    bool
}
type Gradient struct {
	Height   int
	BitDepth int
	Label    string

	// things that are added on run throughs
	startPoint int
	reverse    bool

	// Thigns we generate
	base   control
	colour string
}

type internalHeader struct {
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
