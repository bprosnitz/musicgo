package intervals

import "musicgo"

const (
	Unison musicgo.Interval = iota
	MinorSecond
	MajorSecond
	MinorThird
	MajorThird
	PerfectFourth
	AugmentedFourth
	PerfectFifth
	MinorSixth
	MajorSixth
	MinorSeventh
	MajorSeventh
	Octave
)

const (
	DiminishedFifth = AugmentedFourth
	Fifth           = PerfectFifth
)
