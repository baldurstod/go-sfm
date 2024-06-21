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

	dialog := sound.CreateTrack("Dialog")
	dialog.ClipType = sfm.CLIP_SOUND
	music := sound.CreateTrack("Music")
	music.ClipType = sfm.CLIP_SOUND
	effects := sound.CreateTrack("Effects")
	effects.ClipType = sfm.CLIP_SOUND

	clip.SubClipTrackGroup = sfm.NewTrackGroup("subClipTrackGroup")
	filmTrack := clip.SubClipTrackGroup.CreateTrack("Film")
	filmTrack.ClipType = sfm.CLIP_FILM

	shot1 := sfm.NewClip("shot1")
	filmTrack.AddChildren(shot1)
	shot1.Camera = sfm.NewCamera("camera1")

	shot1.Scene = createScene()
	//shot1.MapName = "maps/dota.vmap"

	shot1.AddAnimationSet(createAnimationSet())

	//log.Println(shot1.Camera)

	return clip
}

func createScene() *sfm.Node {
	scene := sfm.NewNode("scene")

	model1 := sfm.NewGameModel("tiny_04", "models/heroes/tiny/tiny_04/tiny_04.vmdl")
	scene.AddChildren(model1)
	model1.CreateBone("test")

	model2 := sfm.NewGameModel("tiny_astral_order_weapon", "models/items/tiny/tiny_astral_order_weapon/tiny_astral_order_weapon.vmdl")
	scene.AddChildren(model2)
	model2.CreateBone("test")

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
