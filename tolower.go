package piscine

func ToLower(s string) string {
	t := []rune(s)
	for i := 0; i < len(t); i++ {
		if t[i] >= 65 && t[i] <= 90 {
			t[i] += 32
		}
	}
	return string(t)
}
