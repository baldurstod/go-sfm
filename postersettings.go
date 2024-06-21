package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type PosterSettings struct {
	Name            string
	Width           int32
	Height          int32
	DPI             int32
	Units           int32
	ConstrainAspect bool
	HeightInPixels  bool
	WidthInPixels   bool
}

func NewPosterSettings() *PosterSettings {
	return &PosterSettings{
		Name:            "posterSettings",
		Width:           1920,
		Height:          1080,
		DPI:             300,
		ConstrainAspect: true,
		HeightInPixels:  true,
		WidthInPixels:   true,
	}
}

func (ps *PosterSettings) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(ps.Name, "DmElement")

	e.CreateIntAttribute("width", ps.Width)
	e.CreateIntAttribute("height", ps.Height)
	e.CreateIntAttribute("DPI", ps.DPI)
	e.CreateIntAttribute("units", ps.Units)
	e.CreateBoolAttribute("constrainAspect", ps.ConstrainAspect)
	e.CreateBoolAttribute("heightInPixels", ps.HeightInPixels)
	e.CreateBoolAttribute("widthInPixels", ps.WidthInPixels)

	return e
}

/*
	"posterSettings" "DmElement"
	{
		"id" "elementid" "352cde0f-a78b-4da9-95c2-c61a81f66b92"
		"name" "string" "posterSettings"
		"width" "int" "1920"
		"height" "int" "1080"
		"DPI" "int" "300"
		"units" "int" "0"
		"constrainAspect" "bool" "1"
		"heightInPixels" "bool" "1"
		"widthInPixels" "bool" "1"
	}

*/
