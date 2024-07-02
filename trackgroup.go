package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type TrackGroup struct {
	Name            string
	tracks          []*Track
	Visible         bool
	Mute            bool
	DisplayScale    float32
	Minimized       bool
	Volume          float32
	ForceMultiTrack bool
}

func NewTrackGroup(name string) *TrackGroup {
	return &TrackGroup{
		Name:         name,
		Visible:      true,
		DisplayScale: 1,
	}
}

func (tg *TrackGroup) AddTrack(track *Track) {
	tg.tracks = append(tg.tracks, track)
}

func (tg *TrackGroup) CreateTrack(name string, clipType DmeClipType) *Track {
	track := NewTrack(name, clipType)
	tg.tracks = append(tg.tracks, track)
	return track
}

func (tg *TrackGroup) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(tg.Name, "DmeTrackGroup")
}

func (tg *TrackGroup) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateBoolAttribute("mute", tg.Mute)
	e.CreateBoolAttribute("visible", tg.Visible)
	e.CreateBoolAttribute("minimized", tg.Minimized)
	e.CreateBoolAttribute("forcemultitrack", tg.ForceMultiTrack)
	e.CreateFloatAttribute("displayScale", tg.DisplayScale)
	e.CreateFloatAttribute("volume", tg.Volume)

	tracks := e.CreateAttribute("tracks", dmx.AT_ELEMENT_ARRAY)
	for _, tr := range tg.tracks {
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
