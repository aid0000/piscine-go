package strlen

func StrLen(str string) int {
	var nb int = 0
	for range str {
		nb++
	}
	return nb
}
