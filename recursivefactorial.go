package piscine

func RecursiveFactorial(nb int) int {
	if nb > 1 && nb <= 20 {
		return nb * RecursiveFactorial(nb-1)
	} else {
		return 0
	}
}
