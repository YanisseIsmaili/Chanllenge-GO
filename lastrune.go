package piscine

func LastRune(s string) rune {
	s_rune:= []rune(s)
	size := len(s_rune)
	return s_rune[size-1]
}
