package main

import (
	. "fmt"
)

func main() {
	var inputUtente []int

	var q int
	Print("Quanti numeri vuoi inserire: ")
	Scan(&q)

	Println("Inserisci", q, "numeri")
	for i, n := 0, 0; i < q; i++ {
		Scan(&n)
		inputUtente = append(inputUtente, n)
	}

	Println("Numeri in ordine inverso:")
	for i := 0; i < q; i++ {
		Print(inputUtente[q - i - 1], " ")
	}

	Println()
}
