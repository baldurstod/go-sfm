package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type SharedPresetGroupSettings struct {
	Name             string
	presetGroupInfos []*PresetGroupInfo
}

func NewSharedPresetGroupSettings() *SharedPresetGroupSettings {
	return &SharedPresetGroupSettings{
		Name:             "sharedPresetGroupSettings",
		presetGroupInfos: make([]*PresetGroupInfo, 0),
	}
}

func (spgs *SharedPresetGroupSettings) AddPresetGroupInfo(pgi *PresetGroupInfo) {
	spgs.presetGroupInfos = append(spgs.presetGroupInfos, pgi)
}

func (spgs *SharedPresetGroupSettings) CreatePresetGroupInfo(name string, filenameBase string) *PresetGroupInfo {
	pgi := NewPresetGroupInfo(name, filenameBase)
	spgs.presetGroupInfos = append(spgs.presetGroupInfos, pgi)
	return pgi
}

func (spgs *SharedPresetGroupSettings) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(spgs.Name, "DmElement")
}

func (spgs *SharedPresetGroupSettings) isExportable() bool {
	return true
}

func (spgs *SharedPresetGroupSettings) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	presetGroupInfos := e.CreateAttribute("presetGroupInfos", dmx.AT_ELEMENT_ARRAY)
	for _, pgi := range spgs.presetGroupInfos {
		presetGroupInfos.PushElement(serializer.GetElement(pgi))
	}
}
