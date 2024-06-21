package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type ProceduralPresetSettings struct {
	Name              string
	JitterScale       float32
	SmoothScale       float32
	JitterScaleVector float32
	SmoothScaleVector float32
	JitterIterations  int32
	SmoothIterations  int32
	StaggerInterval   Time
}

func NewProceduralPresetSettings() *ProceduralPresetSettings {
	return &ProceduralPresetSettings{
		Name:              "proceduralPresets",
		JitterScale:       1,
		SmoothScale:       1,
		JitterScaleVector: 2.5,
		SmoothScaleVector: 2.5,
		JitterIterations:  5,
		SmoothIterations:  5,
		StaggerInterval:   0.0833,
	}
}

func (pps *ProceduralPresetSettings) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmeProceduralPresetSettings")

	e.CreateStringAttribute("name", pps.Name)
	e.CreateFloatAttribute("jitterscale", pps.JitterScale)
	e.CreateFloatAttribute("smoothscale", pps.SmoothScale)
	e.CreateFloatAttribute("jitterscale_vector", pps.JitterScaleVector)
	e.CreateFloatAttribute("smoothscale_vector", pps.SmoothScaleVector)
	e.CreateIntAttribute("jitteriterations", pps.JitterIterations)
	e.CreateIntAttribute("smoothiterations", pps.JitterIterations)
	e.CreateTimeAttribute("staggerinterval", pps.StaggerInterval)

	return e
}

/*
	"proceduralPresets" "DmeProceduralPresetSettings"
	{
		"id" "elementid" "3f866500-9e1f-4f27-99eb-2a4f15f1e9ae"
		"name" "string" "proceduralPresets"
		"jitterscale" "float" "1"
		"smoothscale" "float" "1"
		"jitterscale_vector" "float" "2.5"
		"smoothscale_vector" "float" "2.5"
		"jitteriterations" "int" "5"
		"smoothiterations" "int" "5"
		"staggerinterval" "time" "0.0833"
	}

*/
