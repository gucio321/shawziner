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

const (
	Chromatic Scale = 3
)

func GetNote(scale Scale, pitch goMusic.Pitch) (NoteIndex, error) {
	date := map[Scale]map[goMusic.Pitch]NoteIndex{
		Chromatic: {},
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
