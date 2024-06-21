package sfm

type Camera struct {
	Name string
	Transform
	FieldOfView         float32
	ZNear               float32
	ZFar                float32
	MicDistance         float32
	EyeOffset           float32
	FocalDistance       float32
	Aperture            float32
	ShutterSpeed        float32
	ToneMapScale        float32
	SSAOBias            float32
	SSAOStrength        float32
	SSAORadius          float32
	DepthOfFieldSamples int
	MotionBlurSamples   int
}

func NewCamera(name string) *Camera {
	return &Camera{
		Name:              name,
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
