package main

import . "fmt"
import "os"
import "strings"

func Occorrenze(s string) [26]int {
	var occorrenze [26]int

	for _, c := range strings.ToLower(s) {
		posizione := c - 97
		occorrenze[posizione]++
	}

	return occorrenze
}

func main() {
	arg := os.Args[1]
	occorrenze := Occorrenze(arg)

	for i, o := range occorrenze {
		Print(string(i + 97), " ", o, " ")
	}

	Println()
}
