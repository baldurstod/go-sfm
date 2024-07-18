package utils

import (
	"fmt"
	"strings"

	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-source2-tools/model"
	"github.com/baldurstod/go-source2-tools/parser"
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

/*
func createAnimationSet(name string) *sfm.AnimationSet {
	animationSet := sfm.NewAnimationSet(name)

	channel := sfm.NewChannel[float32]("rootTransform_scale")
	animationSet.AddOperator(channel)

	return animationSet
}*/

func AddModel(clip *sfm.FilmClip, name string, repository string, filename string) (*sfm.AnimationSet, error) {
	filename = strings.TrimSuffix(filename, ".vmdl_c")
	filename = strings.TrimSuffix(filename, ".vmdl")
	filename += ".vmdl"

	as := clip.CreateAnimationSetForModel(name, filename)

	s2Model, err := GetModel(repository, filename)

	if err != nil {
		return nil, err
	}

	skel, err := s2Model.GetSkeleton()

	if err != nil {
		return nil, err
	}

	bones := make(map[*model.Bone]*sfm.Bone)

	channelsClip := animSetEditorChannels.AddChannelsClip(name)
	channelsClip.AddAnimationSet(as)

	for k, v := range skel.GetBones() {
		bone := as.GetGameModel().CreateBone(as, v.Name, k, v.PosParent, v.RotParent)
		bones[v] = bone
	}

	for _, v := range skel.GetBones() {
		bone := bones[v]
		if v.ParentBone == nil {
			as.GetGameModel().AddChildren(bone)
		} else {
			bones[v.ParentBone].AddChildren(bone)
		}
	}

	err = initFlexes(as, s2Model, channelsClip)
	if err != nil {
		return nil, fmt.Errorf("failed to init flexes: <%w>", err)
	}

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

	gameModel := as.GetGameModel()
	for _, v := range flexes {
		defaultValue := v.GetDefaultValue()
		ope := gameModel.CreateGlobalFlexControllerOperator(v.Name, defaultValue)
		c := as.CreateControl(v.Name)
		c.Channel.ToElement = ope
		c.Channel.ToAttribute = "flexWeight"

		layer := any(c.Channel.Log.GetLayer("float log")).(*sfm.LogLayer[float32])
		layer.SetValue(0, defaultValue)

		c.DefaultValue = defaultValue
		c.Value = defaultValue

		channelsClip.AddChannel(&c.Channel)
	}

	//	CreateGlobalFlexControllerOperator
	return nil
}
