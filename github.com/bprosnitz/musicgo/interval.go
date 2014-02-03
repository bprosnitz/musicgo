package musicgo

type Interval float64

func (i Interval) Larger() Interval {
	return i + 1
}

func (i Interval) Smaller() Interval {
	return i - 1
}
