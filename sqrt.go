package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb == 2{
		return 0
	} 
	if nb > 0 {
		result := 1
		Sqrt := 0
		for a := 1; result <= nb; a++{
			result = a * a
			Sqrt++
			if result == nb {
				return Sqrt
			}
		}
		return 0
	} else {
		return 0
	}
}
