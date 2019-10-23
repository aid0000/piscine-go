package piscine

func IsUpper(str string) bool {
	bool := true

	for _, value := range str {
		if value >= 'A' && value <= 'Z' {
			bool = bool
		} else {
			bool = false
		}
	}
	return bool
}
