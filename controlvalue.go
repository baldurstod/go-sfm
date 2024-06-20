package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type ControlValue struct {
	Name  string
	Value float32
}

func NewControlValue(name string, value float32) *ControlValue {
	return &ControlValue{
		Name:  name,
		Value: value,
	}
}

func (cv *ControlValue) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmElement")

	e.CreateStringAttribute("name", cv.Name)
	e.CreateFloatAttribute("value", cv.Value)

	return e
}

/*
	"DmElement"
	{
		"id" "elementid" "c5def79f-0f45-4146-996b-ebb81d859f60"
		"name" "string" "dimpler"
		"value" "float" "0.1400000006"
	},
*/
