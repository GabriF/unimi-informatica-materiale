package main

import "fmt"

func main() {
	var n int

	fmt.Print("Scrivi un numero intero: ")
	fmt.Scan(&n)

	if n % 2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}
