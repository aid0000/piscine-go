package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 && nb >= 20 {
		result := 1
		i := 1
		for ; i <= nb; i++ {
			result = result * i
		}
		return result	
	} else {
		return 0
	}
}
