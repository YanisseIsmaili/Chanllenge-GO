package piscine

func AlphaCount(s string) int {
	my_array := []rune(s)
	k := 0
	for _, r := range my_array { // for r := 0; r <= len(my_array-1) r++
		if r >= 65 && r <= 90 || r >= 97 && r <= 122 {
			k++
		}
	}
	return k
}

/*
- resulta : le nombre de caracteres de l'alphabet qu'il y a au total
- 1 condition : ne doit pas avoir les nombre 
- 2 condition : pas  avoir de chifre / caractere spesiaux espce vide

string = nombre de charactere 
k = conteur permeten de compter les characteur alphabetique
int = si tu utilise des nombre infini - infini + nobre negatif pas de virguele

- for = parcourire un tableaux
- (_) = pas de changeman du tableaux 
- (i) = index du tableaux
- (r) = variable
- range = taille
- my_array = mon tableaux de rune

- if = donner des condition
- r = variable definisant la taille de mon tableaux
- >= = superieur ou egale 
- 65 = valeur ascci A maj
- && = et
- <= = inferieur et egal
- 90 = valeur ascci Z maj
-  || = ou
- r = variable definisant la taille de mon tableaux
- >= = plus grant et egal
- 97 = valeur ascci de a en miniscul
- && = et 
- r = variable definisant la taille de mon tableaux
- <= = superieur et egale
- 122 = valeur ascci de z miniscul
 
- ! = diferent
- k++ = ajouter 1 a k
- == = strictmen egale
- += = ajout d une valeur a une valeur
- -= = retire une valeur a une valeur
/*
