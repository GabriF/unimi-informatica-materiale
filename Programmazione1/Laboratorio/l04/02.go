package main

import . "fmt"

func main() {
	Print("Inserisci il lato: ")
	var lato int
	Scan(&lato)

	var carattere rune
	for i := 0; i < lato; i++ {

		if i % 2 == 0 {
			carattere = '*'
		} else {
			carattere = '+'
		}

		for j := 0; j < lato; j++ {
			Print(string(carattere), " ")
		}

		Println()
	}
}
