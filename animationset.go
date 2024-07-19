package sfm

import (
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-source2-tools/model"
	"github.com/baldurstod/go-vector"
)

type AnimationSet struct {
	Name                 string
	controls             map[string]IControl
	presetGroups         []*PresetGroup
	operators            []Operator
	RootControlGroup     ControlGroup
	gameModel            *GameModel
	rootTransformControl *TransformControl
}

func newAnimationSet(name string) *AnimationSet {
	as := AnimationSet{
		Name:             name,
		controls:         make(map[string]IControl),
		RootControlGroup: *NewControlGroup(""),
	}

	as.rootTransformControl = as.CreateTransformControl("rootTransform")

	as.rootTransformControl.PositionChannel.Log.GetLayer("vector3 log")
	as.rootTransformControl.OrientationChannel.Log.GetLayer("quaternion log")

	return &as
}

func CreateAnimationSetForModel(name string, filename string) *AnimationSet {
	as := newAnimationSet(name)

	model := newGameModel(name, filename)
	as.setGameModel(model)

	return as
}

func (as *AnimationSet) AddOperator(o Operator) {
	as.operators = append(as.operators, o)
}

func (as *AnimationSet) AddControl(name string, c IControl) {
	as.controls[name] = c
}

func (as *AnimationSet) CreateControl(name string) *Control {
	cg := as.RootControlGroup.GetSubGroup(name)
	c := cg.CreateControl(name)
	as.AddControl(name, c)
	return c
}

func (as *AnimationSet) CreateTransformControl(name string) *TransformControl {
	cg := as.RootControlGroup.GetSubGroup(name)
	c := cg.CreateTransformControl(name)

	c.PositionChannel.ToAttribute = "position"
	c.OrientationChannel.ToAttribute = "orientation"

	as.AddControl(name, c)
	return c
}

func (as *AnimationSet) GetControl(name string) *Control {
	v, ok := as.controls[name]
	if !ok {
		return nil
	}

	c, ok := v.(*Control)
	if !ok {
		return nil
	}

	return c
}

func (as *AnimationSet) GetTransformControl(name string) *TransformControl {
	v, ok := as.controls[name]
	if !ok {
		return nil
	}

	c, ok := v.(*TransformControl)
	if !ok {
		return nil
	}

	return c
}

func (as *AnimationSet) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(as.Name, "DmeAnimationSet")
}

func (as *AnimationSet) isExportable() bool {
	return true
}

func (as *AnimationSet) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	controls := e.CreateAttribute("controls", dmx.AT_ELEMENT_ARRAY)
	for _, c := range as.controls {
		controls.PushElement(serializer.GetElement(c))
	}

	presetGroups := e.CreateAttribute("presetGroups", dmx.AT_ELEMENT_ARRAY)
	for _, pg := range as.presetGroups {
		presetGroups.PushElement(serializer.GetElement(pg))
	}

	operators := e.CreateAttribute("operators", dmx.AT_ELEMENT_ARRAY)
	for _, o := range as.operators {
		operators.PushElement(serializer.GetElement(o))
	}

	e.CreateElementAttribute("rootControlGroup", serializer.GetElement(&as.RootControlGroup))

	if as.gameModel != nil {
		e.CreateElementAttribute("gameModel", serializer.GetElement(as.gameModel))
	} else {
		e.CreateElementAttribute("gameModel", nil)
	}
}

func (as *AnimationSet) setGameModel(model *GameModel) {
	as.gameModel = model

	as.rootTransformControl.PositionChannel.ToElement = model.Transform
	as.rootTransformControl.OrientationChannel.ToElement = model.Transform
}

func (as *AnimationSet) GetGameModel() *GameModel {
	return as.gameModel
}

func (as *AnimationSet) getChannels() []*Channel {
	ret := []*Channel{}
	/*
		ret = append(ret, &as.rootTransformControl.PositionChannel)
		ret = append(ret, &as.rootTransformControl.OrientationChannel)
	*/
	for _, c := range as.controls {
		switch c := c.(type) {
		case *TransformControl:
			ret = append(ret, &c.PositionChannel)
			ret = append(ret, &c.OrientationChannel)
		}
	}

	return ret
}

func (as *AnimationSet) SetFrame(time float32, frame *model.Frame, flexes []model.FlexController) {
	positionChannel := frame.GetChannel("BoneChannel", "Position")
	for _, element := range positionChannel.Datas {

		tc := as.GetTransformControl(element.Name)
		if tc != nil {
			layer := any(tc.PositionChannel.Log.GetLayer("vector3 log")).(*LogLayer[vector.Vector3[float32]])
			layer.SetValue(time, element.Datas.(vector.Vector3[float32]))
		}
	}
	orientationChannel := frame.GetChannel("BoneChannel", "Angle")
	for _, element := range orientationChannel.Datas {

		tc := as.GetTransformControl(element.Name)
		if tc != nil {
			layer := any(tc.OrientationChannel.Log.GetLayer("quaternion log")).(*LogLayer[vector.Quaternion[float32]])
			layer.SetValue(time, (element.Datas.(vector.Quaternion[float32])))
		}
	}

	as.setMorph(time, frame, flexes)
}

func (as *AnimationSet) setMorph(time float32, frame *model.Frame, flexes []model.FlexController) {
	if flexes == nil {
		return
	}

	morphChannel := frame.GetChannel("MorphChannel", "data")
	if morphChannel == nil {
		return
	}

	for _, element := range morphChannel.Datas {

		value := float32(0)
		for _, flex := range flexes {
			if flex.Name == element.Name {
				value = flex.GetControllerValue((element.Datas.(float32)))
				break
			}
		}

		tc := as.GetControl(element.Name)
		if tc != nil {
			layer := any(tc.Channel.Log.GetLayer("float log")).(*LogLayer[float32])
			layer.SetValue(time, value)
		}
	}
}
