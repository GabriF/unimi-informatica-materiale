package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&n)

	if n != 0 {
		if n % 3 == 0 && n % 5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if n % 3 == 0 {
			fmt.Println("Fizz")
		} else if n % 5 == 0 {
			fmt.Println("Buzz")
		}
	}
}
