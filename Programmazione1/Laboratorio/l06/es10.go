package main

import "fmt"

func base10ToBaseX(n, b uint) (string, bool) {
	if b < 2 || b > 16 {
		return "", false
	}

	var nv string

	caratteri := "0123456789ABCDEF"

	for n != 0 {
		resto := n % b
		n = n / b
		nv = string(caratteri[resto]) + nv
	}

	return nv, true
}

func main() {
	fmt.Print("Inserisci la base: ")
	var b uint
	fmt.Scan(&b)

	fmt.Print("Inserisci il numero: ")
	var n uint
	fmt.Scan(&n)

	r, f := base10ToBaseX(n ,b)

	if ! f {
		fmt.Println("Errore nella conversione")
		return
	}

	fmt.Println(n, "in base", b, "si scrive", r)
}
