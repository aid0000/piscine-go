package piscine

func StrLen(str string) int {
	l := 0
	for _, c := range str {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			letter++
		}
	}
	return l
}
