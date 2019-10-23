package piscine

func ToUpper(s string) string {

	a := []rune(s)
	result := ""

	for index, value := range a {
		if value >= 'A' && value <= 'Z' {
			value = value
		} else if value >= 'a' && value <= 'z' {
			a[index] = a[index] - 32
		}
		result += string(a[index])
	}
	return result
}
