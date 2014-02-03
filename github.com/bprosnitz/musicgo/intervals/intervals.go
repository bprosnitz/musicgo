package intervals

import "github.com/bprosnitz/musicgo"

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
