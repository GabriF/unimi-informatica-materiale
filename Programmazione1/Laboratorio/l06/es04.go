package main

import "fmt"
import "strings"

func main() {
	var s string

	for s != "0" {
		fmt.Print("Inserisci una stringa: ")

		_, err := fmt.Scan(&s)
		if err != nil {
			fmt.Println("Errore in input")
			return
		}

		if s != "0" {
			fmt.Println("Stringa in maiuscolo: ", strings.ToUpper(s))
		}
	}
}
