package shawzin

import "testing"

func Test_IntToASCIIPos(t *testing.T) {
	cases := []struct {
		in       int
		expected string
	}{
		{0, "AA"},
		{1, "AB"},
		{25, "AZ"},
		{26, "Aa"},
		{50, "Ay"},
		{51, "Az"},
		{52, "A0"},
		{61, "A9"},
		{62, "BA"},
	}

	for _, c := range cases {
		actual := IntToASCIIPos(c.in)
		if actual != c.expected {
			t.Errorf("IntToASCIIPos(%d) == %s, expected %s, got %s", c.in, actual, c.expected, actual)
		}
	}
}
