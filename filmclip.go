package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type FilmClip struct {
	*Clip
	animationSets     map[*AnimationSet]struct{}
	bookmarkSets      []*BookmarkSet
	ActiveBookmarkSet int32
}

func (*FilmClip) isClip() {}

func NewFilmClip(name string) *FilmClip {
	return &FilmClip{
		Clip:              newClip(name),
		animationSets:     make(map[*AnimationSet]struct{}),
		bookmarkSets:      make([]*BookmarkSet, 0),
		ActiveBookmarkSet: -1,
	}
}

func (fc *FilmClip) AddAnimationSet(as *AnimationSet) {
	fc.animationSets[as] = struct{}{}
}

func (fc *FilmClip) CreateAnimationSet(name string) *AnimationSet {
	as := NewAnimationSet(name)
	fc.AddAnimationSet(as)
	return as
}

func (fc *FilmClip) createDmElement(serializer *Serializer) *dmx.DmElement {
	return fc.Clip.getDmElement(serializer, "DmeFilmClip")
}

func (fc *FilmClip) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	animationSets := e.CreateAttribute("animationSets", dmx.AT_ELEMENT_ARRAY)
	for as := range fc.animationSets {
		animationSets.PushElement(serializer.GetElement(as))
	}

	e.CreateIntAttribute("activeBookmarkSet", fc.ActiveBookmarkSet)

	bookmarkSets := e.CreateAttribute("bookmarkSets", dmx.AT_ELEMENT_ARRAY)
	for _, bs := range fc.bookmarkSets {
		bookmarkSets.PushElement(serializer.GetElement(bs))
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
