package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func NumeroNascosto(testo string) (int, error) {
	var s string

	for _, c := range testo {
		if unicode.IsDigit(c) {
			s += string(c)
		}
	}

	return strconv.Atoi(s)
}

func main() {
	var si string

	fmt.Print("Inserisci una stringa di testo:")
	fmt.Scan(&si)

	nNascosto, errore := NumeroNascosto(si)

	if errore != nil {
		fmt.Println("Errore nella conversione")
		return
	}

	fmt.Println("Doppio del numero nascosto:", nNascosto * 2)
}
