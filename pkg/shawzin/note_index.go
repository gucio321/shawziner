package shawzin

type NoteIndex int

func (n *NoteIndex) String() string {
	return string(rune('A' + *n))
}
