package sfm

import (
	"fmt"
	"strings"
)

type Bone struct {
	Name      string
	NameLower string // For bone merge
	*Node
	transformControl *TransformControl
}

func (b *Bone) isNode() {
}

func (b *Bone) isExportable() bool {
	return true
}

func NewBone(name string, id int) *Bone {
	return &Bone{
		Node:      NewNode(fmt.Sprintf("bone %d (%s)", id, name)),
		Name:      name,
		NameLower: strings.ToLower(name),
	}
}
