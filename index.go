package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	x := 0
	y := 0

	for index := range a {
		index = index
		x++
	}
	for index := range b {
		index = index
		y++
	}
	if y == 0 {
		return 0
	}

	for index, value := range a {
		if value == b[0] && x >= y+index-1 {
			run := 1
			for i := 1; i < y; i++ {
				if b[i] == a[index+i] {
					run++
				}
			}
			if run == y {
				return index
			}
		}
	}
	return -1
}
