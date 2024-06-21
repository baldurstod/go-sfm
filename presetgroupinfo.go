package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type PresetGroupInfo struct {
	Name         string
	FilenameBase string
	presetGroups []*PresetGroup
}

func NewPresetGroupInfo(name string, filenameBase string) *PresetGroupInfo {
	return &PresetGroupInfo{
		Name:         name,
		FilenameBase: filenameBase,
		presetGroups: make([]*PresetGroup, 0),
	}
}

func (pgi *PresetGroupInfo) AddPresetGroup(pg *PresetGroup) {
	pgi.presetGroups = append(pgi.presetGroups, pg)
}

func (pgi *PresetGroupInfo) CreatePresetGroup(name string) *PresetGroup {
	pg := NewPresetGroup(name)
	pgi.presetGroups = append(pgi.presetGroups, pg)
	return pg
}

func (pgi *PresetGroupInfo) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(pgi.Name, "DmePresetGroupInfo")

	presetGroupInfos := e.CreateAttribute("presetGroupInfos", dmx.AT_ELEMENT_ARRAY)
	for _, pgi := range pgi.presetGroups {
		presetGroupInfos.PushElement(serializer.GetElement(pgi))
	}

	return e
}

/*
	"DmePresetGroupInfo"
	{
		"id" "elementid" "85e4d9c3-7676-415d-a257-504beddf0545"
		"name" "string" "tiny_01"
		"filenameBase" "string" "models\\heroes\\tiny\\tiny_01\\tiny_01"
		"presetGroups" "element_array"
		[
			"DmePresetGroup"
			{
				"id" "elementid" "cd059cfb-0b7a-42ac-91c4-48348389a359"
				"name" "string" "phoneme"
				"presets" "element_array"
				[
					"DmePreset"
					{
						"id" "elementid" "6889ae5c-6150-4bde-8d7d-07d93ed5a5ae"
						"name" "string" "b"
						"controlValues" "element_array"
						[
							"DmElement"
							{
								"id" "elementid" "b7b2593b-a55d-44eb-92df-b3e792d75c37"
								"name" "string" "phonemeBMP"
								"value" "float" "1"
							},
							"DmElement"
							{
								"id" "elementid" "627bcfae-7d44-4c9b-981b-cf9520ada7ea"
								"name" "string" "alt_phonemeBMP"
								"value" "float" "1"
							}
						]
						"description" "string" "Big : voiced alveolar stop"
					}
				]
				"visible" "bool" "1"
				"readonly" "bool" "0"
			}
		]
	},

*/
