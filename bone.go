package sfm

import "fmt"

type Bone struct {
	Name string
	*Node
}

func (b *Bone) isNode() {
}

func (b *Bone) isExportable() bool {
	return true
}

func NewBone(name string, id int) *Bone {
	return &Bone{
		Node: NewNode(fmt.Sprintf("bone %d (%s)", id, name)),
		Name: name,
	}
}
