package musicgo

import (
	"fmt"
	"io"
	"math"
	"strings"
)

type Cents int
type NoteIndex int
type Note float64

const nilNote Note = 0

func normalize(n Note) Note {
	return n - Note(math.Floor(float64(n/12))*12)
}

func (n Note) Interval(i Interval) Note {
	return normalize(n + Note(i))
}

func (n Note) Index() NoteIndex {
	m := NoteIndex(math.Floor(float64(n))) % 12
	if m < 0 {
		return NoteIndex(12 + m)
	}
	return m
}

func (n Note) Cents() Cents {
	n64 := float64(n)
	return Cents(math.Floor(0.5 + (n64-math.Floor(n64))*100))
}

func (n Note) String() string {
	name := noteNames[Note(n.Index())]
	c := n.Cents()
	if c != 0 {
		return fmt.Sprintf("%v (%v cents)", name, c)
	}
	return name
}

func ParseNote(s string) (Note, error) {
	return parseNote(strings.NewReader(s))
}

func parseNote(r *strings.Reader) (Note, error) {
	// read the letter
	c, _, err := r.ReadRune()
	if err != nil {
		return nilNote, err
	}
	if c >= 'a' && c <= 'z' {
		c += 'A' - 'a'
	}
	n, ok := noteOffsets[c]
	if !ok {
		return nilNote, fmt.Errorf("Invalid note letter: %c", c)
	}

	// read sharps or flats
	for {
		c, _, err = r.ReadRune()
		if err != nil && err != io.EOF {
			return nilNote, err
		}

		if err == io.EOF {
			break
		}

		switch c {
		case '#':
			n = n + 1
		case 'b':
			n = n - 1
		default:
			return nilNote, fmt.Errorf("Invalid rune: %v", c)
		}
	}

	return normalize(n), nil
}

var noteOffsets map[rune]Note = map[rune]Note{
	'C': 0,
	'D': 2,
	'E': 4,
	'F': 5,
	'G': 7,
	'A': 9,
	'B': 11,
}

var noteNames map[Note]string = map[Note]string{
	0:  "C",
	1:  "C#",
	2:  "D",
	3:  "D#",
	4:  "E",
	5:  "F",
	6:  "F#",
	7:  "G",
	8:  "G#",
	9:  "A",
	10: "A#",
	11: "B",
}
