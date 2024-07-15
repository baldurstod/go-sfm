package utils

import (
	"fmt"
	"strings"

	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-source2-tools/model"
	"github.com/baldurstod/go-source2-tools/parser"
	"github.com/baldurstod/go-vector"
)

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

	channel := sfm.NewChannel[float32]("rootTransform_scale")
	animationSet.AddOperator(channel)

	return animationSet
}

func AddModel(clip *sfm.FilmClip, name string, repository string, filename string) (*sfm.AnimationSet, error) {
	filename = strings.TrimSuffix(filename, ".vmdl_c")
	filename = strings.TrimSuffix(filename, ".vmdl")
	filename += ".vmdl"

	dag := sfm.NewNode(name)

	model1 := sfm.NewGameModel(name, filename)
	dag.AddChildren(model1)
	clip.Scene.AddChildren(dag)

	s2Model, err := GetModel(repository, filename)

	if err != nil {
		return nil, err
	}

	skel, err := s2Model.GetSkeleton()

	if err != nil {
		return nil, err
	}

	bones := make(map[*model.Bone]*sfm.Bone)
	as := createAnimationSet(name)
	as.GameModel = model1

	channelsClip := sfm.NewChannelsClip(name)
	animSetEditorChannels.AddChildren(channelsClip)
	tc := as.CreateTransformControl("rootTransform")

	tc.PositionChannel.ToElement = model1.Transform
	tc.PositionChannel.ToAttribute = "position"
	tc.OrientationChannel.ToElement = model1.Transform
	tc.OrientationChannel.ToAttribute = "orientation"

	channelsClip.AddChannel(&tc.OrientationChannel)
	channelsClip.AddChannel(&tc.PositionChannel)

	posLayer := any(tc.PositionChannel.Log.GetLayer("vector3 log")).(*sfm.LogLayer[vector.Vector3[float32]])
	posLayer.SetValue(0, vector.Vector3[float32]{})
	rotLayer := any(tc.OrientationChannel.Log.GetLayer("quaternion log")).(*sfm.LogLayer[vector.Quaternion[float32]])
	rotLayer.SetValue(0, vector.Quaternion[float32]{})

	//layer := any(tc.PositionChannel.Log.AddLayer()).(*sfm.LogLayer[vector.Vector3[float32]])
	//log.Println(layer)

	/*for i := 0; i < 100; i++ {
		layer.AddValue(int32(i), vector.Vector3[float32]{float32(i), 0, 0})
	}*/

	for k, v := range skel.GetBones() {
		bone := sfm.NewBone(fmt.Sprintf("bone %d (%s)", k, v.Name))
		bone.Transform.Position = v.PosParent
		bone.Transform.Orientation = v.RotParent
		bones[v] = bone
		model1.AddBone(bone)
		tc := as.CreateTransformControl(v.Name)

		tc.PositionChannel.ToElement = bone.Transform
		tc.PositionChannel.ToAttribute = "position"
		tc.OrientationChannel.ToElement = bone.Transform
		tc.OrientationChannel.ToAttribute = "orientation"

		tc.ValuePosition = v.PosParent
		tc.ValueOrientation = v.RotParent

		channelsClip.AddChannel(&tc.OrientationChannel)
		channelsClip.AddChannel(&tc.PositionChannel)

		posLayer := any(tc.PositionChannel.Log.GetLayer("vector3 log")).(*sfm.LogLayer[vector.Vector3[float32]])
		posLayer.SetValue(0, bone.Transform.Position)
		rotLayer := any(tc.OrientationChannel.Log.GetLayer("quaternion log")).(*sfm.LogLayer[vector.Quaternion[float32]])
		rotLayer.SetValue(0, bone.Transform.Orientation)
	}

	for _, v := range skel.GetBones() {
		bone := bones[v]
		if v.ParentBone == nil {
			model1.AddChildren(bone)
		} else {
			bones[v.ParentBone].AddChildren(bone)
		}
	}

	err = initFlexes(as, s2Model, channelsClip)
	if err != nil {
		return nil, fmt.Errorf("failed to init flexes: <%w>", err)
	}

	clip.AddAnimationSet(as)

	return as, nil
}

var models = func() map[string]map[string]*model.Model { return make(map[string]map[string]*model.Model) }()

func findModel(repository string, filename string) *model.Model {
	r, ok := models[repository]
	if !ok {
		return nil
	}
	m, ok := r[filename]
	if !ok {
		return nil
	}
	return m
}

func addModel(repository string, filename string, m *model.Model) {
	r, ok := models[repository]
	if !ok {
		r = make(map[string]*model.Model)
		models[repository] = r
	}

	r[filename] = m
}

func GetModel(repository string, filename string) (*model.Model, error) {
	filename = strings.TrimSuffix(filename, ".vmdl_c")
	filename = strings.TrimSuffix(filename, ".vmdl")
	filename += ".vmdl_c"

	m := findModel(repository, filename)
	if m != nil {
		return m, nil
	}

	file, err := parser.Parse(repository, filename)
	if err != nil {
		return nil, err
	}

	m = model.NewModel()
	m.SetFile(file)

	addModel(repository, filename, m)

	return m, nil

	//log.Println(model.GetSkeleton())
}

func initFlexes(as *sfm.AnimationSet, s2Model *model.Model, channelsClip *sfm.ChannelsClip) error {
	flexes, err := s2Model.GetFlexes()
	if err != nil {
		return fmt.Errorf("failed to get model flexes: <%w>", err)
	}

	gameModel := as.GameModel
	for _, v := range flexes {
		ope := gameModel.CreateGlobalFlexControllerOperator(v.Name, 0.5)
		c := as.CreateControl(v.Name)
		c.Channel.ToElement = ope
		c.Channel.ToAttribute = "flexWeight"

		layer := any(c.Channel.Log.GetLayer("float log")).(*sfm.LogLayer[float32])
		layer.SetValue(0, 0.5)

		c.DefaultValue = 0.5
		c.Value = 0.5

		channelsClip.AddChannel(&c.Channel)
	}

	//	CreateGlobalFlexControllerOperator
	return nil
}
