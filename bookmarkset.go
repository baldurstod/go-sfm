package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type BookmarkSet struct {
	Name             string
	controls         []Element
	Color            Color
	Text             string
	Mute             bool
	trackGroups      []*TrackGroup
	DisplayScale     float32
	mapName          string
	Camera           *Camera
	Scene            *Node
	GlobalState      Element
	FadeIn           Time
	FadeOut          Time
	BackgroundColor  Color
	BackgroundFXClip Element
	operators        []Element
	animationSets    []*AnimationSet
	bookmarkSets     []*BookmarkSet
}

func NewBookmarkSet(name string) *BookmarkSet {
	return &BookmarkSet{
		Name: name,
	}
}

func (c *BookmarkSet) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmeAnimationSet")
	/*
		e.CreateElementAttribute("timeFrame", serializer.GetElement(c.TimeFrame))
		e.CreateColorAttribute("color", c.Color)
		e.CreateStringAttribute("text", c.Text)
		e.CreateBoolAttribute("mute", c.Mute)

		trackGroups := e.CreateAttribute("trackGroups", dmx.AT_ELEMENT_ARRAY)
		for _, tg := range c.trackGroups {
			trackGroups.PushElement(serializer.GetElement(tg))
		}*/

	return e
}
