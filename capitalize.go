package piscine

func Capitalize(s string) string {
	tab := []rune(s)
	for i, t := range tab {
		if lettrenumber(t) {
			if i == 0 || lettrenumber(tab[i-1]) == false {
				if tab[i] >= 97 && tab[i] <= 122 {
					tab[i] = t - 32
				}
			} else {
				if tab[i] >= 65 && tab[i] <= 90 {
					tab[i] = t + 32
				}
			}
		}
	}
	return string(tab)
}

func lettrenumber(r rune) bool {
	if r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122 {
		return true
	} else {
		return false
	}
}
