package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Track struct {
	Name         string
	children     []Element
	Collapsed    bool
	Mute         bool
	Synched      bool
	ClipType     DmeClipType
	Volume       float32
	DisplayScale float32
}

func NewTrack(name string, clipType DmeClipType) *Track {
	return &Track{
		Name:         name,
		Synched:      true,
		Volume:       1,
		DisplayScale: 1,
		ClipType:     clipType,
	}
}

func (t *Track) AddChildren(child Element) {
	t.children = append(t.children, child)
}

func (t *Track) AddChannelsClip(name string) *ChannelsClip {
	channelsClip := newChannelsClip(name)
	t.AddChildren(channelsClip)
	return channelsClip
}

func (t *Track) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(t.Name, "DmeTrack")
}

func (t *Track) isExportable() bool {
	return true
}

func (t *Track) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateBoolAttribute("synched", t.Synched)
	e.CreateIntAttribute("clipType", t.ClipType)
	e.CreateFloatAttribute("volume", t.Volume)
	e.CreateFloatAttribute("displayScale", t.DisplayScale)

	children := e.CreateAttribute("children", dmx.AT_ELEMENT_ARRAY)
	for _, child := range t.children {
		children.PushElement(serializer.GetElement(child))
	}
}
