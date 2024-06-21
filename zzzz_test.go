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

	createClip(session)

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().GetElement(session))
	//log.Println(buf)

	os.WriteFile(path.Join(outputFolder, "test_session.dmx"), buf.Bytes(), 0666)
}

func createClip(session *sfm.Session) *sfm.Clip {
	clip := session.CreateClip("SFM")

	sound := clip.CreateTrackGroup("Sound")
	clip.CreateTrackGroup("Overlay")

	sound.CreateTrack("Dialog")
	sound.CreateTrack("Music")
	sound.CreateTrack("Effects")

	clip.SubClipTrackGroup = sfm.NewTrackGroup("subClipTrackGroup")
	filmTrack := clip.SubClipTrackGroup.CreateTrack("Film")

	shot1 := sfm.NewClip("shot1")
	filmTrack.AddChildren(shot1)
	shot1.Camera = sfm.NewCamera("camera1")

	shot1.Scene = createScene()

	log.Println(shot1.Camera)

	return clip
}

func createScene() *sfm.Scene {
	scene := sfm.NewScene("scene")

	return scene
}

func TestCamera(t *testing.T) {
	//camera := sfm.Camera{}

	//log.Println(camera)
}
