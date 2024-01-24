package main

import (
	"fmt"
	"math"
)

func main() {
	var eta1, eta2 int

	fmt.Scan(&eta1, &eta2)

	media := float64((eta1 + eta2) / 2)
	somma := eta1 + eta2

	mediaDifetto := math.Floor(media)
	mediaEccesso := math.Ceil(media)

	somma10 := somma + 20
	media10 := media + 10

	fmt.Println("Somma =", somma, "\nMedia =", media)
	fmt.Println("Media per difetto =", mediaDifetto, "\nMedia per eccesso =", mediaEccesso)
	fmt.Println("Somma tra 10 anni =", somma10, "\nMedia 10 anni =", media10)
}
