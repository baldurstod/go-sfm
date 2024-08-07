package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Control struct {
	Name         string
	Value        float32
	DefaultValue float32
	Channel      Channel
}

func (*Control) isControl() {}

func NewControl(name string) *Control {
	c := &Control{
		Name:    name,
		Channel: *NewChannel[float32](name + "_flex_channel"),
	}

	c.Channel.FromElement = c
	c.Channel.FromAttribute = "value"
	c.Channel.Mode = 3

	return c
}

func (c *Control) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(c.Name, "DmElement")
}

func (c *Control) isExportable() bool {
	return true
}

func (c *Control) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateFloatAttribute("value", c.Value)
	e.CreateFloatAttribute("defaultValue", c.DefaultValue)
	e.CreateElementAttribute("channel", serializer.GetElement(&c.Channel))
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
