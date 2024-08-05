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
		controlPoints: make([]*ControlPoint, 0, 10),
	}
}

func (gps *GameParticleSystem) isNode() {
}

func (gps *GameParticleSystem) CreateControlPoint(as *AnimationSet, id uint, controlModel *GameModel) *ControlPoint {

	// Create missing points
	l := uint(len(gps.controlPoints))
	if id >= l {
		for i := l; i <= id; i++ {
			cp := newControlPoint(id)
			gps.children = append(gps.children, cp)
			gps.controlPoints = append(gps.controlPoints, cp)

			name := fmt.Sprintf("controlPoint%d", i)
			tc := as.CreateTransformControl(name)
			tc.PositionChannel.ToElement = cp.Transform
			tc.OrientationChannel.ToElement = cp.Transform

			cp.transformControl = tc

			tc.PositionChannel.Log.GetLayer("vector3 log")
			tc.OrientationChannel.Log.GetLayer("quaternion log")
		}
	}

	cp := gps.controlPoints[id]
	//if id >= uint(len(gps.controlPoints)) {
	/*if cp == nil {
		cp = newControlPoint(id)
		//gps.controlPoints = append(gps.controlPoints, cp)
		gps.controlPoints[id] = cp
		gps.children = append(gps.children, cp)
	}*/

	cp.ControlModel = controlModel

	return cp
}

func (gps *GameParticleSystem) getControlPoint(id int) *ControlPoint {
	return gps.controlPoints[id]
}

func (gps *GameParticleSystem) SetParentModel(parent *GameModel) {
	if parent == nil {
		return
	}
	gps.parentModel = parent
	var attachement *Attachment
	for _, controlPoint := range gps.controlPoints {
		controlPoint.ControlModel = parent
		if parent == nil {
			controlPoint.overrideParent = nil
		} else {
			if controlPoint.EntityName == "parent" {
				if parent.parentModel != nil {
					parent = parent.parentModel
				}
			}

			if parent != nil {
				attachement = parent.getAttachmentByName(controlPoint.AttachmentName)
			}

			if attachement != nil {
				parentBone := parent.getBoneByName(attachement.ParentBoneLower)
				if parentBone != nil {
					controlPoint.overrideParent = parentBone
					controlPoint.transformControl.exportable = false
					//controlPoint.Transform.isIdentity = true
					controlPoint.Transform.Position = attachement.Position
					controlPoint.Transform.Orientation = attachement.Orientation
					continue
				}
			}
			controlPoint.overrideParent = nil
			controlPoint.transformControl.exportable = true
			controlPoint.Transform.isIdentity = false
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

	controlPoints := e.CreateAttribute("controlPoints", dmx.AT_ELEMENT_ARRAY)
	controlModels := e.CreateAttribute("controlModels", dmx.AT_ELEMENT_ARRAY)
	for _, cp := range gps.controlPoints {
		controlPoints.PushElement(serializer.GetElement(cp.Transform))
		if cp.ControlModel != nil {
			controlModels.PushElement(serializer.GetElement(cp.ControlModel))
		}
	}

	if gps.parentModel != nil {
		e.CreateElementAttribute("overrideParent", serializer.GetElement(gps.parentModel))
	}
}
