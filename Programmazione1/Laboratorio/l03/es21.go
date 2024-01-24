package main

import "fmt"

func main() {
	fmt.Print("Inserisci una sequenza di numeri (termina con n<=0): ")

	var n, s, i int = 1, 0, -1
	for n > 0 {
		fmt.Scan(&n)

		s += n
		i++;
	}

	fmt.Println("Media: ", s/i)
}
