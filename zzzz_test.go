package sfm_test

import (
	"bytes"
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-sfm"
	_ "log"
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

	shot1.AddAnimationSet(createAnimationSet())

	//log.Println(shot1.Camera)

	return clip
}

func createScene() *sfm.Node {
	scene := sfm.NewNode("scene")

	model := sfm.NewGameModel("tiny_04", "models/heroes/tiny/tiny_04/tiny_04.vmdl")

	scene.AddChildren(model)

	return scene
}

func createAnimationSet() *sfm.AnimationSet {
	animationSet := sfm.NewAnimationSet("test_AnimationSet")

	channel := sfm.NewChannel("rootTransform_scale")

	animationSet.AddOperator(channel)

	return animationSet
}

func TestCamera(t *testing.T) {
	//camera := sfm.Camera{}

	//log.Println(camera)
}
