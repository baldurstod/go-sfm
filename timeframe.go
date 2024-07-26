package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type TimeFrame struct {
	Start    Time
	Duration Time
	Offset   Time
	Scale    float32
}

func newTimeFrame() *TimeFrame {
	return &TimeFrame{
		Duration: 60,
		Scale:    1,
	}
}

func (tf *TimeFrame) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement("", "DmeTimeFrame")
}

func (tf *TimeFrame) isExportable() bool {
	return true
}

func (tf *TimeFrame) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateTimeAttribute("start", tf.Start)
	e.CreateTimeAttribute("duration", tf.Duration)
	e.CreateTimeAttribute("offset", tf.Offset)
	e.CreateFloatAttribute("scale", tf.Scale)
}

func (tf *TimeFrame) GetDuration() float32 {
	return tf.Duration
}

/*
	"timeFrame" "DmeTimeFrame"
	{
		"id" "elementid" "70fca316-57d0-4748-9ba2-106e8a6ace88"
		"start" "time" "0.0000"
		"duration" "time" "60.0000"
		"offset" "time" "0.0000"
		"scale" "float" "1"
	}
*/
