package sfm

import (
	"slices"

	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-vector"
)

type Loggable interface {
	bool | int | float32 | vector.Vector3[float32] | vector.Quaternion[float32]
}

type LogLayer[T Loggable] struct {
	times        []float32
	curveTypes   []int
	values       map[float32]T
	DefaultValue T
	/*
			CDmeLog *m_pOwnerLog;

		mutable int m_lastKey;
		CDmaArray< int > m_times;
		CDmaArray< int > m_CurveTypes;
	*/
}

func NewLogLayer[T Loggable]() *LogLayer[T] {
	return &LogLayer[T]{
		times:      make([]float32, 0),
		curveTypes: make([]int, 0),
		values:     make(map[float32]T, 0),
	}
}

func (*LogLayer[T]) isLogLayer() {}

func (ll *LogLayer[T]) SetValue(time float32, value T) {
	ll.times = append(ll.times, time)
	ll.values[time] = value
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

func (ll *LogLayer[T]) isExportable() bool {
	return true
}

func (ll *LogLayer[T]) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	times := e.CreateAttribute("times", dmx.AT_TIME_ARRAY)
	/*for time := range ll.values {
		times.PushTime(time)
	}*/

	keys := make([]float32, 0, len(ll.values))
	for k := range ll.values {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	/*for _, k := range keys {
		//fmt.Println(k, romanNumeralDict[k])
	}*/

	switch any(ll).(type) {
	/*case *Log[bool]:
	return dmx.NewDmElement("bool log", "DmeBoolLog")*/
	case *LogLayer[float32]:
		values := e.CreateAttribute("values", dmx.AT_FLOAT_ARRAY)
		for _, time := range keys {
			times.PushTime(time)
			values.PushFloat(any(ll.values[time]).(float32))
		}
		if len(keys) == 0 {
			times.PushTime(0)
			values.PushFloat(any(ll.DefaultValue).(float32))
		}

	case *LogLayer[vector.Vector3[float32]]:
		values := e.CreateAttribute("values", dmx.AT_VECTOR3_ARRAY)
		for _, time := range keys {
			times.PushTime(time)
			values.PushVector3(any(ll.values[time]).(vector.Vector3[float32]))
		}
		if len(keys) == 0 {
			times.PushTime(0)
			values.PushVector3(any(ll.DefaultValue).(vector.Vector3[float32]))
		}
	case *LogLayer[vector.Quaternion[float32]]:
		values := e.CreateAttribute("values", dmx.AT_QUATERNION_ARRAY)
		for _, time := range keys {
			times.PushTime(time)
			values.PushQuaternion(any(ll.values[time]).(vector.Quaternion[float32]))
		}
		if len(keys) == 0 {
			times.PushTime(0)
			values.PushQuaternion(any(ll.DefaultValue).(vector.Quaternion[float32]))
		}
	case *LogLayer[bool]:
		values := e.CreateAttribute("values", dmx.AT_BOOL_ARRAY)
		for _, time := range keys {
			times.PushTime(time)
			values.PushBool(any(ll.values[time]).(bool))
		}
		if len(keys) == 0 {
			times.PushTime(0)
			values.PushBool(any(ll.DefaultValue).(bool))
		}
	default:
		panic("code this case")
	}
}
