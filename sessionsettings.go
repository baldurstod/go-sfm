package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type SessionSettings struct {
	Name             string
	TimeSelection    *TimeSelection
	GraphEditorState *GraphEditorState
}

func NewSessionSettings() *SessionSettings {
	return &SessionSettings{
		Name:             "sessionSettings",
		TimeSelection:    NewTimeSelection(),
		GraphEditorState: NewGraphEditorState(),
	}
}

func (s *SessionSettings) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmElement")

	e.CreateStringAttribute("name", s.Name)
	e.CreateElementAttribute("timeSelection", serializer.GetElement(s.TimeSelection))
	e.CreateElementAttribute("graphEditorState", serializer.GetElement(s.GraphEditorState))

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

	return e
}
