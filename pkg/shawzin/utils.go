package shawzin

func asciiTable() map[int]string {
	return map[int]string{
		0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H", 8: "I", 9: "J",
		10: "K", 11: "L", 12: "M", 13: "N", 14: "O", 15: "P", 16: "Q", 17: "R", 18: "S", 19: "T",
		20: "U", 21: "V", 22: "W", 23: "X", 24: "Y", 25: "Z", 26: "a", 27: "b", 28: "c", 29: "d",
		30: "e", 31: "f", 32: "g", 33: "h", 34: "i", 35: "j", 36: "k", 37: "l", 38: "m", 39: "n",
		40: "o", 41: "p", 42: "q", 43: "r", 44: "s", 45: "t", 46: "u", 47: "v", 48: "w", 49: "x",
		50: "y", 51: "z", 52: "0", 53: "1", 54: "2", 55: "3", 56: "4", 57: "5", 58: "6", 59: "7",
		60: "8", 61: "9",
	}
}

// this converts  intager into ASCII-encoded position
// for example 0 becomes 'A', 1 becomes 'B' and so on
// Up to 'zz' which is 57*57 = 3249
func IntToASCIIPos(i int) string {
	data := asciiTable()

	const NewBase = 'Z' - 'A' + ('z' - 'a') + ('9' - '0') + 3
	first := i / NewBase
	second := i % NewBase

	firstS, ok := data[first]
	if !ok {
		panic("Invalid input")
	}

	secondS, ok := data[second]
	if !ok {
		panic("Invalid input")
	}

	return firstS + secondS
}

/*
func IntToASCIIPos(i int) string {
	i = i
	const NewBase = 'Z' - 'A' + ('z' - 'a') + ('9' - '0') + 3

	if i < 0 || i > NewBase*NewBase {
		// panic("Invalid input")
	}

	first := i / NewBase
	second := i % NewBase

	if first >= 'Z'-'A' {
		first += 'a' - 'Z' - 1
	}

	if first >= 'Z'-'A' {
		first +=
	}

	switch {
	case second >= 'Z'-'A':
		second += 'a' - 'Z' - 1
	}

	return string('A'+rune(first)) + string('A'+rune(second))
}
*/
