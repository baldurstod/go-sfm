package sfm

import "fmt"

type Bone struct {
	*Node
	boneName string
}

func (n *Bone) isNode() {
}

func NewBone(name string, id int) *Bone {
	return &Bone{
		Node:     NewNode(fmt.Sprintf("bone %d (%s)", id, name)),
		boneName: name,
	}
}
