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

func (ps *MovieSettings) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmElement")

	e.CreateStringAttribute("name", ps.Name)
	e.CreateIntAttribute("videoTarget", ps.VideoTarget)
	e.CreateIntAttribute("audioTarget", ps.AudioTarget)
	e.CreateBoolAttribute("stereoscopic", ps.Stereoscopic)
	e.CreateBoolAttribute("stereoSingleFile", ps.StereoSingleFile)
	e.CreateBoolAttribute("clearDecals", ps.ClearDecals)
	e.CreateIntAttribute("width", ps.Width)
	e.CreateIntAttribute("height", ps.Height)
	e.CreateStringAttribute("filename", ps.Filename)

	return e
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
