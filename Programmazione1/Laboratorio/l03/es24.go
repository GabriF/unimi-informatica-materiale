package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&n)

	fmt.Print("Divisori propri di ", n, ": ")
	for i := 1; i < n; i++ {
		if n % i == 0 {
			fmt.Print(i, " ")
		}
	}

	fmt.Println()
}
