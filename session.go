package sfm

import (
	"errors"
	"github.com/baldurstod/go-dmx"
)

type Session struct {
	clips      map[*Clip]struct{}
	activeClip *Clip
	dmElement  *dmx.DmElement
}

func NewSession() *Session {
	return &Session{
		clips: make(map[*Clip]struct{}),
	}
}

func (s *Session) CreateClip() *Clip {
	clip := newClip()
	s.clips[clip] = struct{}{}

	if s.activeClip == nil {
		s.activeClip = clip
	}

	return clip
}

func (s *Session) SetActiveClip(clip *Clip) error {
	_, ok := s.clips[clip]

	if !ok {
		return errors.New("Clip doesn't belong to this session")
	}

	s.activeClip = clip

	return nil
}

func (s *Session) ToDmElement() *dmx.DmElement {
	if s.dmElement == nil {
		s.dmElement = dmx.NewDmElement("DmElement")
	}
	e := s.dmElement

	if s.activeClip != nil {
		e.CreateElementAttribute("activeClip", s.activeClip.ToDmElement())
	} else {
		e.CreateElementAttribute("activeClip", nil)

	}

	clipBin := e.CreateAttribute("clipBin", dmx.AT_ELEMENT_ARRAY)
	for k := range s.clips {
		clipBin.PushElement(k.ToDmElement())
	}

	return e
}
