package piscine

func IterativePower(nb int, power int) int {
	if nb == 0 && power == 0 {
		return 1
	} else if nb == 0 {
		return 0
	} else if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	a := 1
	for i := 0; i < power; i++ {
		a *= nb
	}
	return a
}
