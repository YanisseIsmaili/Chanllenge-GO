package piscine

func BasicJoin(elems []string) string {
	var final_string string
	for _, s := range elems {
		final_string += s
	}
	return final_string
}
