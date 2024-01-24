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
		s += " " + letto
	}

	return s[1:]
}

func SeparaParole(s string) []string {
	var l []string

	var parola string
	for _, c := range s {
		if parola != "" && ! unicode.IsLetter(c) {
			l = append(l, parola)
			parola = ""
		} else if unicode.IsLetter(c) {
			parola += string(c)
		}
	}

	if parola != "" {
		l = append(l, parola)
	}

	return l
}

func Occorrenze(sl []string) map[string]int {
	var occorrenze map[string]int = make(map[string]int)

	for _, p := range sl {
		occorrenze[p]++
	}

	return occorrenze
}

func main() {
	testo := LeggiTesto()
	parole := SeparaParole(testo)
	occorrenze := Occorrenze(parole)

	for chiave, valore := range occorrenze {
		Print(chiave, ": ", strings.Repeat("*", valore), "\n")
	}
}
