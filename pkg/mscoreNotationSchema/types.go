package mscoreNotationSchema

import (
	"encoding/xml"
	"io"
)

type MUEFile struct {
	Measures []Measure `xml:"Score>Staff>Measure"`
}

type Measure struct {
	Chords []Chord `xml:"voice>Chord,voice>Rest"`
}

func (m *Measure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, err := d.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		switch t := t.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "Chord":
				var c Chord
				if err := d.DecodeElement(&c, &t); err != nil {
					return err
				}
				m.Chords = append(m.Chords, c)
			case "Rest":
				var r Rest
				if err := d.DecodeElement(&r, &t); err != nil {
					return err
				}
				m.Chords = append(m.Chords, Chord{DurationType: r.DurationType})
			}
		}
	}
}

type Rest struct {
	DurationType string `xml:"durationType"`
}

type Chord struct {
	DurationType string `xml:"durationType"`
	Dots         int    `xml:"dots"`
	Notes        Note   `xml:"Note"`
}

type Note struct {
	Pitch int `xml:"pitch"`
}
