package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	mol := 0
	for _, c := range args {
		numero, err := strconv.Atoi(c)

		if err == nil {
			if mol == 0 {
				mol = 1
			}

			mol *= numero
		}
	}

	Println("La moltiplicazione tra i numeri è: ", mol)
}
