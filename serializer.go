package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Serializer struct {
	elements map[Element]*dmx.DmElement
	q        []Element
}

func NewSerializer() *Serializer {
	return &Serializer{
		elements: make(map[Element]*dmx.DmElement),
	}
}

func (s *Serializer) Serialize(element Element) *dmx.DmElement {
	s.q = make([]Element, 0, 100)
	e := s.GetElement(element)

	for {
		l := len(s.q) - 1
		if l < 0 {
			break
		}

		k := s.q[l]
		s.q = s.q[:l]

		k.toDmElement(s, s.elements[k])
	}

	return e
}

func (s *Serializer) GetElement(element Element) *dmx.DmElement {
	if element == nil {
		return nil
	}
	if !element.isExportable() {
		return nil
	}
	e, ok := s.elements[element]
	if ok {
		return e
	}

	e = element.createDmElement(s)
	s.elements[element] = e
	s.q = append(s.q, element)

	return e
}
