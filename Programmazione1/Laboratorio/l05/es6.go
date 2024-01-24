package main

import "fmt"

func main() {
	var base rune = '\U0001F0B1'

	for i := 0; i < 10; i++ {
		fmt.Println("Simbolo:", string(base), "codice numerico:", base)
		base++
	}
}
