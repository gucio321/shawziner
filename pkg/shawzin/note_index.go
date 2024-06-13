package shawzin

type NoteIndex int

func (n *NoteIndex) String() string {
	indexMap := map[NoteIndex]rune{}

	result, ok := indexMap[*n]
	if !ok {
		return "Unknown"
	}

	return string(result)
}
