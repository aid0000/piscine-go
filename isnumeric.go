package piscine

func IsNumeric(str string) bool {
	bool := true

	for _, value := range str {
		if value >= '0' && value <= '9' {
			bool = bool
		} else {
			bool = false
		}
	}
	return bool
}
