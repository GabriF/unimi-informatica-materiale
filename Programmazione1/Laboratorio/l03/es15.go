package main

import "fmt"

func main() {
	var l1, l2 int;

	fmt.Print("Inserisci l'ampiezza di due angoli di un triangolo: ")
	fmt.Scan(&l1, &l2)

	if l1 + l2 >= 180 {
		fmt.Println("I due angoli formano già 180 gradi")
	} else if l1 <= 0 || l2 <= 0 {
		fmt.Println("I due angoli devono essere maggiori di 0")
	} else {
		fmt.Println("Il terzo angolo misura", 180 - l1 - l2)
	}
}
