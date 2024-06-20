package sfm_test

import (
	"bytes"
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-sfm"
	"log"
	"testing"
)

func TestSession(t *testing.T) {
	session := sfm.NewSession()
	clip := session.CreateClip()

	sound := clip.CreateTrackGroup("Sound")
	clip.CreateTrackGroup("Overlay")

	sound.CreateTrack("Dialog")
	sound.CreateTrack("Music")
	sound.CreateTrack("Effects")

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().GetElement(session))
	log.Println(buf)
}

func TestCamera(t *testing.T) {
	//camera := sfm.Camera{}

	//log.Println(camera)
}
