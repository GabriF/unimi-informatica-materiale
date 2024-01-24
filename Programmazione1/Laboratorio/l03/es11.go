package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci un voto tra 0 e 100: ")
	fmt.Scan(&n)

	if n < 0 || n > 100 {
		fmt.Println("Errore")
	} else if n < 60 {
		fmt.Println("Insuficciente")
	} else if n >= 60 && n < 70 {
		fmt.Println("Sufficiente")
	} else if n >= 70 && n < 80 {
		fmt.Println("Buono")
	} else if n >= 80 && n < 90 {
		fmt.Println("Distinto")
	} else { // n >= 90 && n <= 100
		fmt.Println("Ottimo")
	}
}
