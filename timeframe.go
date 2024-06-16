package sfm

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
