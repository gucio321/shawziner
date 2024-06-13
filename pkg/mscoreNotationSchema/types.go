package mscoreNotationSchema

type MUEFile struct {
	Measures []Measure `xml:"Score>Staff>Measure"`
}

type Measure struct {
	Chords []Chord `xml:"voice>Chord"`
}

type Chord struct {
	DurationType string `xml:"durationType"`
	Dots         int    `xml:"dots"`
	Notes        []Note `xml:"Note"`
}

type Note struct {
	Pitch int `xml:"pitch"`
}
