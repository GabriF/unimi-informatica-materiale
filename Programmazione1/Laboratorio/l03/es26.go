package main

import "fmt"

func main() {
	var prec, n int

	fmt.Print("Inserisci una sequenza di numeri (termina con <=-1): ")
	fmt.Scan(&n)
	for n >= -1 {
		prec = n
		fmt.Scan(&n)

		if n > prec {
			fmt.Println("Crescente")
		} else {
			fmt.Println("Decrescente")
		}
	}
}
