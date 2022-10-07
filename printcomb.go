package piscine

import "github.com/01-edu/z01"

func printcomb() {
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
					if i == 55 && j == 56 && k == 57 {
						z01.PrintRune(10)
					} else {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
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
32