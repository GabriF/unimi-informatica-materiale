package main

import . "fmt"

func main() {
	Print("Inserisci il lato: ")
	var lato int
	Scan(&lato)

	var carattere rune
	for i := 0; i < lato; i++ {

		if i % 3 == 0 {
			carattere = '*'
		} else if i % 3 == 1 {
			carattere = '+'
		} else if i % 3 == 2 {
			carattere = 'o'
		}

		for j := 0; j < lato; j++ {
			Print(string(carattere), " ")
		}

		Println()
	}
}
