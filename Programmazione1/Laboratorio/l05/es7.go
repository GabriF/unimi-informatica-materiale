package main

import "fmt"

func main() {
	var s string

	fmt.Println("Inserisci una stringa:")
	fmt.Scan(&s)

	for _, r := range s {
		fmt.Print(string(r), " ")
	}

	fmt.Println()
}
