package musicgo

type StringIndex int

type FretboardLayout struct {
	strings []FretboardString
}

func NewFretboardLayout(strings ...FretboardString) *FretboardLayout {
	return &FretboardLayout{strings: strings}
}

func (fl *FretboardLayout) Position(String StringIndex, Fret FretboardOffset) FretboardCoordinate {
	return FretboardCoordinate{fretboardString: String, fret: Fret, fretboardLayout: fl}
}

func (fl *FretboardLayout) String(index int) FretboardString {
	return fl.strings[index]
}
