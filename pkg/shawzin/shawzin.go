package shawzin

import (
	"fmt"
	"log"

	"github.com/gucio321/shawziner/pkg/mscoreNotationSchema/goMusic"
)

// GetShawzin returns Shawzin implementation of the given piece.
// According to my reverse-engineering experiments:
// the first character of the result is the scale (as a number)
func GetShawzin(piece *goMusic.Notes, scale Scale) (result string, err error) {
	pos := 0
	for _, n := range piece.Notes {
		log.Printf("Note:%v, pos %v (%v)", n.Pitch, pos, IntToASCIIPos(pos))
		if n.Pitch != goMusic.Pause {
			idx, err := GetNote(scale, n.Pitch)
			if err != nil {
				return "", err
			}

			result += idx.String() + IntToASCIIPos(pos)
		}

		pos += int(n.Duration)
	}

	return fmt.Sprintf("%d%v", scale, result), nil
}
