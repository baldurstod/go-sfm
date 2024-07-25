package sfm

import (
	"github.com/baldurstod/go-vector"
)

type Attachment struct {
	Name        string
	ParentBone  string
	Position    vector.Vector3[float32]
	Orientation vector.Quaternion[float32]
}
