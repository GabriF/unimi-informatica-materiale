package main

import . "fmt"

func main() {
	Print("Inserisci il lato: ")
	var lato int
	Scan(&lato)

	for i := 0; i < lato; i++ {
		for j := 0; j < lato; j++ {
			Print("* ")
		}

		Println()
	}
}
