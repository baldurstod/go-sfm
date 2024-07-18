package sfm

import (
	"github.com/baldurstod/go-dmx"
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

func NewAnimationSet(name string) *AnimationSet {
	as := AnimationSet{
		Name:             name,
		controls:         make(map[string]IControl),
		RootControlGroup: *NewControlGroup(""),
	}

	as.rootTransformControl = as.CreateTransformControl("rootTransform")
	as.rootTransformControl.PositionChannel.ToAttribute = "position"
	as.rootTransformControl.OrientationChannel.ToAttribute = "orientation"

	posLayer := any(as.rootTransformControl.PositionChannel.Log.GetLayer("vector3 log")).(*LogLayer[vector.Vector3[float32]])
	posLayer.SetValue(0, vector.Vector3[float32]{})
	rotLayer := any(as.rootTransformControl.OrientationChannel.Log.GetLayer("quaternion log")).(*LogLayer[vector.Quaternion[float32]])
	rotLayer.SetValue(0, vector.Quaternion[float32]{})

	return &as
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

func (as *AnimationSet) SetGameModel(model *GameModel) {
	as.gameModel = model

	as.rootTransformControl.PositionChannel.ToElement = model.Transform
	as.rootTransformControl.OrientationChannel.ToElement = model.Transform
}

func (as *AnimationSet) GetGameModel() *GameModel {
	return as.gameModel
}

func (as *AnimationSet) AddToChannelsClip(clip *ChannelsClip) {
	clip.AddChannel(&as.rootTransformControl.OrientationChannel)
	clip.AddChannel(&as.rootTransformControl.PositionChannel)
}
