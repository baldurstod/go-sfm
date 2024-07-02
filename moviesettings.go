package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type MovieSettings struct {
	Name             string
	VideoTarget      int32
	AudioTarget      int32
	Stereoscopic     bool
	StereoSingleFile bool
	ClearDecals      bool
	Width            int32
	Height           int32
	Filename         string
}

func NewMovieSettings() *MovieSettings {
	return &MovieSettings{
		Name:        "movieSettings",
		VideoTarget: 6,
		AudioTarget: 2,
		Width:       1280,
		Height:      720,
	}
}

func (ms *MovieSettings) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(ms.Name, "DmElement")
}

func (ms *MovieSettings) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateIntAttribute("videoTarget", ms.VideoTarget)
	e.CreateIntAttribute("audioTarget", ms.AudioTarget)
	e.CreateBoolAttribute("stereoscopic", ms.Stereoscopic)
	e.CreateBoolAttribute("stereoSingleFile", ms.StereoSingleFile)
	e.CreateBoolAttribute("clearDecals", ms.ClearDecals)
	e.CreateIntAttribute("width", ms.Width)
	e.CreateIntAttribute("height", ms.Height)
	e.CreateStringAttribute("filename", ms.Filename)
}

/*
	"movieSettings" "DmElement"
	{
		"id" "elementid" "659693c4-e148-44e6-8975-9b4120477f88"
		"name" "string" "movieSettings"
		"videoTarget" "int" "6"
		"audioTarget" "int" "2"
		"stereoscopic" "bool" "0"
		"stereoSingleFile" "bool" "0"
		"clearDecals" "bool" "0"
		"width" "int" "1280"
		"height" "int" "720"
		"filename" "string" ""
	}

*/
