package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type ProgressiveRefinementSettings struct {
	Name                             string
	On                               bool
	UseDepthOfField                  bool
	OverrideDepthOfFieldSamples      bool
	OverrideDepthOfFieldSamplesValue int32
	UseMotionBlur                    bool
	OverrideMotionBlurSamples        bool
	OverrideMotionBlurSamplesValue   int32
	UseAntialiasing                  bool
	OverrideShutterSpeed             bool
	OverrideShutterSpeedValue        float32
}

func NewProgressiveRefinementSettings() *ProgressiveRefinementSettings {
	return &ProgressiveRefinementSettings{
		Name:                             "ProgressiveRefinementSettings",
		On:                               true,
		UseDepthOfField:                  true,
		OverrideDepthOfFieldSamplesValue: 64,
		UseMotionBlur:                    true,
		OverrideMotionBlurSamplesValue:   8,
		UseAntialiasing:                  true,
		OverrideShutterSpeedValue:        1. / 48,
	}
}

func (prs *ProgressiveRefinementSettings) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(prs.Name, "DmElement")
}

func (prs *ProgressiveRefinementSettings) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateBoolAttribute("on", prs.On)
	e.CreateBoolAttribute("useDepthOfField", prs.UseDepthOfField)
	e.CreateBoolAttribute("overrideDepthOfFieldSamples", prs.OverrideDepthOfFieldSamples)
	e.CreateIntAttribute("overrideDepthOfFieldSamplesValue", prs.OverrideDepthOfFieldSamplesValue)
	e.CreateBoolAttribute("useMotionBlur", prs.UseMotionBlur)
	e.CreateBoolAttribute("overrideMotionBlurSamples", prs.OverrideMotionBlurSamples)
	e.CreateIntAttribute("overrideMotionBlurSamplesValue", prs.OverrideMotionBlurSamplesValue)
	e.CreateBoolAttribute("useAntialiasing", prs.UseAntialiasing)
	e.CreateBoolAttribute("overrideShutterSpeed", prs.OverrideShutterSpeed)
	e.CreateFloatAttribute("overrideShutterSpeedValue", prs.OverrideShutterSpeedValue)
}

/*
	"ProgressiveRefinement" "DmElement"
	{
		"id" "elementid" "b2add799-236e-4b46-980a-bac07c39679c"
		"name" "string" "ProgressiveRefinementSettings"
		"on" "bool" "1"
		"useDepthOfField" "bool" "1"
		"overrideDepthOfFieldSamples" "bool" "0"
		"overrideDepthOfFieldSamplesValue" "int" "64"
		"useMotionBlur" "bool" "1"
		"overrideMotionBlurSamples" "bool" "0"
		"overrideMotionBlurSamplesValue" "int" "8"
		"useAntialiasing" "bool" "1"
		"overrideShutterSpeed" "bool" "0"
		"overrideShutterSpeedValue" "float" "0.020833334"
	}

*/
