package piscine

func TrimAtoi(s string) int {
	runes := []rune(s)
	sign := 1
	result := 0
	foundDigit := false
	for _, r := range runes {
		if !foundDigit {
			if r == '-' {
				sign = -1
			} else if r == '+' {
				sign = 1
			}
		}
		if r >= '0' && r <= '9' {
			foundDigit = true
			result = 10*result + int(r-'0')
		}
	}
	return sign * result
}
