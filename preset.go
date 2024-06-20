package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Preset struct {
	Name          string
	controlValues []*ControlValue
	Description   string
}

func NewPreset(name string) *Preset {
	return &Preset{
		Name: name,
	}
}

func (p *Preset) AddControlValue(cv *ControlValue) {
	p.controlValues = append(p.controlValues, cv)
}

func (p *Preset) CreateControlValue(name string, value float32) *ControlValue {
	cv := NewControlValue(name, value)
	p.controlValues = append(p.controlValues, cv)
	return cv
}

func (p *Preset) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmePreset")

	e.CreateStringAttribute("name", p.Name)
	e.CreateStringAttribute("description", p.Description)

	controlValues := e.CreateAttribute("controlValues", dmx.AT_ELEMENT_ARRAY)
	for _, cv := range p.controlValues {
		controlValues.PushElement(serializer.GetElement(cv))
	}

	return e
}

/*
	"DmePreset"
	{
		"id" "elementid" "45aecc03-3eba-4eb8-827f-9de57ea0472b"
		"name" "string" "er"
		"controlValues" "element_array"
		[
			"DmElement"
			{
				"id" "elementid" "568d201f-9880-4e3a-9807-a9f3c0aa993d"
				"name" "string" "dimpler"
				"value" "float" "0.1400000006"
			},
			"DmElement"
			{
				"id" "elementid" "4649ed39-685e-4a08-bffa-b5c835a096ad"
				"name" "string" "upperLipsTowardAndPart"
				"value" "float" "0.0399999991"
			},
			"DmElement"
			{
				"id" "elementid" "fa52996d-a0b7-430f-8d01-2221db3fd2ad"
				"name" "string" "lowerLipsTowardAndPart"
				"value" "float" "0.0199999996"
			},
			"DmElement"
			{
				"id" "elementid" "fa45a511-3177-4f39-84d1-339bee6fc818"
				"name" "string" "lowerLipDepressorAndChinRaiser"
				"value" "float" "0"
			},
			"DmElement"
			{
				"id" "elementid" "67e153dd-40cd-4ccf-b5d8-ffdfd6ce3f48"
				"name" "string" "lipCornerDepressorAndSharpLipPuller"
				"value" "float" "0"
			},
			"DmElement"
			{
				"id" "elementid" "8160f1de-5e1c-48ef-bfd0-d03701097617"
				"name" "string" "lipCornerPuller"
				"value" "float" "0"
			},
			"DmElement"
			{
				"id" "elementid" "48559f42-8632-4b21-8994-a5dab547d01d"
				"name" "string" "lipStretcher"
				"value" "float" "0"
			},
			"DmElement"
			{
				"id" "elementid" "7a0e3fae-f832-4df8-b802-75d99ec66555"
				"name" "string" "lipPuckerer"
				"value" "float" "0.4799999893"
			},
			"DmElement"
			{
				"id" "elementid" "59957dc5-5818-4305-812a-4dd6c31b1f26"
				"name" "string" "upperLipFunneler"
				"value" "float" "0.5400000215"
			},
			"DmElement"
			{
				"id" "elementid" "802ab04c-382d-4deb-9be8-fa98f311c653"
				"name" "string" "lowerLipFunneler"
				"value" "float" "0.6299999952"
			},
			"DmElement"
			{
				"id" "elementid" "6c34119c-0206-40e4-a163-c6e713af5552"
				"name" "string" "jawSuckAndThrust"
				"value" "float" "0"
			},
			"DmElement"
			{
				"id" "elementid" "e51e4d3a-d971-4bec-b96b-1dc0bbbd4b75"
				"name" "string" "jawOpen"
				"value" "float" "0.2300000042"
			},
			"DmElement"
			{
				"id" "elementid" "886dc5a9-ac63-4db0-bf31-7310b7ffe2dc"
				"name" "string" "alt_dimpler"
				"value" "float" "0.1400000006"
			},
			"DmElement"
			{
				"id" "elementid" "afa78749-c7df-4d68-8b20-ef8d59c9c7b1"
				"name" "string" "alt_upperLipsTowardAndPart"
				"value" "float" "0.0399999991"
			},
			"DmElement"
			{
				"id" "elementid" "952d07a3-d777-47d8-8ad6-0122451a155a"
				"name" "string" "alt_lowerLipsTowardAndPart"
				"value" "float" "0.0199999996"
			},
			"DmElement"
			{
				"id" "elementid" "e9cd57d4-3f0d-4554-977d-ffd6af92054f"
				"name" "string" "alt_lowerLipDepressorAndChinRaiser"
				"value" "float" "0"
			},
			"DmElement"
			{
				"id" "elementid" "fa4f63b8-eba9-4452-b57e-618231dedace"
				"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
				"value" "float" "0"
			},
			"DmElement"
			{
				"id" "elementid" "70bfc919-458c-4f99-9082-7f8c6419e12d"
				"name" "string" "alt_lipCornerPuller"
				"value" "float" "0"
			},
			"DmElement"
			{
				"id" "elementid" "ca769bbb-003c-4755-b545-010219dc7f8e"
				"name" "string" "alt_lipStretcher"
				"value" "float" "0"
			},
			"DmElement"
			{
				"id" "elementid" "3e7ed4f5-3aa4-4f5e-83ce-db751d1ae429"
				"name" "string" "alt_lipPuckerer"
				"value" "float" "0.4799999893"
			},
			"DmElement"
			{
				"id" "elementid" "bd0511f0-0158-451f-844b-564f4232b8a4"
				"name" "string" "alt_upperLipFunneler"
				"value" "float" "0.5400000215"
			},
			"DmElement"
			{
				"id" "elementid" "f4eedeb1-d52d-46fc-8766-e65756b32dc0"
				"name" "string" "alt_lowerLipFunneler"
				"value" "float" "0.6299999952"
			},
			"DmElement"
			{
				"id" "elementid" "a4f877f4-3363-4807-a6fc-d17b4c1ed47d"
				"name" "string" "alt_jawSuckAndThrust"
				"value" "float" "0"
			},
			"DmElement"
			{
				"id" "elementid" "17fb79e8-ca52-41bd-b0b4-aecad9df8ce1"
				"name" "string" "alt_jawOpen"
				"value" "float" "0.2300000042"
			},
			"DmElement"
			{
				"id" "elementid" "b890c642-9bc0-403c-a55c-555734eef460"
				"name" "string" "tongueUp"
				"value" "float" "0.8000000119"
			}
		]
		"description" "string" "URn : rhotacized schwa"
	},
*/
