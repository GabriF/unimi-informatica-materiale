package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n_lati int
		lunghezza float64
	)

	fmt.Print("Scrivi il numero di lati: ")
	fmt.Scan(&n_lati)

	fmt.Print("Scrivi la lunghezza dei lati: ")
	fmt.Scan(&lunghezza)

	var area float64 = (float64(n_lati) * math.Pow(lunghezza, 2)) / (4 * math.Tan(math.Pi/float64(n_lati)))
	fmt.Println("Area = ", area)
}
