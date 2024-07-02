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

var animSetEditorChannels *sfm.Track

func CreateClip(session *sfm.Session) (*sfm.FilmClip, error) {
	clip := session.CreateFilmClip("SFM")

	sound := clip.CreateTrackGroup("Sound")
	clip.CreateTrackGroup("Overlay")

	sound.CreateTrack("Dialog", sfm.CLIP_SOUND)
	sound.CreateTrack("Music", sfm.CLIP_SOUND)
	sound.CreateTrack("Effects", sfm.CLIP_SOUND)

	clip.SubClipTrackGroup = sfm.NewTrackGroup("subClipTrackGroup")
	filmTrack := clip.SubClipTrackGroup.CreateTrack("Film", sfm.CLIP_FILM)

	shot1 := sfm.NewFilmClip("shot1")
	channelTrackGroup := shot1.CreateTrackGroup("channelTrackGroup")
	animSetEditorChannels = channelTrackGroup.CreateTrack("animSetEditorChannels", sfm.CLIP_CHANNEL)
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

func createAnimationSet(name string) *sfm.AnimationSet {
	animationSet := sfm.NewAnimationSet(name)

	channel := sfm.NewChannel("rootTransform_scale")

	animationSet.AddOperator(channel)

	return animationSet
}

func AddModel(clip *sfm.FilmClip, name string, filename string, f2 string) error {
	dag := sfm.NewNode(name)

	model1 := sfm.NewGameModel(name, filename)
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
	as := createAnimationSet(name)
	as.CreateTransformControl("rootTransform")

	channelsClip := sfm.NewChannelsClip(name)
	animSetEditorChannels.AddChildren(channelsClip)

	for k, v := range skel.GetBones() {
		bone := sfm.NewBone(fmt.Sprintf("bone %d (%s)", k, v.Name))
		bone.Transform.Position = v.PosParent
		bone.Transform.Orientation = v.RotParent
		bones[v] = bone
		model1.AddBone(bone)
		tc := as.CreateTransformControl(v.Name)

		tc.PositionChannel.ToElement = bone
		tc.PositionChannel.ToAttribute = "position"
		tc.OrientationChannel.ToElement = bone
		tc.OrientationChannel.ToAttribute = "orientation"

		channelsClip.AddChannel(&tc.OrientationChannel)
		channelsClip.AddChannel(&tc.PositionChannel)
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
