package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	var numero_generato int = rand.Intn(100)

	var n int = -1
	for n != numero_generato {
		fmt.Print("Indovina il numero: ")
		fmt.Scan(&n)

		if n > numero_generato {
			fmt.Println("Hai scritto un numero superiore")
		} else if n < numero_generato {
			fmt.Println("Hai scritto un numero inferiore")
		}
	}
}
