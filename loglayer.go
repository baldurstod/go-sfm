package sfm

import (
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-vector"
)

type Loggable interface {
	bool | int | float32 | vector.Vector3[float32] | vector.Quaternion[float32]
}

type LogLayer[T Loggable] struct {
	owner      *Log[T]
	times      []int
	curveTypes []int
	/*
			CDmeLog *m_pOwnerLog;

		mutable int m_lastKey;
		CDmaArray< int > m_times;
		CDmaArray< int > m_CurveTypes;
	*/
}

func newLogLayer(owner *Log[int]) *LogLayer[int] {
	return &LogLayer[int]{
		owner:      owner,
		times:      make([]int, 0),
		curveTypes: make([]int, 0),
	}
}

func (s *LogLayer[int]) toDmElement(serializer *Serializer, e *dmx.DmElement) {
}
