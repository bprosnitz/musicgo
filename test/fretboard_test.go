package test

import (
	"testing"

	. "github.com/bprosnitz/musicgo"
	"github.com/bprosnitz/musicgo/guitar"
	"github.com/bprosnitz/musicgo/notes"
)

func TestFretboardStringFret(t *testing.T) {
	type expectedResult struct {
		in    FretboardString
		pitch Pitch
		out   FretboardOffset
		isErr bool
	}
	tests := []expectedResult{
		expectedResult{guitar.StandardTuning.FretboardString(0), notes.E.Octave(4), 0, false},
		expectedResult{guitar.StandardTuning.FretboardString(3), notes.F.Octave(3), 3, false},
		expectedResult{guitar.StandardTuning.FretboardString(3), notes.F.Octave(3) + 0.02, 0, true},
		expectedResult{guitar.StandardTuning.FretboardString(3), notes.F.Octave(-3), 0, true},
	}

	for _, test := range tests {
		offset, err := test.in.Fret(test.pitch)
		if err != nil && !test.isErr {
			t.Errorf("Unexpected error occured: %v", err)
		}
		if err == nil && test.isErr {
			t.Errorf("Expected error to occur, but it didn't.")

		}
		if !test.isErr && offset != test.out {
			t.Errorf("Fretboard offset different from expectation. Got %v. Expected %v.", offset, test.out)
		}
	}
}

func TestFretboardCoordinatePitch(t *testing.T) {
	type expectedResult struct {
		in  FretboardCoordinate
		out Pitch
	}
	tests := []expectedResult{
		expectedResult{guitar.StandardTuning.Position(0, 0), notes.E.Octave(4)},
		expectedResult{guitar.StandardTuning.Position(1, 1), notes.C.Octave(4)},
		expectedResult{guitar.StandardTuning.Position(2, 2), notes.A.Octave(3)},
		expectedResult{guitar.StandardTuning.Position(3, 3), notes.F.Octave(3)},
		expectedResult{guitar.StandardTuning.Position(4, 4), notes.CSharp.Octave(3)},
		expectedResult{guitar.StandardTuning.Position(5, 5), notes.A.Octave(2)},
	}
	for _, test := range tests {
		if test.in.Pitch() != test.out {
			t.Errorf("Pitch for %v was %v. Expected %v.", test.in, test.in.Pitch(), test.out)
		}
	}
}

func TestFretboardCoordinateSwitchFretboard(t *testing.T) {
	type expectedResult struct {
		coord FretboardCoordinate
		newFb *FretboardLayout
		out   FretboardCoordinate
		fail  bool
	}
	tests := []expectedResult{
		expectedResult{guitar.StandardTuning.Position(1, 0), guitar.OpenGTuning, guitar.OpenGTuning.Position(1, 0), false},
		expectedResult{guitar.StandardTuning.Position(5, 0), guitar.OpenGTuning, guitar.OpenGTuning.Position(5, 2), false},
		expectedResult{guitar.OpenGTuning.Position(5, 2), guitar.StandardTuning, guitar.StandardTuning.Position(5, 0), false},
		expectedResult{guitar.OpenGTuning.Position(5, 0), guitar.StandardTuning, guitar.StandardTuning.Position(5, 2), true},
	}
	for _, test := range tests {
		newCoord, err := test.coord.SwitchFretboard(test.newFb)
		if test.fail && err == nil {
			t.Errorf("Expected to get error when converting %v to %v. Got %v.", test.coord, test.newFb, newCoord)
		}
		if !test.fail && err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}
		if !test.fail && newCoord != test.out {
			t.Errorf("Fret coord for %v in %v was %v. Expected %v.", test.coord, test.newFb, newCoord, test.out)
		}
	}
}
