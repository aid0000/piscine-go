package piscine

func AlphaCount(str string) int {
	l := 0
	for _, c := range str {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			l++
		}
	}
	return l
}
