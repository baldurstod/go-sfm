package sfm

import (
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-vector"
)

type Transform struct {
	Name        string
	Position    vector.Vector3[float32]
	Orientation vector.Quaternion[float32]
	Scale       float32
	isIdentity  bool
}

func NewTransform(name string) *Transform {
	return &Transform{
		Name:        name,
		Orientation: [...]float32{0, 0, 0, 1},
		Scale:       1,
	}
}

func (t *Transform) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(t.Name, "DmeTransform")
}

func (t *Transform) isExportable() bool {
	return true
}

func (t *Transform) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	if t.isIdentity {
		e.CreateVector3Attribute("position", [...]float32{0, 0, 0})
		e.CreateQuaternionAttribute("orientation", [...]float32{0, 0, 0, 1})
		e.CreateFloatAttribute("scale", 1)
	} else {
		e.CreateVector3Attribute("position", t.Position)
		e.CreateQuaternionAttribute("orientation", t.Orientation)
		e.CreateFloatAttribute("scale", t.Scale)
	}
}

/*
"DmeTransform"
{
	"id" "elementid" "9168f52d-a1ab-47f5-8e90-4b09e5c88eba"
	"position" "vector3" "-1422.6496582031 -1422.6496582031 1306.7924804688"
	"orientation" "quaternion" "0 0 0.9238795042 -0.3826834559"
	"scale" "float" "0.9300000072"
}
*/
