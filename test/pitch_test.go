package test

// test is a separate package to remove an import cycle (for musicgo/notes)

import (
	"testing"

	. "github.com/bprosnitz/musicgo"
	"github.com/bprosnitz/musicgo/intervals"
	"github.com/bprosnitz/musicgo/notes"
)

func TestPitchString(t *testing.T) {
	type expectedResult struct {
		in  Pitch
		out string
	}
	tests := []expectedResult{
		expectedResult{notes.A.Octave(0), "A0"},
		expectedResult{notes.Db.Octave(3), "C#3"},
		expectedResult{notes.Db.Octave(-1), "C#[-1]"},
		expectedResult{notes.B.Octave(2) + 0.2, "B2 (20 cents)"},
		expectedResult{notes.F.Octave(-1) + 0.342, "F[-1] (34 cents)"},
	}
	for _, test := range tests {
		if test.in.String() != test.out {
			t.Errorf("Pitch string output didn't match expectation. Was %v. Expected %v.", test.in.String(), test.out)
		}
	}
}

func TestPitchInterval(t *testing.T) {
	type expectedResult struct {
		initial  Pitch
		interval Interval
		final    Pitch
	}
	tests := []expectedResult{
		expectedResult{notes.C.Octave(2), intervals.Octave, notes.C.Octave(3)},
		expectedResult{notes.B.Octave(1), -2 * intervals.Octave, notes.B.Octave(-1)},
		expectedResult{notes.D.Octave(2) + 0.2, 13.7, notes.DSharp.Octave(3) + 0.9},
	}
	for _, test := range tests {
		if test.initial.Interval(test.interval) != test.final {
			t.Errorf("Interval failed: initial %v interval %v final %v", test.initial, test.interval, test.final)
		}
	}
}

func TestPitchIndex(t *testing.T) {
	type expectedResult struct {
		input  Pitch
		output NoteIndex
	}
	tests := []expectedResult{
		expectedResult{notes.C.Octave(5), 0},
		expectedResult{notes.B.Octave(-1), 11},
		expectedResult{notes.G.Octave(2), 7},
		expectedResult{notes.C.Octave(5) + 0.3, 0},
		expectedResult{notes.C.Octave(5) - 0.3, 11},
	}
	for _, test := range tests {
		if test.input.Index() != test.output {
			t.Errorf("Invalid pitch index for %s: %v. Expected %v.", test.input, test.input.Index(), test.output)
		}
	}
}

func TestPitchCents(t *testing.T) {
	type expectedResult struct {
		input  Pitch
		output Cents
	}
	tests := []expectedResult{
		expectedResult{notes.C.Octave(4), 0},
		expectedResult{notes.B.Octave(-1) + 0.2, 20},
		expectedResult{notes.A.Octave(5) - 0.3, 70},
	}
	for _, test := range tests {
		if test.input.Cents() != test.output {
			t.Errorf("Invalid cents for %s: %v. Expected %v.", test.input, test.input.Cents(), test.output)
		}
	}
}

func TestPitchOctave(t *testing.T) {
	type expectedResult struct {
		input  Pitch
		output Octave
	}
	tests := []expectedResult{
		expectedResult{notes.C.Octave(0), 0},
		expectedResult{notes.B.Octave(0), 0},
		expectedResult{notes.B.Octave(1) + 0.8, 1},
		expectedResult{notes.C.Octave(0) - 0.1, -1},
		expectedResult{notes.A.Octave(-1), -1},
		expectedResult{notes.F.Octave(9) + 0.5, 9},
		expectedResult{notes.G.Octave(4), 4},
	}
	for _, test := range tests {
		if test.input.Octave() != test.output {
			t.Errorf("Invalid octave for %s: %v. Expected %v.", test.input, test.input.Octave(), test.output)
		}
	}
}
