package sfm

type Session struct {
	mapName string
	clips map[*Clip]struct{}
}

func NewSession() *Session {
	return &Session{
	}
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
	return clip
}
