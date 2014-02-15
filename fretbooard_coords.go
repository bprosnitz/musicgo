package musicgo

import (
	"fmt"
	"math"
)

type FretboardCoordinate struct {
	fretboardString StringIndex
	fret            FretboardOffset
	fretboardLayout *FretboardLayout
}

var zeroFretCoord FretboardCoordinate = FretboardCoordinate{}

func (fc FretboardCoordinate) String() string {
	return fmt.Sprintf("[string: %v (%v), fret: %v, layout: %v]", fc.fretboardLayout.FretboardString(fc.fretboardString), fc.fretboardString, fc.fret, fc.fretboardLayout)
}

func (fc FretboardCoordinate) Pitch() Pitch {
	return Pitch(fc.fretboardLayout.strings[fc.fretboardString]).Interval(Interval(fc.fret))
}

func (fc FretboardCoordinate) SwitchFretboard(newLayout *FretboardLayout) (FretboardCoordinate, error) {
	pitch := fc.Pitch()
	targetString := float64(fc.fretboardString) * float64(newLayout.NumString()) / float64(fc.fretboardLayout.NumString())

	closest := int(math.Floor(targetString + 0.5))

	// go outwards from closest string to increase stability
	for r := 0; r < len(newLayout.strings); r++ {
		indexHigh := closest + r
		if indexHigh >= 0 && indexHigh < newLayout.NumString() {
			offset, err := newLayout.strings[indexHigh].Fret(pitch)
			if err == nil {
				return FretboardCoordinate{fretboardString: StringIndex(indexHigh), fret: FretboardOffset(offset), fretboardLayout: newLayout}, nil
			}
		}

		indexLow := closest - r
		if indexLow >= 0 && indexLow < newLayout.NumString() {
			offset, err := newLayout.strings[indexLow].Fret(pitch)
			if err == nil {
				return FretboardCoordinate{fretboardString: StringIndex(indexLow), fret: FretboardOffset(offset), fretboardLayout: newLayout}, nil
			}
		}
	}

	return zeroFretCoord, fmt.Errorf("No matching note could be found on new fretboard")
}
