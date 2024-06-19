package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type TrackGroup struct {
	Name            string
	tracks          []*Track
	Visible         bool
	Mute            bool
	DisplayScale    float64
	Minimized       bool
	Volume          float64
	ForceMultiTrack bool
}

func NewTrackGroup(name string) *TrackGroup {
	return &TrackGroup{
		Name:         name,
		Visible:      true,
		DisplayScale: 1,
	}
}

func (t *TrackGroup) AddTrack(track *Track) {
	t.tracks = append(t.tracks, track)
}

func (t *TrackGroup) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmeTrackGroup")

	e.CreateStringAttribute("name", t.Name)
	e.CreateBoolAttribute("mute", t.Mute)
	e.CreateBoolAttribute("visible", t.Visible)
	e.CreateBoolAttribute("minimized", t.Minimized)
	e.CreateBoolAttribute("forcemultitrack", t.ForceMultiTrack)
	e.CreateFloatAttribute("displayScale", t.DisplayScale)
	e.CreateFloatAttribute("volume", t.Volume)

	tracks := e.CreateAttribute("tracks", dmx.AT_ELEMENT_ARRAY)
	for _, tr := range t.tracks {
		tracks.PushElement(serializer.GetElement(tr))
	}

	/*
		e.CreateElementAttribute("timeFrame", c.TimeFrame.ToDmElement())
		e.CreateColorAttribute("color", c.Color)
		e.CreateStringAttribute("text", c.Text)
		e.CreateBoolAttribute("mute", c.Mute)


		trackGroups := e.CreateAttribute("trackGroups", dmx.AT_ELEMENT_ARRAY)
		for _, tg := range c.trackGroups {
			trackGroups.PushElement(tg.ToDmElement())
		}
	*/

	return e
}

/*
	"DmeTrackGroup"
	{
		"id" "elementid" "da4acc8f-978b-4cc9-899d-da7298e4b61e"
		"name" "string" "Sound"
		"tracks" "element_array"
		[
			"DmeTrack"
			{
				"id" "elementid" "164d4b8c-5ddb-4a21-bc01-834680cc4a09"
				"name" "string" "Dialog"
				"children" "element_array"
				[
				]
				"collapsed" "bool" "0"
				"mute" "bool" "0"
				"synched" "bool" "1"
				"clipType" "int" "1"
				"volume" "float" "1"
				"displayScale" "float" "1"
			},
			"DmeTrack"
			{
				"id" "elementid" "635e17c5-5e52-4ae8-ae13-21f71475eb68"
				"name" "string" "Music"
				"children" "element_array"
				[
				]
				"collapsed" "bool" "0"
				"mute" "bool" "0"
				"synched" "bool" "1"
				"clipType" "int" "1"
				"volume" "float" "1"
				"displayScale" "float" "1"
			},
			"DmeTrack"
			{
				"id" "elementid" "8c6f4a78-a0fc-4fbd-b98e-ceab1ee70597"
				"name" "string" "Effects"
				"children" "element_array"
				[
				]
				"collapsed" "bool" "0"
				"mute" "bool" "0"
				"synched" "bool" "1"
				"clipType" "int" "1"
				"volume" "float" "1"
				"displayScale" "float" "1"
			}
		]
		"visible" "bool" "1"
		"mute" "bool" "0"
		"displayScale" "float" "1"
		"minimized" "bool" "0"
		"volume" "float" "1"
		"forcemultitrack" "bool" "0"
	},*/
