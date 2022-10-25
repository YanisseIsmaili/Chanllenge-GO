package piscine

func RecursiveFactorial(nb int) int {
	a := 1
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb <= 25 {
		nb *= RecursiveFactorial(nb - 1)
	} else {
		return 0
	}
	return a
}
