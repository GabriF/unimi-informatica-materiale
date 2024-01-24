package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci il numeri di cui vuoi la tabellina: ")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Println(n, "x", i, "=", n * i)
	}
}
