package goMusic

type Duration int

const (
	// Well, minimal Shawzin's duration is 1/32 of a note, so
	WholeNote        Duration = 32
	HalfNote         Duration = 16
	QuarterNote      Duration = 8
	EighthNote       Duration = 4
	SixteenthNote    Duration = 2
	ThirtySecondNote Duration = 1
)

func (d *Duration) Dot() *Duration {
	*d = *d * 3 / 2
	return d
}
