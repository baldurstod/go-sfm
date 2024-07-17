package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type ChannelsClip struct {
	*Clip
	channels map[*Channel]struct{}
}

func (*ChannelsClip) isClip() {}

func NewChannelsClip(name string) *ChannelsClip {
	return &ChannelsClip{
		Clip:     newClip(name),
		channels: make(map[*Channel]struct{}),
	}
}

func (cc *ChannelsClip) AddChannel(ch *Channel) {
	cc.channels[ch] = struct{}{}
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
}

/*

	"DmeChannelsClip"
	{
		"id" "elementid" "e9068407-be37-4484-8540-8f7b9e78255b"
		"name" "string" "tiny_04_1"
		"timeFrame" "DmeTimeFrame"
		{
			"id" "elementid" "ee52a31e-01e0-47ee-9e64-e3cc595c60f0"
			"start" "time" "-5.0000"
			"duration" "time" "70.0000"
			"offset" "time" "0.0000"
			"scale" "float" "1"
		}

		"color" "color" "0 0 0 0"
		"text" "string" ""
		"mute" "bool" "0"
		"trackGroups" "element_array"
		[
		]
		"displayScale" "float" "1"
		"channels" "element_array"
		[
			"element" "25f69b5e-cb37-44f0-945f-1b042bf6287c",
			"DmeChannel"
			{
				"id" "elementid" "13e948fe-1f0d-4732-9535-c7c451e291dd"
				"name" "string" "scaled_rootTransform_scale_channel"
				"fromElement" "element" "195243ee-5973-42b9-9fd3-06b351e861b4"
				"fromAttribute" "string" "result"
				"fromIndex" "int" "0"
				"toElement" "element" "9168f52d-a1ab-47f5-8e90-4b09e5c88eba"
				"toAttribute" "string" "scale"
				"toIndex" "int" "0"
				"mode" "int" "1"
				"log" "DmeFloatLog"
				{
					"id" "elementid" "a4215837-2231-43d8-8dcb-71ff5cf2bd1c"
					"name" "string" "float log"
					"layers" "element_array"
					[
						"DmeFloatLogLayer"
						{
							"id" "elementid" "71fe343d-53fe-4157-81ad-4b6d9540f0e4"
							"name" "string" "float log"
							"times" "time_array"
							[
							]
							"curvetypes" "int_array"
							[
							]
							"values" "float_array"
							[
							]
							"compressed" "binary"
							"
							"
						}
					]
					"curveinfo" "element" ""
					"usedefaultvalue" "bool" "0"
					"defaultvalue" "float" "0"
					"bookmarks" "time_array"
					[
					]
				}

			}
		]
	},
*/
