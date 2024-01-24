package main

import "fmt"
import "strings"

func main() {
	var n int

	fmt.Print("Inserisci un intero maggiore di 0: ")
	_, err := fmt.Scan(&n)

	if err != nil || n <= 0 {
		fmt.Println("Errore di input")
		return
	}

	for i := 0; i < n; i++ {
		c := strings.Repeat("*", i + 1)

		fmt.Println(c)
	}
}
