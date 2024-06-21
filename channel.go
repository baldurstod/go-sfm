package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Channel struct {
	Name          string
	FromElement   Element
	FromAttribute string
	FromIndex     int32
	ToElement     Element
	ToAttribute   string
	ToIndex       int32
	Mode          int32
	Log           Element
}

func NewChannel(name string) *Channel {
	return &Channel{
		Name: name,
	}
}

func (c *Channel) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(c.Name, "DmeChannel")

	e.CreateElementAttribute("fromElement", serializer.GetElement(c.FromElement))
	e.CreateStringAttribute("fromAttribute", c.FromAttribute)
	e.CreateIntAttribute("fromIndex", c.FromIndex)

	e.CreateElementAttribute("toElement", serializer.GetElement(c.ToElement))
	e.CreateStringAttribute("toAttribute", c.ToAttribute)
	e.CreateIntAttribute("toIndex", c.ToIndex)

	e.CreateIntAttribute("mode", c.Mode)
	e.CreateElementAttribute("log", serializer.GetElement(c.Log))

	return e
}

/*
"DmeChannel"
{
	"id" "elementid" "58bf2a7d-8fff-425f-9fde-2241804fa0ad"
	"name" "string" "innerBrowRaiser_flex_channel"
	"fromElement" "element" "5bb1548e-9693-4cca-9e6d-4d59754e9c37"
	"fromAttribute" "string" "value"
	"fromIndex" "int" "0"
	"toElement" "element" "66a1d5a0-8425-4d1f-b344-fdd442dcce4e"
	"toAttribute" "string" "flexWeight"
	"toIndex" "int" "0"
	"mode" "int" "3"
	"log" "DmeFloatLog"
	{
		"id" "elementid" "17ca9687-e075-4a97-8438-ba11d2c98530"
		"name" "string" "float log"
		"layers" "element_array"
		[
			"DmeFloatLogLayer"
			{
				"id" "elementid" "d04a49b1-ce07-4fd9-ae1f-342b6ff09023"
				"name" "string" "float log"
				"times" "time_array"
				[
					"0.0000"
				]
				"curvetypes" "int_array"
				[
				]
				"values" "float_array"
				[
					"0"
				]
				"compressed" "binary"
				"
				"
			}
		]
		"curveinfo" "element" ""
		"usedefaultvalue" "bool" "1"
		"defaultvalue" "float" "0"
		"bookmarks" "time_array"
		[
		]
	}

}
*/
