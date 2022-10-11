package piscine

func IsNumeric(s string) bool {
    t := []rune(s)
	a := true
	for _, r := range t {
		if !(r >= 48 && r <= 57) {
			a = false
		}
	}
+	return a
}
