package main

import . "fmt"

func main() {
	var s string
	Scan(&s)

	var sMinuscola, sMaiuscola string

	for _, c := range s {
		if c >= 65 && c <= 90 { // c è maiuscola
			sMaiuscola += string(c)
			sMinuscola += string(c + 32)
		} else if c >= 97 && c <= 122 { // c è minuscola
			sMaiuscola += string(c - 32)
			sMinuscola += string(c)
		}
	}

	Println("Minuscola: ", sMinuscola)
	Println("Maiuscola: ", sMaiuscola)
}
