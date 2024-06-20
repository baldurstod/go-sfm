package sfm

type Vec2 [2]float32
type Vec3 [3]float32
type Vec4 [4]float32

func initQuaternion() Vec4 {
	return [...]float32{0, 0, 0, 1}
}
