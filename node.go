package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Node struct {
	Name           string
	Transform      *Transform
	Visible        bool
	children       []INode
	overrideParent INode
	overridePos    bool
	overrideRot    bool
}

func (n *Node) isNode() {
}

func NewNode(name string) *Node {
	return &Node{
		Name:           name,
		Transform:      NewTransform(name),
		Visible:        true,
		children:       make([]INode, 0),
		overrideParent: nil,
		overridePos:    true,
		overrideRot:    true,
	}
}

func (n *Node) AddChildren(child INode) {
	n.children = append(n.children, child)
}

func (n *Node) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(n.Name, "DmeDag")
}

func (n *Node) isExportable() bool {
	return true
}

func (n *Node) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateElementAttribute("transform", serializer.GetElement(n.Transform))
	e.CreateBoolAttribute("visible", n.Visible)

	children := e.CreateAttribute("children", dmx.AT_ELEMENT_ARRAY)
	for _, child := range n.children {
		children.PushElement(serializer.GetElement(child))
	}

	if n.overrideParent != nil {
		e.CreateElementAttribute("overrideParent", serializer.GetElement(n.overrideParent))
		e.CreateBoolAttribute("overridePos", n.overridePos)
		e.CreateBoolAttribute("overrideRot", n.overrideRot)
	}
}

/*
"DmeDag"
{
	"id" "elementid" "08f3d6d7-31d5-4ddb-97ed-6e4bd2c9f9f6"
	"name" "string" "bone 15 (clavicle_R)"
	"transform" "element" "280dc5cf-9ff9-45a9-b269-140445269a02"
	"shape" "element" ""
	"visible" "bool" "1"
	"children" "element_array"
	[
		"DmeDag"
		{
			"id" "elementid" "41bd1e92-0565-4728-a36e-4c47fa783a2e"
			"name" "string" "bone 16 (arm_upper_R)"
			"transform" "element" "f6830c75-983b-46af-917f-bca4cf58a08e"
			"shape" "element" ""
			"visible" "bool" "1"
			"children" "element_array"
			[
			]
			"overrideParent" "element" "738b7f76-e3f9-4f6c-b6b9-c708484ba9e1"
			"overridePos" "bool" "1"
			"overrideRot" "bool" "1"
		}
	]
	"overrideParent" "element" "55e21315-7fae-43d0-be78-3d670c3b4a4c"
	"overridePos" "bool" "1"
	"overrideRot" "bool" "1"
}
*/
