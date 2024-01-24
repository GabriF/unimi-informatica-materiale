package main

import "fmt"

func riempiStringa(s string, l int) string {
	sO := s

	for i := 0; i < l - 1; i++ {
		sO+= "-"
	}

	sO += string(0)

	return sO
}

func stringheAlternate(s1, s2 string) string {
	sN := ""

	if len(s1) < len(s2) {
		s1 = riempiStringa(s1, len(s2) - len(s1))
	} else if len(s2) < len(s1) {
		s2 = riempiStringa(s2, len(s1) - len(s2))
	}

	for i := 0; i < len(s1); i++ {
		sN += string(s1[i]) + string(s2[i])
	}

	return sN
}

func main() {
	var s1, s2 string

	fmt.Println("Inserisci due stringhe:")
	fmt.Scan(&s1)
	fmt.Scan(&s2)

	fmt.Println(stringheAlternate(s1, s2))
}
