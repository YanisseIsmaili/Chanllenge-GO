package piscine

func RecursiveFactirial(nb int) int {
	a := 1
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb <= 25 {
		nb *= RecursiveFactirial(nb - 1)
	} else {
		return 0
	}
	return a
}
