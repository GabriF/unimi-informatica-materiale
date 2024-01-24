package main

import "fmt"

func ePrimo(numero int) bool {
	if numero == 0 {
		return false
	} else if numero == 1 {
		return true
	}

	for i := 2; i < numero; i++ {
		if numero % i == 0 {
			return false
		}
	}

	return true
}

func numeriPrimi(soglia int) {
	for i := 0; i < soglia; i++ {
		if ePrimo(i) {
			fmt.Println(i, " è primo")
		}
	}
}

func main() {
	var soglia int

	fmt.Scan(&soglia)

	numeriPrimi(soglia)
}
