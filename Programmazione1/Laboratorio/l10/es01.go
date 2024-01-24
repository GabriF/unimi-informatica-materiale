package main

import (
	. "fmt"
	"bufio"
	"os"
	"unicode"
	"strings"
)

func LeggiTesto() string {
	var s string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		letto := scanner.Text()
		s += letto
	}

	return s
}

func Occorrenze(s string) map[rune]int {
	var occorrenze map[rune]int = make(map[rune]int)

	for _, r := range s {
		if unicode.IsLetter(r) {
			occorrenze[r]++
		}
	}

	return occorrenze
}

func main() {
	testo := LeggiTesto()
	occorrenze := Occorrenze(testo)

	for chiave, valore := range occorrenze {
		Print(string(chiave), ": ", strings.Repeat("*", valore), "\n")
	}
}
