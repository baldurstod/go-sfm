package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Camera struct {
	Name                string
	Transform           *Transform
	Visible             bool
	FieldOfView         float32
	ZNear               float32
	ZFar                float32
	MicDistance         float32
	EyeOffset           float32
	FocalDistance       float32
	Aperture            float32
	ShutterSpeed        Time
	ToneMapScale        float32
	SSAOBias            float32
	SSAOStrength        float32
	SSAORadius          float32
	DepthOfFieldSamples int32
	MotionBlurSamples   int32
}

func NewCamera(name string) *Camera {
	return &Camera{
		Name:                name,
		Transform:           NewTransform(""),
		Visible:             true,
		FieldOfView:         45,
		ZNear:               10,
		ZFar:                20000,
		FocalDistance:       72,
		Aperture:            0.2,
		ShutterSpeed:        0.0208,
		ToneMapScale:        1,
		SSAOBias:            0.5,
		SSAOStrength:        2,
		SSAORadius:          30,
		DepthOfFieldSamples: 64,
		MotionBlurSamples:   8,
	}
}

func (c *Camera) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(c.Name, "DmeCamera")
}

func (c *Camera) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateElementAttribute("transform", serializer.GetElement(c.Transform))
	e.CreateBoolAttribute("visible", c.Visible)
	e.CreateFloatAttribute("fieldOfView", c.FieldOfView)
	e.CreateFloatAttribute("znear", c.ZNear)
	e.CreateFloatAttribute("zfar", c.ZFar)
	e.CreateFloatAttribute("eyeOffset", c.EyeOffset)
	e.CreateFloatAttribute("focalDistance", c.FocalDistance)
	e.CreateFloatAttribute("aperture", c.Aperture)
	e.CreateTimeAttribute("shutterSpeed", c.ShutterSpeed)
	e.CreateFloatAttribute("toneMapScale", c.ToneMapScale)
	e.CreateFloatAttribute("SSAOBias", c.SSAOBias)
	e.CreateFloatAttribute("SSAOStrength", c.SSAOStrength)
	e.CreateFloatAttribute("SSAORadius", c.SSAORadius)
	e.CreateIntAttribute("depthOfFieldSamples", c.DepthOfFieldSamples)
	e.CreateIntAttribute("motionBlurSamples", c.MotionBlurSamples)
}

/*
"camera" "DmeCamera"
{
	"id" "elementid" "29e5cfcb-d276-4dff-8a8c-c0095f8d8884"
	"name" "string" "camera1"
	"transform" "DmeTransform"
	{
		"id" "elementid" "a3d8190a-4cc1-499d-82ec-e54bcb44105a"
		"position" "vector3" "-1657.3859863281 -1764.8508300781 1462.5373535156"
		"orientation" "quaternion" "-0.0299309008 0.0761209726 0.3647043407 0.9275238514"
	}

	"shape" "element" ""
	"visible" "bool" "1"
	"children" "element_array"
	[
	]
	"fieldOfView" "float" "45"
	"znear" "float" "12"
	"zfar" "float" "24000"
	"micDistance" "float" "0"
	"eyeOffset" "float" "0"
	"focalDistance" "float" "72"
	"aperture" "float" "0.200000003"
	"shutterSpeed" "time" "0.0208"
	"toneMapScale" "float" "1"
	"SSAOBias" "float" "0.5"
	"SSAOStrength" "float" "2"
	"SSAORadius" "float" "30"
	"depthOfFieldSamples" "int" "64"
	"motionBlurSamples" "int" "8"
}
*/
