package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type AnimationSet struct {
	Name             string
	controls         map[IControl]struct{}
	presetGroups     []*PresetGroup
	operators        []Operator
	RootControlGroup ControlGroup
	GameModel        *GameModel
}

func NewAnimationSet(name string) *AnimationSet {
	return &AnimationSet{
		Name:             name,
		controls:         make(map[IControl]struct{}),
		RootControlGroup: *NewControlGroup(""),
	}
}

func (as *AnimationSet) AddOperator(o Operator) {
	as.operators = append(as.operators, o)
}

func (as *AnimationSet) AddControl(c IControl) {
	as.controls[c] = struct{}{}
}

func (as *AnimationSet) CreateTransformControl(name string) *TransformControl {
	cg := as.RootControlGroup.GetSubGroup(name)
	c := cg.CreateTransformControl(name)
	as.AddControl(c)
	return c
}

func (as *AnimationSet) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(as.Name, "DmeAnimationSet")
}

func (as *AnimationSet) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	controls := e.CreateAttribute("controls", dmx.AT_ELEMENT_ARRAY)
	for c := range as.controls {
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

	if as.GameModel != nil {
		e.CreateElementAttribute("gameModel", serializer.GetElement(as.GameModel))
	} else {
		e.CreateElementAttribute("gameModel", nil)
	}
}
