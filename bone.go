package sfm

import "fmt"

type Bone struct {
	Name string
	*Node
}

func (n *Bone) isNode() {
}

func NewBone(name string, id int) *Bone {
	return &Bone{
		Node: NewNode(fmt.Sprintf("bone %d (%s)", id, name)),
		Name: name,
	}
}
