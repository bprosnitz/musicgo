package musicgo

import (
	"fmt"
	"math"
)

type Octave int
type Pitch float64

const nilPitch Pitch = 0

func (p Pitch) Interval(i Interval) Pitch {
	return p + Pitch(i)
}

func (p Pitch) Index() NoteIndex {
	m := NoteIndex(math.Floor(float64(p))) % 12
	if m < 0 {
		return 12 + m
	}
	return m
}

func (p Pitch) Cents() Cents {
	n64 := float64(p)
	return Cents(math.Floor(0.5 + (n64-math.Floor(n64))*100))
}

func (p Pitch) Octave() Octave {
	return Octave(math.Floor(float64(p / 12)))
}

func (p Pitch) String() string {
	noteName := noteNames[Note(p.Index())]
	o := p.Octave()
	c := p.Cents()
	var octaveName string
	if o < 0 {
		octaveName = fmt.Sprintf("[%v]", o)
	} else {
		octaveName = fmt.Sprintf("%v", o)
	}
	if c != 0 {
		return fmt.Sprintf("%v%v (%v cents)", noteName, octaveName, c)
	}
	return noteName + octaveName
}

func (p Pitch) Note() Note {
	return normalize(Note(p))
}
