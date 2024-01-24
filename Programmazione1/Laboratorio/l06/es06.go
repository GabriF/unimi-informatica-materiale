package main

import "fmt"

func main() {
	var si string = ""
	var cognome string

	for si != "0" {
		fmt.Scan(&si)

		if cognome < si {
			cognome = si
		}
	}

	fmt.Println(cognome)
}
