package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Bone struct {
	Name        string
	Position    [3]float32
	Orientation [4]float32
}

func NewBone(name string) *Bone {
	return &Bone{
		Name:        name,
		Orientation: [...]float32{0, 0, 0, 1},
	}
}

func (b *Bone) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(b.Name, "DmeAnimationSet")

	e.CreateVector3Attribute("position", b.Position)
	e.CreateQuaternionAttribute("orientation", b.Orientation)

	return e
}

/*
"DmeTransform"
{
	"id" "elementid" "d94d28cb-65eb-4f35-a986-c8625d80f70d"
	"name" "string" "bone 6 (elbow_R)"
	"position" "vector3" "-64.4283981323 0.0000076294 0"
	"orientation" "quaternion" "0 -0.7285940051 0 -0.6849458218"
}
*/
