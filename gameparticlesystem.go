package sfm

import (
	"fmt"

	"github.com/baldurstod/go-dmx"
)

type GameParticleSystem struct {
	*Node
	systemName    string
	controlPoints []*ControlPoint
	parentModel   *GameModel
}

func newGameParticleSystem(name string, systemName string) *GameParticleSystem {
	return &GameParticleSystem{
		Node:          NewNode(name),
		systemName:    systemName,
		controlPoints: make([]*ControlPoint, 0, 1),
	}
}

func (gps *GameParticleSystem) isNode() {
}

func (gps *GameParticleSystem) CreateControlPoint(as *AnimationSet, id int) *ControlPoint {
	cp := newControlPoint(id)
	gps.controlPoints = append(gps.controlPoints, cp)

	name := fmt.Sprintf("controlPoint%d", id)
	tc := as.CreateTransformControl(name)
	tc.PositionChannel.ToElement = cp.Transform
	tc.OrientationChannel.ToElement = cp.Transform

	cp.transformControl = tc

	tc.PositionChannel.Log.GetLayer("vector3 log")
	tc.OrientationChannel.Log.GetLayer("quaternion log")

	gps.children = append(gps.children, cp)
	//any(tc.PositionChannel.Log.GetLayer("vector3 log")).(*LogLayer[vector.Vector3[float32]]).defaultValue = cp.Transform.Position
	//any(tc.OrientationChannel.Log.GetLayer("quaternion log")).(*LogLayer[vector.Quaternion[float32]]).defaultValue = cp.Transform.Orientation

	return cp
}

func (gps *GameParticleSystem) getControlPoint(id int) *ControlPoint {
	return gps.controlPoints[id]
}

func (gps *GameParticleSystem) SetParentModel(parent *GameModel) {
	gps.parentModel = parent
	for _, controlPoint := range gps.controlPoints {
		if parent == nil {
			controlPoint.overrideParent = nil
		} else {
			parentBone := parent.getBoneByName(controlPoint.Name)
			if parentBone != nil {
				controlPoint.overrideParent = parentBone
				controlPoint.transformControl.exportable = false
				controlPoint.Transform.isIdentity = true
			} else {
				controlPoint.overrideParent = nil
				controlPoint.transformControl.exportable = true
				controlPoint.Transform.isIdentity = false
			}
		}
	}
}

func (gps *GameParticleSystem) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(gps.Name, "DmeGameParticleSystem")
}

func (gps *GameParticleSystem) isExportable() bool {
	return true
}

func (gps *GameParticleSystem) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateStringAttribute("particleSystemType", gps.systemName)

	e.CreateElementAttribute("transform", serializer.GetElement(gps.Transform))
	e.CreateBoolAttribute("visible", gps.Visible)
	e.CreateBoolAttribute("simulating", true)
	e.CreateBoolAttribute("emitting", true)
	e.CreateBoolAttribute("inEndCap", false)
	e.CreateBoolAttribute("visible", gps.Visible)
	e.CreateIntAttribute("randomSeed", 0)
	e.CreateFloatAttribute("simulationTimeScale", 1)
	e.CreateFloatAttribute("depthSortBias", 0)

	children := e.CreateAttribute("children", dmx.AT_ELEMENT_ARRAY)
	for _, child := range gps.children {
		children.PushElement(serializer.GetElement(child))
	}

	bones := e.CreateAttribute("controlPoints", dmx.AT_ELEMENT_ARRAY)
	for _, cp := range gps.controlPoints {
		bones.PushElement(serializer.GetElement(cp.Transform))
	}

	if gps.parentModel != nil {
		e.CreateElementAttribute("overrideParent", serializer.GetElement(gps.parentModel))
	}
}
