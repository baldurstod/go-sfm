package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Serializable interface {
	toDmElement(*Serializer) *dmx.DmElement
}

type Serializer struct {
	elements map[Serializable]*dmx.DmElement
}

func NewSerializer() *Serializer {
	return &Serializer{
		elements: make(map[Serializable]*dmx.DmElement),
	}
}

func (s *Serializer) GetElement(serializable Serializable) *dmx.DmElement {
	e, ok := s.elements[serializable]
	if ok {
		return e
	}

	e = serializable.toDmElement(s)
	s.elements[serializable] = e
	/*
		clip := newClip()
		s.clips[clip] = struct{}{}

		if s.activeClip == nil {
			s.activeClip = clip
		}
	*/

	return e
}
