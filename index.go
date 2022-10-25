package piscine


func Index(s string, toFind string) int {
	var index int
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			index = i
			if len(toFind) == 1 {
				return index
			}
			for j := 1; (j < len(toFind) && j+i < len(s)); j++ {
				if s[i] != toFind[j] {
					return -1
				} else if s[i] == toFind[len(toFind)-1] {
					return index
				}
			}
		}
	}
	return index
}

/*
func Index(s string, toFind string) int {
	strtofind := NbretoFind(toFind)
	secondIndex := 0
	for _, j := range toFind {
		for i1, i2 := range s {
			if i2 == j {
				if strtofind == 1 {
					return i1
				} else if strtofind > 1 {
					for k := 0; k < strtofind; k++ {
						if s[i1+k] == toFind[k] {
							secondIndex++
						} else {
							return -1
						}
					}
					if secondIndex == strtofind {
						return i1
					}
				} else {
					return -1
				}
			}
		}
		if secondIndex <= 0 {
			return -1
		}
	}
	return strtofind
}

func NbretoFind(s string) int {
	compteur := 0
	for i := range s {
		compteur = i + 1
	}
	return compteur
}
*/
