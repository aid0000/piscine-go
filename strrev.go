package piscine

func StrRev(s string) string {
	var a int

	for index := range s {
		a = index
	}

	Word := []byte(s)
	b := 0
	for i := a; i >= 0; i-- {
		Word[i] = byte(s[b])
		b++
	}
	Final := string(Word)
	return Final
}
