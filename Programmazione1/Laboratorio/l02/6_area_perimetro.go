package main

import "fmt"

func main() {
	var altezza, base float64

	fmt.Print("Inserisci altezza e base di un rettangolo: ")
	fmt.Scan(&altezza, &base)

	fmt.Println("Area =", base * altezza)
	fmt.Println("Perimetro =", (base + altezza) * 2)
}
