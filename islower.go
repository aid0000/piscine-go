package piscine

func IsLower(str string) bool {
	bool := true

	for _, value := range str {
		if value >= 'a' && value <= 'z' {
			bool = bool
		} else {
			bool = false
		}
	}
	return bool
}
