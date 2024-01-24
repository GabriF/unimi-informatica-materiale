package main

import (
	. "fmt"
	"os"
	_ "unicode"
	"strconv"
)

func main() {
	args := os.Args[1:]

	var prec int
	for i, c := range args {
		intero, err := strconv.Atoi(c)

		// controlla se la stringa è un intero
		if err != nil {
			Println("Valore in posizione", i, "non valido")
			return
		}

		// controlla se il numero in posizione dispari è minore del precedente e se il numero in posizioen pari è maggiore del precedente
		if i % 2 == 0 {
			if intero <= prec && i != 0 {
				Println("Valore in posizione", i, "non valido")
				return
			}
		} else if intero >= prec { // i % 2 == 1
			Println("Valore in posizione", i, "non valido")
			return
		}

		prec = intero
	}

	Println("Sequenza valida")
}
