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

func numeriGemelli(soglia int) {
	for p := 2; p < soglia; p++ {
		q := p - 2

		if ePrimo(q) && ePrimo(p) {
			fmt.Println(q, "e", p, "sono numeri primi gemelli")
		}
	}
}

func main() {
	var soglia int

	fmt.Scan(&soglia)

	if soglia <= 0 {
		fmt.Println("La soglia inserita non è positiva")
		return
	}

	numeriGemelli(soglia)
}
