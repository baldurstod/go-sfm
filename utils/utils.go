package utils

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-source2-tools/model"
	"github.com/baldurstod/go-source2-tools/parser"
	"github.com/baldurstod/go-source2-tools/particles"
	"github.com/baldurstod/go-vector"
)

var animSetEditorChannels *sfm.Track

func CreateClip(session *sfm.Session) *sfm.FilmClip {
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

	shot1.Scene = createScene()
	//shot1.MapName = "maps/dota.vmap"

	//shot1.AddAnimationSet(createAnimationSet())

	//log.Println(shot1.Camera)

	return shot1
}

func createScene() *sfm.Node {
	return sfm.NewNode("scene")
}

/*
func createAnimationSet(name string) *sfm.AnimationSet {
	animationSet := sfm.NewAnimationSet(name)

	channel := sfm.NewChannel[float32]("rootTransform_scale")
	animationSet.AddOperator(channel)

	return animationSet
}*/

func AddModel(clip *sfm.FilmClip, name string, repository string, filename string, parent sfm.INode) (*sfm.AnimationSet, error) {
	filename = strings.TrimSuffix(filename, ".vmdl_c")
	filename = strings.TrimSuffix(filename, ".vmdl")
	filename += ".vmdl"

	as := clip.CreateAnimationSetForModel(name, filename, parent)

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

	err = initFlexes(as, s2Model)
	if err != nil {
		return nil, fmt.Errorf("failed to init flexes: <%w>", err)
	}

	err = initAttachments(as, s2Model)
	if err != nil {
		return nil, fmt.Errorf("failed to init attachments: <%w>", err)
	}

	return as, nil
}

func PlayActivity(as *sfm.AnimationSet, activity *model.Activity, duration float32) error {
	if as == nil {
		return errors.New("empty animation set")
	}

	model := as.GetGameModel()
	if model == nil {
		return errors.New("animation set doesn't have any model")
	}

	s2Model, err := GetModel("dota2", model.ModelName)
	if err != nil {
		return err
	}

	seq, err := s2Model.GetSequence(activity)
	if err != nil {
		return err
	}
	playSequence(as, s2Model, seq, duration)

	log.Println(seq)

	return nil
}

func PlaySequence(as *sfm.AnimationSet, animation string, duration float32) error {
	if as == nil {
		return errors.New("empty animation set")
	}

	model := as.GetGameModel()
	if model == nil {
		return errors.New("animation set doesn't have any model")
	}

	s2Model, err := GetModel("dota2", model.ModelName)
	if err != nil {
		return err
	}

	seq, err := s2Model.GetSequenceByName(animation)
	if err != nil {
		return err
	}
	playSequence(as, s2Model, seq, duration)

	log.Println(seq)

	return nil
}

func playSequence(as *sfm.AnimationSet, model *model.Model, sequence *model.Sequence, duration float32) error {
	fps := sequence.GetFps()

	frames := int(duration * float32(fps))

	flexes, err := model.GetFlexes()
	if err != nil {
		return err
	}

	for frameId := 0; frameId < frames; frameId++ {
		frame, err := sequence.GetFrame(frameId)
		frameTime := float32(frameId) / float32(fps)
		if err != nil {
			return err
		}
		//log.Println(frame)
		positionChannel := frame.GetChannel("BoneChannel", "Position")
		for _, element := range positionChannel.Datas {

			tc := as.GetTransformControl(element.Name)
			if tc != nil {
				layer := any(tc.PositionChannel.Log.GetLayer("vector3 log")).(*sfm.LogLayer[vector.Vector3[float32]])
				layer.SetValue(frameTime, element.Datas.(vector.Vector3[float32]))
			}
		}
		orientationChannel := frame.GetChannel("BoneChannel", "Angle")
		for _, element := range orientationChannel.Datas {

			tc := as.GetTransformControl(element.Name)
			if tc != nil {
				layer := any(tc.OrientationChannel.Log.GetLayer("quaternion log")).(*sfm.LogLayer[vector.Quaternion[float32]])
				layer.SetValue(frameTime, (element.Datas.(vector.Quaternion[float32])))
			}
		}

		morphChannel := frame.GetChannel("MorphChannel", "data")
		for _, element := range morphChannel.Datas {

			value := float32(0)
			for _, flex := range flexes {
				if flex.Name == element.Name {
					value = flex.GetControllerValue((element.Datas.(float32))) //((element.Datas.(float32)) - flex.Min) / (flex.Max - flex.Min)
					break
				}
			}

			tc := as.GetControl(element.Name)
			if tc != nil {
				layer := any(tc.Channel.Log.GetLayer("float log")).(*sfm.LogLayer[float32])
				layer.SetValue(frameTime, value)
			}
		}
	}

	return nil
}

var models = func() map[string]map[string]any { return make(map[string]map[string]any) }()

func findModel(repository string, filename string) any {
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

func addModel(repository string, filename string, m any) {
	r, ok := models[repository]
	if !ok {
		r = make(map[string]any)
		models[repository] = r
	}

	r[filename] = m
}

func GetModel(repository string, filename string) (*model.Model, error) {
	filename = strings.TrimSuffix(filename, ".vmdl_c")
	filename = strings.TrimSuffix(filename, ".vmdl")
	filename += ".vmdl_c"

	if m := findModel(repository, filename); m != nil {
		if model, ok := m.(*model.Model); ok {
			return model, nil
		} else {
			return nil, errors.New("can't convert to *model.Model")
		}
	}

	file, err := parser.Parse(repository, filename)
	if err != nil {
		return nil, err
	}

	m := model.NewModel()
	m.SetFile(file)

	addModel(repository, filename, m)

	return m, nil
}

func GetSystem(repository string, filename string) (*particles.ParticleSystem, error) {
	filename = strings.TrimSuffix(filename, ".vpcf_c")
	filename = strings.TrimSuffix(filename, ".vpcf")
	filename += ".vpcf_c"

	if s := findModel(repository, filename); s != nil {
		if system, ok := s.(*particles.ParticleSystem); ok {
			return system, nil
		} else {
			return nil, errors.New("can't convert to *model.Model")
		}
	}

	file, err := parser.Parse(repository, filename)
	if err != nil {
		return nil, err
	}

	s := particles.NewParticleSystem()
	s.SetFile(file)

	addModel(repository, filename, s)

	return s, nil
}

func initFlexes(as *sfm.AnimationSet, s2Model *model.Model) error {
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
	}

	//	CreateGlobalFlexControllerOperator
	return nil
}

func initAttachments(as *sfm.AnimationSet, s2Model *model.Model) error {
	attachements, err := s2Model.GetAttachements()
	if err != nil {
		return fmt.Errorf("failed to get model attachements: <%w>", err)
	}

	gameModel := as.GetGameModel()
	for _, v := range attachements {

		gameModel.CreateAttachment(v)
		/*
			defaultValue := v.GetDefaultValue()
			ope := gameModel.CreateGlobalFlexControllerOperator(v.Name, defaultValue)
			c := as.CreateControl(v.Name)
			c.Channel.ToElement = ope
			c.Channel.ToAttribute = "flexWeight"

			layer := any(c.Channel.Log.GetLayer("float log")).(*sfm.LogLayer[float32])
			layer.SetValue(0, defaultValue)

			c.DefaultValue = defaultValue
			c.Value = defaultValue*/
	}

	return nil
}

func AddParticleSystem(clip *sfm.FilmClip, name string, repository string, filename string, parent sfm.INode) (*sfm.AnimationSet, error) {
	filename = strings.TrimSuffix(filename, ".vpcf_c")
	filename = strings.TrimSuffix(filename, ".vpcf")
	filename += ".vpcf"

	as := clip.CreateAnimationSetForParticleSystem(name, filename, parent)
	as.GetParticleSystem().CreateControlPoint(as, 0, nil)

	s2system, err := GetSystem(repository, filename)
	if err != nil {
		return nil, err
	}

	channelsClip := animSetEditorChannels.AddChannelsClip(name)
	channelsClip.AddAnimationSet(as)

	ch := sfm.NewChannel[bool]("simulating channel")
	ch.ToElement = as.GetParticleSystem()
	ch.FromElement = as.GetParticleSystem()
	ch.ToAttribute = "simulating"
	ch.FromAttribute = "simulating"

	layer := any(ch.Log.GetLayer("bool log")).(*sfm.LogLayer[bool])
	layer.SetValue(0, true)
	layer.DefaultValue = true

	channelsClip.AddChannel(ch)

	/*c.DefaultValue = true
	c.Value = true*/

	log.Println(s2system.GetControlPointConfigurationById(0))
	controlPointConfig, err := s2system.GetControlPointConfigurationById(0)
	if err != nil {
		return nil, err
	}

	for _, driver := range controlPointConfig.Drivers {
		cp := as.GetParticleSystem().CreateControlPoint(as, uint(driver.ControlPoint), nil)
		if cp == nil {
			continue
		}

		cp.AttachType = driver.AttachType
		cp.AttachmentName = driver.AttachmentName
		cp.EntityName = driver.EntityName
	}

	return as, nil
}
