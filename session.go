package sfm

import (
	"errors"
	"github.com/baldurstod/go-dmx"
)

type Session struct {
	clips      map[*Clip]struct{}
	activeClip *Clip
}

func NewSession() *Session {
	return &Session{}
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
	e := dmx.NewDmElement()
	return e
}
