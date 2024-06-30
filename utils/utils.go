package utils

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-source2-tools/model"
	"github.com/baldurstod/go-source2-tools/parser"
)

const varFolder = "./var/"

func CreateClip(session *sfm.Session) (*sfm.Clip, *sfm.Node, error) {
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
		return nil, nil, err
	}
	shot1.Scene = scene
	//shot1.MapName = "maps/dota.vmap"

	//shot1.AddAnimationSet(createAnimationSet())

	//log.Println(shot1.Camera)

	return clip, scene, nil
}

func createScene() (*sfm.Node, error) {
	scene := sfm.NewNode("scene")

	//model1 := sfm.NewGameModel("tiny_04", "models/heroes/tiny/tiny_01/tiny_01.vmdl")
	//scene.AddChildren(model1)

	return scene, nil
	/*
	   s2Model, err := getModel("tiny_01.vmdl_c")

	   	if err != nil {
	   		return scene, nil
	   	}

	   log.Println(s2Model.GetSkeleton())
	   skel, err := s2Model.GetSkeleton()

	   	if err != nil {
	   		return scene, nil
	   	}

	   bones := make(map[*model.Bone]*sfm.Bone)

	   	for k, v := range skel.GetBones() {
	   		bone := sfm.NewBone(fmt.Sprintf("bone %d (%s)", k, v.Name))
	   		bone.Transform.Position = v.PosParent
	   		bone.Transform.Orientation = v.RotParent
	   		bones[v] = bone
	   		model1.AddBone(bone)
	   	}

	   	for _, v := range skel.GetBones() {
	   		bone := bones[v]
	   		if v.ParentBone == nil {
	   			model1.AddChildren(bone)
	   		} else {
	   			bones[v.ParentBone].AddChildren(bone)
	   		}
	   	}

	   return scene, nil
	*/
}

func createAnimationSet() *sfm.AnimationSet {
	animationSet := sfm.NewAnimationSet("test_AnimationSet")

	channel := sfm.NewChannel("rootTransform_scale")

	animationSet.AddOperator(channel)

	return animationSet
}

func AddModel(parent *sfm.Node, filename string, f2 string) error {
	model1 := sfm.NewGameModel("model", "models/heroes/tiny/tiny_01/tiny_01.vmdl")
	parent.AddChildren(model1)

	s2Model, err := getModel(f2)

	if err != nil {
		return err
	}

	log.Println(s2Model.GetSkeleton())
	skel, err := s2Model.GetSkeleton()

	if err != nil {
		return err
	}

	bones := make(map[*model.Bone]*sfm.Bone)

	for k, v := range skel.GetBones() {
		bone := sfm.NewBone(fmt.Sprintf("bone %d (%s)", k, v.Name))
		bone.Transform.Position = v.PosParent
		bone.Transform.Orientation = v.RotParent
		bones[v] = bone
		model1.AddBone(bone)
	}

	for _, v := range skel.GetBones() {
		bone := bones[v]
		if v.ParentBone == nil {
			model1.AddChildren(bone)
		} else {
			bones[v.ParentBone].AddChildren(bone)
		}
	}

	return nil
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
