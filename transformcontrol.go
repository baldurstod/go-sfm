package sfm

import (
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-vector"
)

type TransformControl struct {
	Name               string
	ValuePosition      vector.Vector3[float32]
	ValueOrientation   vector.Quaternion[float32]
	PositionChannel    Channel
	OrientationChannel Channel
}

func (*TransformControl) isControl() {}

func NewTransformControl(name string) *TransformControl {
	return &TransformControl{
		Name: name,
	}
}

func (tc *TransformControl) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(tc.Name, "DmeTransformControl")

	e.CreateVector3Attribute("value", tc.ValuePosition)
	e.CreateQuaternionAttribute("defaultValue", tc.ValueOrientation)
	e.CreateElementAttribute("positionChannel", serializer.GetElement(&tc.PositionChannel))
	e.CreateElementAttribute("orientationChannel", serializer.GetElement(&tc.OrientationChannel))

	return e
}

/*
"DmeTransformControl"
{
	"id" "elementid" "e9f7377f-1f74-4596-a876-6b7002d1ab24"
	"name" "string" "rootTransform"
	"valuePosition" "vector3" "-1422.6496582031 -1422.6496582031 1306.7924804688"
	"valueOrientation" "quaternion" "0 0 0.9238795042 -0.3826834559"
	"positionChannel" "element" "1ce0e6d2-4e40-4240-b9f9-e4db532702d7"
	"orientationChannel" "element" "12977884-3ccb-438d-b311-0c0f57336216"
}
*/
