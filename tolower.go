package piscine

func ToLower(s string) string {
	a := []rune(s)
	result := ""

	for index, value := range a {
		if value >= 'a' && value <= 'z' {
			value = value
		} else if value >= 'A' && value <= 'Z' {
			a[index] = a[index] + 32
		}
		result += string(a[index])
	}
	return result
}
