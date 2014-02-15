package musicgo

import (
	"bytes"
)

type StringIndex int

type FretboardLayout struct {
	strings []FretboardString
}

func NewFretboardLayout(strings ...FretboardString) *FretboardLayout {
	return &FretboardLayout{strings: strings}
}

func (fl *FretboardLayout) String() string {
	var buf bytes.Buffer
	for i, strng := range fl.strings {
		if i > 0 {
			buf.WriteRune(' ')
		}
		buf.WriteString(strng.String())
	}
	return buf.String()
}

func (fl *FretboardLayout) Position(String StringIndex, Fret FretboardOffset) FretboardCoordinate {
	if int(String) < 0 || int(String) >= len(fl.strings) {
		panic("string index out of bounds")
	}
	if Fret < 0 {
		panic("fretboard offset much be non-negative")
	}
	return FretboardCoordinate{fretboardString: String, fret: Fret, fretboardLayout: fl}
}

func (fl *FretboardLayout) NumString() int {
	return len(fl.strings)
}

func (fl *FretboardLayout) FretboardString(index StringIndex) FretboardString {
	return fl.strings[index]
}
