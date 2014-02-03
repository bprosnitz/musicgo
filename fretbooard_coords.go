package musicgo

type FretboardCoordinate struct {
	fretboardString StringIndex
	fret            FretboardOffset
	fretboardLayout *FretboardLayout
}

func (fc FretboardCoordinate) Pitch() Pitch {
	return Pitch(fc.fretboardLayout.strings[fc.fretboardString]).Interval(Interval(fc.fret))
}
