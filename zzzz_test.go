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

	clip.AddTrackGroup(sfm.NewTrackGroup("Sound"))
	clip.AddTrackGroup(sfm.NewTrackGroup("Overlay"))

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().GetElement(session))
	log.Println(buf)
}

func TestCamera(t *testing.T) {
	//camera := sfm.Camera{}

	//log.Println(camera)
}
