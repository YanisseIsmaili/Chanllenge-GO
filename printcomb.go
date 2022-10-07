package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if i < j && j < k {
					z01.PrintRune(get_digit(i))
					z01.PrintRune(get_digit(j))
					z01.PrintRune(get_digit(k))
					if i == 7 && j == 8 && k == 9 {
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}

func get_digit(number int) rune {
	return rune(number + 48)
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
