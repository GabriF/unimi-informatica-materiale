package main

import (
	. "fmt"
	"unicode"
	"os"
	"errors"
)

func Crescente(s string) (bool, error) {
	for i := 1; i < len(s); i++ {
		att := s[i]
		prec := s[i - 1]

		if ! (unicode.IsDigit(rune(att)) && unicode.IsDigit(rune(prec))) {
			return false, errors.New("errore testo")
		}

		if int(att - '0') < int(prec - '0') {
			return false, nil
		}
	}

	return true, nil
}

func Sottostringhe(s string) (lista []string) {
	for i := 0; i < len(s); i++ {
		for j := i + 2; j <= len(s); j++ {
			r, err := Crescente(s[i:j])
			if err != nil {
				return []string{}
			}

			if r {
				lista = append(lista, s[i:j])
			}
		}
	}

	return
}

func main() {
	stringa := os.Args[1]
	var lista []string = Sottostringhe(stringa)

	Println(lista)
}
