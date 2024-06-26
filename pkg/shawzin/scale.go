package shawzin

import (
	"errors"
	"fmt"

	"github.com/gucio321/shawziner/pkg/mscoreNotationSchema/goMusic"
)

var (
	ErrScaleNotFound = errors.New("scale not found")
	ErrInvalidNote   = errors.New("invalid note")
)

type Scale int

//go:generate stringer -type=Scale
const (
	Chromatic Scale = 3
)

func GetNote(scale Scale, pitch goMusic.Pitch) (NoteIndex, error) {
	date := map[Scale]map[goMusic.Pitch]NoteIndex{
		Chromatic: {
			goMusic.C1:      1,
			goMusic.C1Sharp: 2,
			goMusic.D1:      4,
			goMusic.D1Sharp: 9,
			goMusic.E1:      10,
			goMusic.F1:      12,
			goMusic.F1Sharp: 17,
			goMusic.G1:      18,
			goMusic.G1Sharp: 20,
			goMusic.A1:      33,
			goMusic.A1Sharp: 34,
			goMusic.B1:      36,
			goMusic.C2:      52,
			goMusic.C2Sharp: 44,
			goMusic.D2:      28,
		},
	}

	sc, Ok := date[scale]
	if !Ok {
		return 0, ErrScaleNotFound
	}

	note, Ok := sc[pitch]
	if !Ok {
		return 0, fmt.Errorf("Invalid scale %s for note %v: %w", scale, pitch, ErrInvalidNote)
	}

	return note, nil
}
