package main

import . "fmt"

func main() {
	Print("Inserisci la stringa: ")
	var s string
	Scan(&s)

	var sGir string
	for _, c := range s {
		if c >= 65 && c <= 90 {
			c += 32
		}

		sGir = string(c) + sGir
	}

	if s == sGir {
		Println("La parola è palindroma")
	} else {
		Println("La arola non è palindroma")
	}
}
