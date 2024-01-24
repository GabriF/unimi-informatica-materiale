package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&n)

	for i, cor, cor2 := 0, 0, 1; i < n; i++ {
		cor, cor2 = cor + cor2, cor

		fmt.Println(cor)
	}
}
