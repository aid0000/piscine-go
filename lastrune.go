package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	var lastRune rune
	for _, r := range runes {
		lastRune = r
	}
	return lastRune
}
