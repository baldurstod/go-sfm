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

func (ges *GraphEditorState) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(ges.Name, "DmeGraphEditorState")

	e.CreateBoolAttribute("displayGrid", ges.DisplayGrid)

	return e
}

/*
	"graphEditorState" "DmeGraphEditorState"
	{
		"id" "elementid" "091a352e-281b-4574-9b52-ee8674aa7ee7"
		"name" "string" "graphEditorState"
		"displayGrid" "bool" "1"
	}

*/
