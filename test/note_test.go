package test

// test is a separate package to remove an import cycle (for musicgo/notes)

import (
	"testing"

	. "github.com/bprosnitz/musicgo"
	"github.com/bprosnitz/musicgo/intervals"
	"github.com/bprosnitz/musicgo/notes"
)

func TestNoteString(t *testing.T) {
	type expectedResult struct {
		in  Note
		out string
	}
	tests := []expectedResult{
		expectedResult{notes.A, "A"},
		expectedResult{notes.ASharp, "A#"},
		expectedResult{notes.B, "B"},
		expectedResult{notes.C, "C"},
		expectedResult{notes.CSharp, "C#"},
		expectedResult{notes.D, "D"},
		expectedResult{notes.DSharp, "D#"},
		expectedResult{notes.E, "E"},
		expectedResult{notes.F, "F"},
		expectedResult{notes.FSharp, "F#"},
		expectedResult{notes.G, "G"},
		expectedResult{notes.GSharp, "G#"},
		expectedResult{notes.BSharp, "C"},
		expectedResult{notes.ESharp, "F"},
		expectedResult{notes.Ab, "G#"},
		expectedResult{notes.Bb, "A#"},
		expectedResult{notes.Cb, "B"},
		expectedResult{notes.Db, "C#"},
		expectedResult{notes.Eb, "D#"},
		expectedResult{notes.Fb, "E"},
		expectedResult{notes.Gb, "F#"},
		expectedResult{notes.A + 0.52, "A (52 cents)"},
	}
	for _, test := range tests {
		if test.in.String() != test.out {
			t.Errorf("Note string output didn't match expectation. Was %v. Expected %v.", test.in.String(), test.out)
		}
	}
}

func TestParseNote(t *testing.T) {
	type equalData struct {
		first  string
		second string
	}
	type unequalData struct {
		first  string
		second string
	}
	type errorData struct {
		str string
	}
	var testData []interface{} = []interface{}{
		&equalData{"A", "a"},
		&equalData{"A#", "Bb"},
		&equalData{"b", "A##"},
		&equalData{"C", "B#"},
		&equalData{"B##", "Db"},
		&unequalData{"A", "C"},
		&errorData{"x"},
		&errorData{"BW"},
		&errorData{""},
	}
	for _, data := range testData {
		switch data.(type) {
		case *equalData:
			ed := data.(*equalData)
			n1, err := ParseNote(ed.first)
			if err != nil {
				t.Errorf("Error parsing %v: %v", ed.first, err)
			}
			n2, err := ParseNote(ed.second)
			if err != nil {
				t.Errorf("Error parsing %v: %v", ed.second, err)
			}
			if n1 != n2 {
				t.Errorf("Notes are unexpectedly not equal: %v and %v (values %f and %f)", ed.first, ed.second, n1, n2)
			}
		case *unequalData:
			ud := data.(*unequalData)
			n1, err := ParseNote(ud.first)
			if err != nil {
				t.Errorf("Error parsing %v: %v", ud.first, err)
			}
			n2, err := ParseNote(ud.second)
			if err != nil {
				t.Errorf("Error parsing %v: %v", ud.second, err)
			}
			if n1 == n2 {
				t.Errorf("Notes are unexpectedly equal: %v and %v (values %f and %f)", ud.first, ud.second, n1, n2)
			}
		case *errorData:
			ed := data.(*errorData)
			_, err := ParseNote(ed.str)
			if err == nil {
				t.Errorf("Error was expected with %v", ed.str)
			}
		}
	}
}

func TestNoteInterval(t *testing.T) {
	type expectedResult struct {
		initial  Note
		interval Interval
		final    Note
	}
	tests := []expectedResult{
		expectedResult{notes.C, intervals.Unison, notes.C},
		expectedResult{notes.B, intervals.MajorSecond, notes.CSharp},
		expectedResult{notes.D, -intervals.MajorThird, notes.ASharp},
		expectedResult{notes.C, intervals.Octave, notes.C},
		expectedResult{notes.C, intervals.Octave + intervals.MinorSecond, notes.CSharp},
	}
	for _, test := range tests {
		if test.initial.Interval(test.interval) != test.final {
			t.Errorf("Interval failed: initial %v interval %v final %v", test.initial, test.interval, test.final)
		}
	}
}

func TestNoteIndex(t *testing.T) {
	type expectedResult struct {
		input  Note
		output NoteIndex
	}
	tests := []expectedResult{
		expectedResult{notes.C, 0},
		expectedResult{notes.B, 11},
		expectedResult{notes.G, 7},
		expectedResult{notes.C + 0.50, 0},
		expectedResult{notes.C - 0.50, 11},
		expectedResult{notes.D - 0.50, 1},
		expectedResult{notes.C.Interval(intervals.MajorSecond), 2},
		expectedResult{notes.C.Interval(-intervals.MajorSecond), 10},
	}
	for _, test := range tests {
		if test.input.Index() != test.output {
			t.Errorf("Invalid note index for %s: %v. Expected %v.", test.input, test.input.Index(), test.output)
		}
	}
}

func TestNoteCents(t *testing.T) {
	type expectedResult struct {
		input  Note
		output Cents
	}
	tests := []expectedResult{
		expectedResult{notes.C, 0},
		expectedResult{notes.B, 0},
		expectedResult{notes.G, 0},
		expectedResult{100, 0},
		expectedResult{-100, 0},
		expectedResult{-0.49, 51},
		expectedResult{0.49, 49},
		expectedResult{notes.C + 1.19, 19},
	}
	for _, test := range tests {
		if test.input.Cents() != test.output {
			t.Errorf("Invalid cents for %s: %v. Expected %v.", test.input, test.input.Cents(), test.output)
		}
	}
}

func TestNoteOctave(t *testing.T) {
	type expectedResult struct {
		inNote   Note
		inOctave Octave
		outPitch Pitch
	}
	tests := []expectedResult{
		expectedResult{notes.C, 0, 0},
		expectedResult{notes.B, 0, 11},
		expectedResult{notes.B, -1, -1},
		expectedResult{notes.D, 4, 50},
		expectedResult{notes.CSharp + 0.5, 1, 13.5},
	}
	for _, test := range tests {
		if test.inNote.Octave(test.inOctave) != test.outPitch {
			t.Errorf("Expected pitch of %v in octave %v to be %f but was %f", test.inNote, test.inOctave, test.outPitch, test.inNote.Octave(test.inOctave))
		}
	}
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
