package sfm

import (
	"errors"
)

type Session struct {
	mapName    string
	clips      map[*Clip]struct{}
	activeClip *Clip
}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) GetMap() string {
	return s.mapName
}

func (s *Session) SetMap(name string) {
	s.mapName = name
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
