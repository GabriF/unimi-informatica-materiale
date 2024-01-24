package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int

	fmt.Print("Inserisci due interi: ")
	fmt.Scan(&x, &y)

	var maggiore, minore, somma, differenza, prodotto, divisione, media, potenza_for, potenza_pow int

	if x > y {
		maggiore = x
		minore = y
	} else {
		maggiore = y
		minore = x
	}

	somma = x + y

	differenza = x - y

	prodotto = x * y

	if y != 0 {
		divisione = x / y
	} else {
		divisione = -1
	}

	media = (x + y) / 2

	if y == 0 {
		potenza_for = 1
	} else {
		potenza_for = x;
		for i := 0; i < y; potenza_for, i = potenza_for * x, i + 1 {}
	}

	potenza_pow = int(math.Pow(float64(x), float64(y)))

	fmt.Println("Maggiore: ", maggiore)
	fmt.Println("Minore: ", minore)
	fmt.Println("Somma: ", somma)
	fmt.Println("Differenza: ", differenza)
	fmt.Println("Prodotto: ", prodotto)
	fmt.Println("Divisione: ", divisione)
	fmt.Println("Media: ", media)
	fmt.Println("Potenza for: ", potenza_for)
	fmt.Println("Potenza pow: ", potenza_pow)
}
