package shawzin

type NoteIndex int

func (n *NoteIndex) String() string {
	return asciiTable()[int(*n)]
}
