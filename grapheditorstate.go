package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type GraphEditorState struct {
	Name        string
	DisplayGrid bool
}

func NewGraphEditorState() *GraphEditorState {
	return &GraphEditorState{
		Name:        "graphEditorState",
		DisplayGrid: true,
	}
}

func (ges *GraphEditorState) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(ges.Name, "DmeGraphEditorState")
}

func (ges *GraphEditorState) isExportable() bool {
	return true
}

func (ges *GraphEditorState) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateBoolAttribute("displayGrid", ges.DisplayGrid)
}

/*
	"graphEditorState" "DmeGraphEditorState"
	{
		"id" "elementid" "091a352e-281b-4574-9b52-ee8674aa7ee7"
		"name" "string" "graphEditorState"
		"displayGrid" "bool" "1"
	}

*/
