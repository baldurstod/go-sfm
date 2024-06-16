package sfm

type Camera struct {
	Name string
	Transform
	FieldOfView         float
	ZNear               float
	ZFar                float
	MicDistance         float
	EyeOffset           float
	FocalDistance       float
	Aperture            float
	ShutterSpeed        float
	ToneMapScale        float
	SSAOBias            float
	SSAOStrength        float
	SSAORadius          float
	DepthOfFieldSamples int
	MotionBlurSamples   int
}

func newCamera() *Camera {
	return &Camera{
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
