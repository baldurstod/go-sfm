package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Serializer struct {
	elements map[Element]*dmx.DmElement
}

func NewSerializer() *Serializer {
	return &Serializer{
		elements: make(map[Element]*dmx.DmElement),
	}
}

func (s *Serializer) GetElement(element Element) *dmx.DmElement {
	if element == nil {
		return nil
	}
	e, ok := s.elements[element]
	if ok {
		return e
	}

	e = element.toDmElement(s)
	s.elements[element] = e

	return e
}
