package main

import "github.com/01-edu/z01"

func PrintComb() {

	for i := 48; i <= 57; i++ {
		a := rune(j)

		for j := 48; j <= 57; j++ {
			b := rune(k)

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
	z01.PrintRun(36)
}

/*
i = 0	i+48 (valeur aski)
J = 0
K = 0

for = boucle
X++ = augmentation de 1
X-- = diminution de 1

for i=0; i<=9; i++				(Boucle i)
	for j=0; j<=9 j++			(Boucle j)
		for k=0; i<=9 k++		(Boucle K)
			if i < j && j < k    (regle (i) plus peti que (j) et = (&&) (j) plus petit que (k))
				print (i,j,k)

*/
