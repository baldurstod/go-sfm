package sfm_test

import (
	"bytes"
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-sfm"
	"log"
	"os"
	"path"
	"testing"
)

const outputFolder = "./var/"

func TestSession(t *testing.T) {
	session := sfm.NewSession()

	clip := session.CreateClip("SFM")

	sound := clip.CreateTrackGroup("Sound")
	clip.CreateTrackGroup("Overlay")

	sound.CreateTrack("Dialog")
	sound.CreateTrack("Music")
	sound.CreateTrack("Effects")

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().GetElement(session))
	log.Println(buf)

	os.WriteFile(path.Join(outputFolder, "test_session.dmx"), buf.Bytes(), 0666)
}

func TestCamera(t *testing.T) {
	//camera := sfm.Camera{}

	//log.Println(camera)
}
