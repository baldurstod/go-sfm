package sfm

import (
	"errors"
	"github.com/baldurstod/go-dmx"
)

type Session struct {
	Name       string
	clips      map[*Clip]struct{}
	activeClip *Clip
	settings   *SessionSettings
}

func NewSession() *Session {
	return &Session{
		Name:     "session",
		clips:    make(map[*Clip]struct{}),
		settings: NewSessionSettings(),
	}
}

func (s *Session) CreateClip(name string) *Clip {
	clip := NewClip(name)
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

func (s *Session) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(s.Name, "DmElement")

	if s.activeClip != nil {
		e.CreateElementAttribute("activeClip", serializer.GetElement(s.activeClip))
	} else {
		e.CreateElementAttribute("activeClip", nil)
	}
	e.CreateElementAttribute("settings", serializer.GetElement(s.settings))

	clipBin := e.CreateAttribute("clipBin", dmx.AT_ELEMENT_ARRAY)
	for k := range s.clips {
		clipBin.PushElement(serializer.GetElement(k))
	}

	return e
}
