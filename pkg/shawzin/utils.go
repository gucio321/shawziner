package shawzin

import "fmt"

// this converts  intager into ASCII-encoded position
// for example 0 becomes 'A', 1 becomes 'B' and so on
// Up to 'zz' which is 57*57 = 3249
func IntToASCIIPos(i int) string {
	i = i
	const NewBase = 'Z' - 'A' + ('z' - 'a') + 2

	if i < 0 || i > NewBase*NewBase {
		// panic("Invalid input")
	}

	fmt.Println(i)
	first := i / NewBase
	fmt.Println(first)
	second := i % NewBase
	fmt.Println(second)

	if first > NewBase/2 {
		first += 'a' - 'Z' - 1
	}

	if second > NewBase/2-1 {
		second += 'a' - 'Z' - 1
	}

	return string('A'+rune(first)) + string('A'+rune(second))
}
