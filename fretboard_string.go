package musicgo

import (
	"fmt"
	"math"
)

type FretboardString Pitch
type FretboardOffset int

const nilOffset FretboardOffset = 0

func (fs FretboardString) Fret(pitch Pitch) (FretboardOffset, error) {
	offset := FretboardString(pitch) - fs
	pos := FretboardOffset(math.Floor(float64(offset + 0.5)))
	if pos < 0 {
		return pos, fmt.Errorf("Fretboard position negative")
	}
	if math.Abs(float64(offset)-float64(pos)) >= 0.01 {
		return pos, fmt.Errorf("No pitch on this string is within 1 cent error of desired pitch")
	}
	return pos, nil
}
