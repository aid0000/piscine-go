package piscine

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}
	a := []rune(str)
	for _, letter := range a {
		if letter >= 32 && letter <= 126 {
			continue
		} else {
			return false
		}
	}
	return true
}
