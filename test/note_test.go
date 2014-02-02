package test

// test is a separate package to remove an import cycle (for musicgo/notes)

import (
	"testing"

	"musicgo"
	"musicgo/intervals"
	"musicgo/notes"
)

func TestString(t *testing.T) {
	for n := musicgo.Note(0); n < 12; n++ {
		letter := "B"
		for letter_val, num := range noteOffsets {
			if num > n {
				break
			}
			if num == n {
				letter = string(letter_val)
			} else {
				letter = string(letter_val) + "#"
			}
		}
		if musicgo.Note(n).String() != letter {
			t.Errorf("Note value %v and representation %v don't match", n, letter)
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
			n1, err := musicgo.ParseNote(ed.first)
			if err != nil {
				t.Errorf("Error parsing %v: %v", ed.first, err)
			}
			n2, err := musicgo.ParseNote(ed.second)
			if err != nil {
				t.Errorf("Error parsing %v: %v", ed.second, err)
			}
			if n1 != n2 {
				t.Errorf("Notes are unexpectedly not equal: %v and %v (values %f and %f)", ed.first, ed.second, n1, n2)
			}
		case *unequalData:
			ud := data.(*unequalData)
			n1, err := musicgo.ParseNote(ud.first)
			if err != nil {
				t.Errorf("Error parsing %v: %v", ud.first, err)
			}
			n2, err := musicgo.ParseNote(ud.second)
			if err != nil {
				t.Errorf("Error parsing %v: %v", ud.second, err)
			}
			if n1 == n2 {
				t.Errorf("Notes are unexpectedly equal: %v and %v (values %f and %f)", ud.first, ud.second, n1, n2)
			}
		case *errorData:
			ed := data.(*errorData)
			_, err := musicgo.ParseNote(ed.str)
			if err == nil {
				t.Errorf("Error was expected with %v", ed.str)
			}
		}
	}
}

func TestInterval(t *testing.T) {
	type expectedResult struct {
		initial  musicgo.Note
		interval musicgo.Interval
		final    musicgo.Note
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

func TestIndex(t *testing.T) {
	type expectedResult struct {
		input  musicgo.Note
		output musicgo.NoteIndex
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

func TestCents(t *testing.T) {
	type expectedResult struct {
		input  musicgo.Note
		output musicgo.Cents
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

var noteOffsets map[rune]musicgo.Note = map[rune]musicgo.Note{
	'C': 0,
	'D': 2,
	'E': 4,
	'F': 5,
	'G': 7,
	'A': 9,
	'B': 11,
}
