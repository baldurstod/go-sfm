package sfm

import (
	"bytes"
	"errors"
	"os"

	"github.com/baldurstod/go-dmx"
)

type Session struct {
	Name       string
	clips      map[IClip]struct{}
	activeClip IClip
	settings   *SessionSettings
}

func NewSession() *Session {
	return &Session{
		Name:     "session",
		clips:    make(map[IClip]struct{}),
		settings: NewSessionSettings(),
	}
}

func (s *Session) CreateFilmClip(name string) *FilmClip {
	clip := NewFilmClip(name)
	s.clips[clip] = struct{}{}

	if s.activeClip == nil {
		s.activeClip = clip
	}

	return clip
}

func (s *Session) SetActiveClip(clip IClip) error {
	_, ok := s.clips[clip]

	if !ok {
		return errors.New("Clip doesn't belong to this session")
	}

	s.activeClip = clip

	return nil
}

func (s *Session) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(s.Name, "DmElement")
}

func (s *Session) isExportable() bool {
	return true
}

func (s *Session) toDmElement(serializer *Serializer, e *dmx.DmElement) {
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
}

func (s *Session) WriteTextFile(name string) error {
	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, NewSerializer().Serialize(s))

	os.WriteFile(name, buf.Bytes(), 0666)
	return nil
}

func (s *Session) WriteBinaryFile(name string) error {
	buf := new(bytes.Buffer)
	dmx.SerializeBinary(buf, NewSerializer().Serialize(s), "sfm_session", 22)

	os.WriteFile(name, buf.Bytes(), 0666)
	return nil
}
