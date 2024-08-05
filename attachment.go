package sfm

import (
	"strings"

	"github.com/baldurstod/go-vector"
)

type Attachment struct {
	Name            string
	ParentBone      string
	ParentBoneLower string
	Position        vector.Vector3[float32]
	Orientation     vector.Quaternion[float32]
}

func (a *Attachment) SetParentBone(name string) {
	a.ParentBone = name
	a.ParentBoneLower = strings.ToLower(name)

}
