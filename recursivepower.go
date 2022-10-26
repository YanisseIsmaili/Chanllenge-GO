package piscine

func RecursivePower(nb int, power int) int {
	if nb == 0 && power == 0 {
		return 1
	} else if nb == 0 {
		return 0
	} else if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	a := nb
	a *= RecursivePower(nb, power-1)
	return a
}
