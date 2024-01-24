package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&n)

	if n % 10 != 0 || n == 0 {
		fmt.Println(n, "non è multiplo di 10")
	} else {
		fmt.Println(n, "è multiplo di 10")
	}
}
