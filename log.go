package sfm

import (
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-vector"
)

type Log[T Loggable] struct {
	Name string
	/*

		CDmaElementArray< CDmeLogLayer >	m_Layers;
		CDmaElement< CDmeCurveInfo > m_CurveInfo;
	*/
}

func newLog[T Loggable]() *Log[T] {
	return &Log[T]{}
}

func (*Log[T]) isLog() {}

func (l *Log[T]) createDmElement(serializer *Serializer) *dmx.DmElement {
	switch any(l).(type) {
	case *Log[bool]:
		return dmx.NewDmElement("bool log", "DmeBoolLog")
	case *Log[float32]:
		return dmx.NewDmElement("float log", "DmeFloatLog")
	case *Log[vector.Vector3[float32]]:
		return dmx.NewDmElement("vector3 log", "DmeVector3Log")
	case *Log[vector.Quaternion[float32]]:
		return dmx.NewDmElement("quaternion log", "DmeQuaternionLog")
	default:
		panic("code this case")
	}
}

func (l *Log[T]) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	//panic("todo")
}
