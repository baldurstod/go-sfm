package sfm

type Vec2 [2]float64
type Vec3 [3]float64
type Vec4 [4]float64

func initQuaternion() Vec4 {
	return [...]float64{0, 0, 0, 1}
}
