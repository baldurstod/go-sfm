package utils

import (
	"fmt"
	"os"
	"path"

	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-source2-tools/model"
	"github.com/baldurstod/go-source2-tools/parser"
)

const varFolder = "./var/"

func CreateClip(session *sfm.Session) (*sfm.Clip, error) {
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

	//shot1.AddAnimationSet(createAnimationSet())

	//log.Println(shot1.Camera)

	return shot1, nil
}

func createScene() (*sfm.Node, error) {
	scene := sfm.NewNode("scene")

	return scene, nil
}

func createAnimationSet() *sfm.AnimationSet {
	animationSet := sfm.NewAnimationSet("test_AnimationSet")

	channel := sfm.NewChannel("rootTransform_scale")

	animationSet.AddOperator(channel)

	return animationSet
}

func AddModel(clip *sfm.Clip, filename string, f2 string) error {
	dag := sfm.NewNode("model dag")

	model1 := sfm.NewGameModel("model", filename)
	dag.AddChildren(model1)
	clip.Scene.AddChildren(dag)

	s2Model, err := getModel(f2)

	if err != nil {
		return err
	}

	skel, err := s2Model.GetSkeleton()

	if err != nil {
		return err
	}

	bones := make(map[*model.Bone]*sfm.Bone)
	as := createAnimationSet()

	for k, v := range skel.GetBones() {
		bone := sfm.NewBone(fmt.Sprintf("bone %d (%s)", k, v.Name))
		bone.Transform.Position = v.PosParent
		bone.Transform.Orientation = v.RotParent
		bones[v] = bone
		model1.AddBone(bone)
		as.CreateTransformControl(v.Name)
	}

	for _, v := range skel.GetBones() {
		bone := bones[v]
		if v.ParentBone == nil {
			model1.AddChildren(bone)
		} else {
			bones[v.ParentBone].AddChildren(bone)
		}
	}

	clip.AddAnimationSet(as)
	as.GameModel = model1

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
