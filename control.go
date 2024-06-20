package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Control struct {
	Name         string
	Value        float32
	DefaultValue float32
	Channel      *Channel
}

func NewControl(name string) *Control {
	return &Control{
		Name: name,
	}
}

func (c *Control) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmElement")

	e.CreateStringAttribute("name", c.Name)
	e.CreateFloatAttribute("value", c.Value)
	e.CreateFloatAttribute("defaultValue", c.DefaultValue)
	e.CreateElementAttribute("channel", serializer.GetElement(c.Channel))

	return e
}

/*
"DmElement"
{
	"id" "elementid" "5bb1548e-9693-4cca-9e6d-4d59754e9c37"
	"name" "string" "innerBrowRaiser"
	"defaultValue" "float" "0"
	"value" "float" "0"
	"channel" "element" "58bf2a7d-8fff-425f-9fde-2241804fa0ad"
}
*/
