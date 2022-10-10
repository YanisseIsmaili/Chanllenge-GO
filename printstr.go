package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	s_rune := []rune(s)
	size := len(s_rune)
	for i := 0; i < (size); i++ {
		z01.PrintRune(s_rune[i])
	}
	
}
