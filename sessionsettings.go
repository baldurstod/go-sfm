package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type SessionSettings struct {
	Name                      string
	TimeSelection             *TimeSelection
	GraphEditorState          *GraphEditorState
	ProceduralPresetSettings  *ProceduralPresetSettings
	RenderSettings            *RenderSettings
	PosterSettings            *PosterSettings
	MovieSettings             *MovieSettings
	SharedPresetGroupSettings *SharedPresetGroupSettings
}

func NewSessionSettings() *SessionSettings {
	return &SessionSettings{
		Name:                      "sessionSettings",
		TimeSelection:             NewTimeSelection(),
		GraphEditorState:          NewGraphEditorState(),
		ProceduralPresetSettings:  NewProceduralPresetSettings(),
		RenderSettings:            NewRenderSettings(),
		PosterSettings:            NewPosterSettings(),
		MovieSettings:             NewMovieSettings(),
		SharedPresetGroupSettings: NewSharedPresetGroupSettings(),
	}
}

func (s *SessionSettings) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(s.Name, "DmElement")
}

func (s *SessionSettings) isExportable() bool {
	return true
}

func (s *SessionSettings) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateElementAttribute("timeSelection", serializer.GetElement(s.TimeSelection))
	e.CreateElementAttribute("graphEditorState", serializer.GetElement(s.GraphEditorState))
	e.CreateElementAttribute("proceduralPresets", serializer.GetElement(s.ProceduralPresetSettings))
	e.CreateElementAttribute("renderSettings", serializer.GetElement(s.RenderSettings))
	e.CreateElementAttribute("posterSettings", serializer.GetElement(s.PosterSettings))
	e.CreateElementAttribute("movieSettings", serializer.GetElement(s.MovieSettings))
	e.CreateElementAttribute("sharedPresetGroupSettings", serializer.GetElement(s.SharedPresetGroupSettings))

	/*

		if s.activeClip != nil {
			e.CreateElementAttribute("activeClip", serializer.GetElement(s.activeClip))
		} else {
			e.CreateElementAttribute("activeClip", nil)

		}

		clipBin := e.CreateAttribute("clipBin", dmx.AT_ELEMENT_ARRAY)
		for k := range s.clips {
			clipBin.PushElement(serializer.GetElement(k))
		}
	*/
}
