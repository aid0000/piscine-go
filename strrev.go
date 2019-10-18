package piscine

func StrRev(s string) string {
	var a int

	for index := range s {
		a = index
	}

	wordInBytes := []byte(s)
	j := 0
	for i := a; i >= 0; i-- {
		wordInBytes[i] = byte(s[j])
		j++
	}
	FinalStr := string(wordInBytes)
	return FinalStr
}
