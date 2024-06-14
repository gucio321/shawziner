package shawzin

import (
	"fmt"
	"log"

	"github.com/gucio321/shawziner/pkg/mscoreNotationSchema/goMusic"
)

// GetShawzin returns Shawzin implementation of the given piece.
// According to my reverse-engineering experiments:
// the first character of the result is the scale (as a number)
func GetShawzin(piece *goMusic.Notes, scale Scale) (string, error) {
	notes := ""
	pos := 0
	for _, n := range piece.Notes {
		idx, err := GetNote(scale, n.Pitch)
		if err != nil {
			return "", err
		}

		log.Printf("Note: %v=%v (%v), pos %v (%v)", idx, idx.String(), n.Pitch, pos, IntToASCIIPos(pos))
		notes += idx.String() + IntToASCIIPos(pos)

		pos += int(n.Duration)
	}

	return fmt.Sprintf("%d%v", scale, notes), nil
}
