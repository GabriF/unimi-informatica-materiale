package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Inserisci due numeri interi: ")
	fmt.Scan(&a, &b)

	if b == 0 {
		fmt.Println("Impossibile")
	} else {
		fmt.Println("a/b =", float32(a) / float32(b))
	}
}
