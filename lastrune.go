package piscine

func NRune(s string, n int) rune {
	for index, r := range []rune(s) {
		if index+1 == n {
			return r
		}
	}
	return '\x00'
}
