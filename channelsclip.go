package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type ChannelsClip struct {
	*Clip
	channels      map[*Channel]struct{}
	animationSets map[*AnimationSet]struct{}
}

func (*ChannelsClip) isClip() {}

func newChannelsClip(name string) *ChannelsClip {
	cc := ChannelsClip{
		Clip:          newClip(name),
		channels:      make(map[*Channel]struct{}),
		animationSets: make(map[*AnimationSet]struct{}),
	}

	cc.TimeFrame.Start = -5
	cc.TimeFrame.Duration = 70

	return &cc
}

func (cc *ChannelsClip) AddChannel(ch *Channel) {
	cc.channels[ch] = struct{}{}
}

func (cc *ChannelsClip) AddAnimationSet(as *AnimationSet) {
	cc.animationSets[as] = struct{}{}
}

func (cc *ChannelsClip) createDmElement(serializer *Serializer) *dmx.DmElement {
	return cc.Clip.getDmElement(serializer, "DmeChannelsClip")
}

func (cc *ChannelsClip) isExportable() bool {
	return true
}

func (cc *ChannelsClip) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	channels := e.CreateAttribute("channels", dmx.AT_ELEMENT_ARRAY)
	for ch := range cc.channels {
		channels.PushElement(serializer.GetElement(ch))
	}

	for as := range cc.animationSets {
		for _, ch := range as.getChannels() {
			channels.PushElement(serializer.GetElement(ch))
		}
	}
}
