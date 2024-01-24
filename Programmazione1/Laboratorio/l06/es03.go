package main

import "fmt"

func main() {
	stringheUnite := ""
	s := ""

	for s != "0" {
		_, err := fmt.Scan(&s)

		if err != nil {
			fmt.Println("Errore in input")
			return
		}

		if s == "0" {
			stringheUnite = stringheUnite[:len(stringheUnite)] // rimuove lo spazio scritto precedentemente
		} else {
			stringheUnite += s + " "
		}
	}

	fmt.Println(stringheUnite)
}
