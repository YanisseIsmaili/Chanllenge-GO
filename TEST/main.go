package main

import "github.com/01-edu/z01"

func main() {

	for i := 48; i <= 57; i++ {
		a := rune(i)

		for j := 48; j <= 57; j++ {
			b := rune(j)

			for k := 48; k <= 57; k++ {
				c := rune(k)

				if i < j && j < k {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		}
	}
	z01.PrintRune(36)
}
