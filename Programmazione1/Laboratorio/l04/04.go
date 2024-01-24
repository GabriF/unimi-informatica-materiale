package main

import . "fmt"

func main() {
	var s string
	Scan(&s)

	var vocali, consonanti, minuscole, maiuscole int

	for _, c := range s {
		if c >= 65 && c <= 90 {
			maiuscole++
		} else if c >= 97 && c <= 122 {
			minuscole++
		} else {
			return
		}

		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vocali++
		} else if c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
			vocali++
		} else {
			consonanti++
		}
	}

	Println("Vocali: ", vocali)
	Println("Consonanti: ", consonanti)
	Println("Minuscole: ", minuscole)
	Println("Maiuscole: ", maiuscole)
}
