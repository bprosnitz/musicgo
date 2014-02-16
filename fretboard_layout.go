package musicgo

import (
	"bytes"
	"sync"
)

type StringIndex int

type FretboardLayout struct {
	strings []FretboardString
}

type fretboardLayoutHashConsCacheType struct {
	cache []*FretboardLayout
	mutex sync.Mutex
}

func (flc *fretboardLayoutHashConsCacheType) HashCons(fl *FretboardLayout) *FretboardLayout {
	flc.mutex.Lock()
	for _, existingLayout := range flc.cache {
		existingStrings := existingLayout.strings
		if len(existingStrings) != len(fl.strings) {
			continue
		}
		foundDifference := false
		for i, existingString := range existingLayout.strings {
			if fl.strings[i] != existingString {
				foundDifference = true
				break
			}
		}

		if !foundDifference {
			flc.mutex.Unlock()
			return existingLayout
		}
	}
	flc.cache = append(flc.cache, fl)

	flc.mutex.Unlock()
	return fl
}

var fretboardLayoutCache fretboardLayoutHashConsCacheType = fretboardLayoutHashConsCacheType{}

func NewFretboardLayout(strings ...FretboardString) *FretboardLayout {
	return fretboardLayoutCache.HashCons(&FretboardLayout{strings: strings})
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
