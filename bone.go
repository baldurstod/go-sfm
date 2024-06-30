package sfm

type Bone struct {
	*Node
}

func (n *Bone) isNode() {
}

func NewBone(name string) *Bone {
	return &Bone{
		Node: NewNode(name),
	}
}
