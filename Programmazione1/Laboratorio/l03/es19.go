package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Inserisci due numeri interi pari: ")
	fmt.Scan(&a, &b)

	var vp int = a
	if a % 2 != 0 {
		vp -= 1
	}

	var s int = 0;
	for i := vp + 2; i > a && i < b; i += 2 {
		s += i;
	}
	fmt.Println("La somma dei numeri pari tra", a, "e", b, "è", s)
}
