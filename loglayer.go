package sfm

import (
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-vector"
)

type Loggable interface {
	bool | int | float32 | vector.Vector3[float32] | vector.Quaternion[float32]
}

type LogLayer[T Loggable] struct {
	times      []int32
	curveTypes []int
	values     []T
	/*
			CDmeLog *m_pOwnerLog;

		mutable int m_lastKey;
		CDmaArray< int > m_times;
		CDmaArray< int > m_CurveTypes;
	*/
}

func newLogLayer[T Loggable]() *LogLayer[T] {
	return &LogLayer[T]{
		times:      make([]int32, 0),
		curveTypes: make([]int, 0),
		values:     make([]T, 0),
	}
}

func (*LogLayer[T]) isLogLayer() {}

func (ll *LogLayer[T]) AddValue(time int32, value T) {
	ll.times = append(ll.times, time)
	ll.values = append(ll.values, value)
}

func (ll *LogLayer[T]) createDmElement(serializer *Serializer) *dmx.DmElement {
	switch any(ll).(type) {
	case *LogLayer[bool]:
		return dmx.NewDmElement("bool log", "DmeBoolLogLayer")
	case *LogLayer[float32]:
		return dmx.NewDmElement("float log", "DmeFloatLogLayer")
	case *LogLayer[vector.Vector3[float32]]:
		return dmx.NewDmElement("vector3 log", "DmeVector3LogLayer")
	case *LogLayer[vector.Quaternion[float32]]:
		return dmx.NewDmElement("quaternion log", "DmeQuaternionLogLayer")
	default:
		panic("code this case")
	}
}

func (ll *LogLayer[T]) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	times := e.CreateAttribute("times", dmx.AT_TIME_ARRAY)
	for _, time := range ll.times {
		times.PushTime(float32(time) / 20.0)
	}

	switch any(ll).(type) {
	/*case *Log[bool]:
		return dmx.NewDmElement("bool log", "DmeBoolLog")
	case *Log[float32]:
		return dmx.NewDmElement("float log", "DmeFloatLog")*/
	case *LogLayer[vector.Vector3[float32]]:
		values := e.CreateAttribute("values", dmx.AT_VECTOR3_ARRAY)
		for _, value := range ll.values {
			values.PushVector3(any(value).(vector.Vector3[float32]))
		}
	/*case *Log[vector.Quaternion[float32]]:
	return dmx.NewDmElement("quaternion log", "DmeQuaternionLog")*/
	default:
		panic("code this case")
	}
}
