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

func (cv *ControlValue) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(cv.Name, "DmElement")
}

func (cv *ControlValue) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateFloatAttribute("value", cv.Value)
}

/*
	"DmElement"
	{
		"id" "elementid" "c5def79f-0f45-4146-996b-ebb81d859f60"
		"name" "string" "dimpler"
		"value" "float" "0.1400000006"
	},
*/
