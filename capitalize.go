package piscine

func Capitalize(s string) string {
	a := []rune(s)
	var result string
	k := 0

	for i := range s {
		i = i
		k++
	}

	if a[0] >= 'a' && a[0] <= 'z' {
		a[0] = a[0] - 32
	}

	for i := 1; i < k; i++ {
		if a[i] >= 'A' && a[i] <= 'Z' {
			if a[i-1] >= 'A' && a[i-1] <= 'Z' {
				a[i] = a[i] + 32
			}
			if a[i-1] >= 'a' && a[i-1] <= 'z' {
				a[i] = a[i] + 32
			}
			if a[i-1] >= '0' && a[i-1] <= '9' {
				a[i] = a[i] + 32
			}
		}
		if a[i] >= 'a' && a[i] <= 'z' {
			if a[i-1] >= 'A' && a[i-1] <= 'Z' {
				continue
			}
			if a[i-1] >= 'a' && a[i-1] <= 'z' {
				continue
			}
			if a[i-1] >= '0' && a[i-1] <= '9' {
				continue
			} else {
				a[i] = a[i] - 32
			}
		}
	}
	result = string(a)
	return result
}
