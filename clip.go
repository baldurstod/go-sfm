package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Clip struct {
	Name              string
	TimeFrame         *TimeFrame
	Color             Color
	Text              string
	Mute              bool
	trackGroups       []*TrackGroup
	DisplayScale      float32
	GlobalState       Element
	FadeIn            Time
	FadeOut           Time
	BackgroundColor   Color
	BackgroundFXClip  Element
	SubClipTrackGroup *TrackGroup
}

func newClip(name string) *Clip {
	return &Clip{
		Name:            name,
		TimeFrame:       newTimeFrame(),
		Color:           [...]byte{0, 0, 0, 0},
		BackgroundColor: [...]byte{64, 64, 64, 255},
		trackGroups:     make([]*TrackGroup, 0),
	}
}

func (c *Clip) AddTrackGroup(tg *TrackGroup) {
	c.trackGroups = append(c.trackGroups, tg)
}

func (c *Clip) CreateTrackGroup(name string) *TrackGroup {
	tg := NewTrackGroup(name)
	c.trackGroups = append(c.trackGroups, tg)
	return tg
}

func (c *Clip) toDmElement(serializer *Serializer) *dmx.DmElement {
	panic("do not use toDmElement, derived structs must use getDmElement instead")
}

func (c *Clip) getDmElement(serializer *Serializer, elementType string) *dmx.DmElement {
	e := dmx.NewDmElement(c.Name, elementType)

	e.CreateElementAttribute("timeFrame", serializer.GetElement(c.TimeFrame))
	e.CreateColorAttribute("color", c.Color)
	e.CreateStringAttribute("text", c.Text)
	e.CreateBoolAttribute("mute", c.Mute)

	trackGroups := e.CreateAttribute("trackGroups", dmx.AT_ELEMENT_ARRAY)
	for _, tg := range c.trackGroups {
		trackGroups.PushElement(serializer.GetElement(tg))
	}

	if c.SubClipTrackGroup != nil {
		e.CreateElementAttribute("subClipTrackGroup", serializer.GetElement(c.SubClipTrackGroup))
	} else {
		e.CreateElementAttribute("subClipTrackGroup", nil)
	}

	return e
}
