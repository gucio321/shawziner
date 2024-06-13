package mscoreNotationSchema

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/gucio321/shawziner/pkg/mscoreNotationSchema/goMusic"
)

var InvalidDurationError = errors.New("Invalid duration")

func GetDuration(dur string, dots int) (goMusic.Duration, error) {
	data := map[string]goMusic.Duration{
		"whole":   goMusic.WholeNote,
		"half":    goMusic.HalfNote,
		"quarter": goMusic.QuarterNote,
		"eighth":  goMusic.EighthNote,
		"16th":    goMusic.SixteenthNote,
		"32nd":    goMusic.ThirtySecondNote,
	}

	result, ok := data[dur]
	if !ok {
		return 0, InvalidDurationError
	}

	switch dots {
	case 0:
		// do nothing
	case 1:
		result.Dot()
	default:
		// not implemented
		return 0, fmt.Errorf("more than 1 dot is not implemented yet: %w", InvalidDurationError)
	}

	return result, nil
}

func (m *MUEFile) Unmarshal(data []byte) error {
	return xml.Unmarshal(data, m)
}

func (m *MUEFile) AsGoMusic() (*goMusic.Notes, error) {
	notes := new(goMusic.Notes)
	for _, measure := range m.Measures {
		for _, chord := range measure.Chords {
			duration, err := GetDuration(chord.DurationType, chord.Dots)
			if err != nil {
				return nil, fmt.Errorf("error while parsing duration %s: %w", chord.DurationType, err)
			}

			notes.Notes = append(notes.Notes, goMusic.Note{
				Duration: duration,
			})
		}
	}

	return notes, nil
}
