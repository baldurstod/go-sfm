package sfm

import (
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-source2-tools/model"
	"github.com/baldurstod/go-vector"
)

type GameModel struct {
	*Node
	ModelName               string
	Skin                    int32
	flexWeights             []float32
	MeshGroupMask           uint64
	bones                   []*Bone
	attachments             map[string]*Attachment
	globalFlexControllers   []*GlobalFlexControllerOperator
	ComputeBounds           bool
	EvaluateProceduralBones bool
	parentModel             *GameModel
}

func newGameModel(name string, modelName string) *GameModel {
	return &GameModel{
		Node:                    NewNode(name),
		ModelName:               modelName,
		Skin:                    0,
		flexWeights:             make([]float32, 0),
		bones:                   make([]*Bone, 0),
		attachments:             make(map[string]*Attachment, 0),
		MeshGroupMask:           0xffffffffffffffff,
		globalFlexControllers:   make([]*GlobalFlexControllerOperator, 0),
		ComputeBounds:           true,
		EvaluateProceduralBones: true,
	}
}

func (gm *GameModel) isNode() {
}

func (gm *GameModel) AddChildren(child INode) {
	gm.children = append(gm.children, child)
}

func (gm *GameModel) CreateBone(as *AnimationSet, name string, id int, position vector.Vector3[float32], orientation vector.Quaternion[float32]) *Bone {
	bone := NewBone(name, id)
	gm.bones = append(gm.bones, bone)
	bone.Transform.Position = position
	bone.Transform.Orientation = orientation

	tc := as.CreateTransformControl(name)
	tc.PositionChannel.ToElement = bone.Transform
	//tc.PositionChannel.ToAttribute = "position"
	tc.OrientationChannel.ToElement = bone.Transform
	//tc.OrientationChannel.ToAttribute = "orientation"

	tc.ValuePosition = position
	tc.ValueOrientation = orientation

	bone.transformControl = tc

	any(tc.PositionChannel.Log.GetLayer("vector3 log")).(*LogLayer[vector.Vector3[float32]]).DefaultValue = bone.Transform.Position
	any(tc.OrientationChannel.Log.GetLayer("quaternion log")).(*LogLayer[vector.Quaternion[float32]]).DefaultValue = bone.Transform.Orientation

	return bone
}

func (gm *GameModel) CreateAttachment(src *model.Attachment) *Attachment {
	attachment := Attachment{}

	attachment.Name = src.Name
	if len(src.Influences) > 0 {
		// we can only use one influence ?
		influence := src.Influences[0]

		attachment.SetParentBone(influence.Name)
		attachment.Position = influence.Offset
		attachment.Orientation = influence.Rotation
	}

	gm.attachments[attachment.Name] = &attachment
	return &attachment
}

func (gm *GameModel) CreateGlobalFlexControllerOperator(name string, flexWeight float32) *GlobalFlexControllerOperator {
	o := NewGlobalFlexControllerOperator(name, flexWeight, gm)
	gm.globalFlexControllers = append(gm.globalFlexControllers, o)
	gm.flexWeights = append(gm.flexWeights, flexWeight)
	return o
}

func (gm *GameModel) getBoneByName(name string) *Bone {
	for _, bone := range gm.bones {
		if bone.NameLower == name {
			return bone
		}
	}
	return nil
}

func (gm *GameModel) getAttachmentByName(name string) *Attachment {
	return gm.attachments[name]
}

func (gm *GameModel) SetParentModel(parent *GameModel) {
	gm.parentModel = parent
	for _, bone := range gm.bones {
		if parent == nil {
			bone.overrideParent = nil
		} else {
			parentBone := parent.getBoneByName(bone.NameLower)
			if parentBone != nil {
				bone.overrideParent = parentBone
				bone.transformControl.exportable = false
				bone.Transform.isIdentity = true
			} else {
				bone.overrideParent = nil
				bone.transformControl.exportable = true
				bone.Transform.isIdentity = false
			}
		}
	}
}

func (gm *GameModel) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(gm.Name, "DmeGameModel")
}

func (gm *GameModel) isExportable() bool {
	return true
}

func (gm *GameModel) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateStringAttribute("modelName", gm.ModelName)
	e.CreateIntAttribute("skin", gm.Skin)
	e.CreateElementAttribute("transform", serializer.GetElement(gm.Transform))
	e.CreateBoolAttribute("visible", gm.Visible)
	e.CreateUint64Attribute("meshGroupMask", gm.MeshGroupMask)
	e.CreateBoolAttribute("computeBounds", gm.ComputeBounds)
	e.CreateBoolAttribute("evaluateProceduralBones", gm.EvaluateProceduralBones)

	children := e.CreateAttribute("children", dmx.AT_ELEMENT_ARRAY)
	for _, child := range gm.children {
		children.PushElement(serializer.GetElement(child))
	}

	flexWeights := e.CreateAttribute("flexWeights", dmx.AT_FLOAT_ARRAY)
	for _, weights := range gm.flexWeights {
		flexWeights.PushFloat(weights)
	}

	bones := e.CreateAttribute("bones", dmx.AT_ELEMENT_ARRAY)
	for _, bone := range gm.bones {
		bones.PushElement(serializer.GetElement(bone.Transform))
	}

	globalFlexControllers := e.CreateAttribute("globalFlexControllers", dmx.AT_ELEMENT_ARRAY)
	for _, gfc := range gm.globalFlexControllers {
		globalFlexControllers.PushElement(serializer.GetElement(gfc))
	}

	if gm.parentModel != nil {
		e.CreateElementAttribute("overrideParent", serializer.GetElement(gm.parentModel))
	}
}
