package sfm_test

import (
	"bytes"
	"fmt"
	"log"
	_ "log"
	"os"
	"path"
	"testing"

	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-source2-tools/model"
	"github.com/baldurstod/go-source2-tools/parser"
)

const varFolder = "./var/"

func TestSession(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	session := sfm.NewSession()

	_, err := createClip(session)
	if err != nil {
		t.Error(err)
		return
	}

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().GetElement(session))
	//log.Println(buf)

	os.WriteFile(path.Join(varFolder, "test_session.dmx"), buf.Bytes(), 0666)
}

func createClip(session *sfm.Session) (*sfm.Clip, error) {
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

	var err error
	scene, err := createScene()
	if err != nil {
		return nil, err
	}
	shot1.Scene = scene
	//shot1.MapName = "maps/dota.vmap"

	shot1.AddAnimationSet(createAnimationSet())

	//log.Println(shot1.Camera)

	return clip, nil
}

func createScene() (*sfm.Node, error) {
	scene := sfm.NewNode("scene")

	model1 := sfm.NewGameModel("tiny_04", "models/heroes/tiny/tiny_01/tiny_01.vmdl")
	scene.AddChildren(model1)

	s2Model, err := getModel("tiny_01.vmdl_c")
	//defer return scene, nil

	if err != nil {
		return scene, nil
	}

	log.Println(s2Model.GetSkeleton())
	skel, err := s2Model.GetSkeleton()
	if err != nil {
		return scene, nil
	}

	for k, v := range skel.GetBones() {
		bone := model1.CreateBone(fmt.Sprintf("bone %d (%s)", k, v.Name))
		bone.Transform.Position = v.PosParent
		bone.Transform.Orientation = v.RotParent
	}

	/*
		model2 := sfm.NewGameModel("tiny_astral_order_weapon", "models/items/tiny/tiny_astral_order_weapon/tiny_astral_order_weapon.vmdl")
		scene.AddChildren(model2)
		model2.CreateBone("test")*/

	return scene, nil
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

func getModel(filename string) (*model.Model, error) {
	b, err := os.ReadFile(path.Join(varFolder, filename))
	if err != nil {
		return nil, err
	}
	file, err := parser.Parse(b)
	if err != nil {
		return nil, err
	}

	model := model.NewModel()
	model.SetFile(file)

	return model, nil

	//log.Println(model.GetSkeleton())
}
