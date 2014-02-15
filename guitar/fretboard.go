package guitar

import (
	. "github.com/bprosnitz/musicgo"
	"github.com/bprosnitz/musicgo/notes"
)

var StandardTuning *FretboardLayout = NewFretboardLayout(
	FretboardString(notes.E.Octave(4)),
	FretboardString(notes.B.Octave(3)),
	FretboardString(notes.G.Octave(3)),
	FretboardString(notes.D.Octave(3)),
	FretboardString(notes.A.Octave(2)),
	FretboardString(notes.E.Octave(2)))

var DropDTuning *FretboardLayout = NewFretboardLayout(
	FretboardString(notes.D.Octave(4)),
	FretboardString(notes.B.Octave(3)),
	FretboardString(notes.G.Octave(3)),
	FretboardString(notes.D.Octave(3)),
	FretboardString(notes.A.Octave(2)),
	FretboardString(notes.E.Octave(2)))

var OpenGTuning *FretboardLayout = NewFretboardLayout(
	FretboardString(notes.D.Octave(4)),
	FretboardString(notes.B.Octave(3)),
	FretboardString(notes.G.Octave(3)),
	FretboardString(notes.D.Octave(3)),
	FretboardString(notes.G.Octave(2)),
	FretboardString(notes.D.Octave(2)))
