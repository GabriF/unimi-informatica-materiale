package main

import (
	"fmt"
	"math"
)

func main() {
	var raggio float64

	fmt.Print("Inserisci il raggio: ")
	fmt.Scan(&raggio)

	fmt.Println("Area =", raggio * raggio * math.Pi)
	fmt.Println("Circonferenza =", 2 * raggio * math.Pi)
}
