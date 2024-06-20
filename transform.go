package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Transform struct {
	Name        string
	Position    [3]float64
	Orientation [4]float64
	Scale       float64
}

func NewTransform(name string) *Transform {
	return &Transform{
		Name:        name,
		Orientation: [...]float64{0, 0, 0, 1},
		Scale:       1,
	}
}

func (t *Transform) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmeAnimationSet")

	e.CreateStringAttribute("name", t.Name)
	e.CreateVector3Attribute("position", t.Position)
	e.CreateQuaternionAttribute("orientation", t.Orientation)
	e.CreateFloatAttribute("scale", t.Scale)

	return e
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
