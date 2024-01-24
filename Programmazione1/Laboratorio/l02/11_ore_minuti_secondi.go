package main

import "fmt"

func main() {
	fmt.Print("Numero di secondi: ")

	var s int
	fmt.Scan(&s)

	minuti := s / 60 % 60
	ore := s / 60 / 60
	secondi := s % 60

	fmt.Print("h:m:s ", ore, ":", minuti, ":", secondi, "\n")
}
