package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type TimeFrame struct {
	Start     Time
	Duration  Time
	Offset    Time
	Scale     float64
	dmElement *dmx.DmElement
}

func newTimeFrame() *TimeFrame {
	return &TimeFrame{
		Duration: 60,
		Scale:    1,
	}
}

func (tf *TimeFrame) ToDmElement() *dmx.DmElement {
	if tf.dmElement == nil {
		tf.dmElement = dmx.NewDmElement("DmeTimeFrame")
	}
	e := tf.dmElement

	e.CreateFloatAttribute("start", tf.Start)
	e.CreateFloatAttribute("duration", tf.Duration)
	e.CreateFloatAttribute("offset", tf.Offset)
	e.CreateFloatAttribute("scale", tf.Scale)

	return e
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
