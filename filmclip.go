package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type FilmClip struct {
	*Clip
	Scene             INode
	Camera            *Camera
	MapName           string
	animationSets     map[*AnimationSet]struct{}
	bookmarkSets      []*BookmarkSet
	ActiveBookmarkSet int32
	operators         []Element
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

/*
func (fc *FilmClip) CreateAnimationSet(name string) *AnimationSet {
	as := newAnimationSet(name)
	fc.AddAnimationSet(as)
	return as
}
*/

func (fc *FilmClip) createDmElement(serializer *Serializer) *dmx.DmElement {
	return fc.Clip.getDmElement(serializer, "DmeFilmClip")
}

func (fc *FilmClip) isExportable() bool {
	return true
}

func (fc *FilmClip) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateElementAttribute("scene", serializer.GetElement(fc.Scene))

	if fc.Camera != nil {
		e.CreateElementAttribute("camera", serializer.GetElement(fc.Camera))
	} else {
		e.CreateElementAttribute("camera", nil)
	}

	e.CreateStringAttribute("mapname", fc.MapName)

	animationSets := e.CreateAttribute("animationSets", dmx.AT_ELEMENT_ARRAY)
	for as := range fc.animationSets {
		animationSets.PushElement(serializer.GetElement(as))
	}

	e.CreateIntAttribute("activeBookmarkSet", fc.ActiveBookmarkSet)

	bookmarkSets := e.CreateAttribute("bookmarkSets", dmx.AT_ELEMENT_ARRAY)
	for _, bs := range fc.bookmarkSets {
		bookmarkSets.PushElement(serializer.GetElement(bs))
	}

	operators := e.CreateAttribute("operators", dmx.AT_ELEMENT_ARRAY)
	for _, o := range fc.operators {
		operators.PushElement(serializer.GetElement(o))
	}
}

func (fc *FilmClip) CreateAnimationSetForModel(name string, filename string, parent INode) *AnimationSet {
	as := CreateAnimationSetForModel(name, filename)

	dag := NewNode(name)

	dag.AddChildren(as.GetGameModel())
	if parent != nil {
		parent.AddChildren(dag)
	}

	fc.animationSets[as] = struct{}{}

	return as
}

func (fc *FilmClip) CreateAnimationSetForParticleSystem(name string, filename string, parent INode) *AnimationSet {
	as := CreateAnimationSetForParticleSystem(name, filename)

	dag := NewNode(name)

	dag.AddChildren(as.GetParticleSystem())
	if parent != nil {
		parent.AddChildren(dag)
	}

	fc.animationSets[as] = struct{}{}

	return as
}

func (fc *FilmClip) GetDuration() float32 {
	return fc.TimeFrame.GetDuration()
}
