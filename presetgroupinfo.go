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

func (pgi *PresetGroupInfo) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(pgi.Name, "DmePresetGroupInfo")
}

func (pgi *PresetGroupInfo) isExportable() bool {
	return true
}

func (pgi *PresetGroupInfo) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	presetGroupInfos := e.CreateAttribute("presetGroupInfos", dmx.AT_ELEMENT_ARRAY)
	for _, pgi := range pgi.presetGroups {
		presetGroupInfos.PushElement(serializer.GetElement(pgi))
	}
}
