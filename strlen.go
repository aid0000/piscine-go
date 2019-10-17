package strlen

func StrLen(str string) int {
	var count int = 0
	for range str {
		count++
	}
	return count
}
