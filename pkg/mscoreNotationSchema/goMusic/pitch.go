package goMusic

type Pitch int

// According to my experiment,
// MuseScore stores pitch as an intager number.
//
// Because of our specific usage in this project,
// we'll not put all of the available pitches here.
// Lets start from Low G
//
// Also let me generate names for them:
//
//go:generate stringer -type=Pitch
const (
	G0 Pitch = 53 + iota
	G0Sharp
	A0
	A0Sharp
	B0
	C1
	C1Sharp
	D1
	D1Sharp
	E1
	F1
	F1Sharp
	G1
	G1Sharp
	A1
	A1Sharp
	B1
	C2
	C2Sharp
	D2
	D2Sharp
	E2
	F2
	F2Sharp
	G2
	G2Sharp
	// because of Shawzin's range, I think we can end up here.

	Pause Pitch = 0
)
