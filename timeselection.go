package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type TimeSelection struct {
	Name              string
	Enabled           bool
	Relative          bool
	FalloffLeft       Time
	FalloffRight      Time
	HoldLeft          Time
	HoldRight         Time
	InterpolatorLeft  int32
	InterpolatorRight int32
	Threshold         float32
	ResampleInterval  Time
	RecordingState    int32
}

func NewTimeSelection() *TimeSelection {
	return &TimeSelection{
		Name:              "timeSelection",
		Enabled:           true,
		FalloffLeft:       -214748.3647,
		FalloffRight:      214748.3647,
		HoldLeft:          -214748.3647,
		HoldRight:         214748.3647,
		InterpolatorLeft:  6,
		InterpolatorRight: 6,
		Threshold:         0.0005,
		ResampleInterval:  0.0100,
		RecordingState:    2,
	}
}

func (ts *TimeSelection) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(ts.Name, "DmeTimeSelection")
}

func (ts *TimeSelection) isExportable() bool {
	return true
}

func (ts *TimeSelection) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateBoolAttribute("enabled", ts.Enabled)
	e.CreateBoolAttribute("relative", ts.Relative)
	e.CreateTimeAttribute("falloff_left", ts.FalloffLeft)
	e.CreateTimeAttribute("falloff_right", ts.FalloffRight)
	e.CreateTimeAttribute("hold_left", ts.HoldLeft)
	e.CreateTimeAttribute("hold_right", ts.HoldRight)
	e.CreateIntAttribute("interpolator_left", ts.InterpolatorLeft)
	e.CreateIntAttribute("interpolator_right", ts.InterpolatorRight)
	e.CreateFloatAttribute("threshold", ts.Threshold)
	e.CreateTimeAttribute("resampleinterval", ts.ResampleInterval)
	e.CreateIntAttribute("recordingstate", ts.RecordingState)
}

/*
	"timeSelection" "DmeTimeSelection"
	{
		"id" "elementid" "8ff082b5-a31d-46dd-94cc-26eb66b71d32"
		"name" "string" "timeSelection"
		"enabled" "bool" "1"
		"relative" "bool" "0"
		"falloff_left" "time" "-214748.3647"
		"falloff_right" "time" "214748.3647"
		"hold_left" "time" "-214748.3647"
		"hold_right" "time" "214748.3647"
		"interpolator_left" "int" "6"
		"interpolator_right" "int" "6"
		"threshold" "float" "0.0005"
		"resampleinterval" "time" "0.0100"
		"recordingstate" "int" "2"
	}

*/
