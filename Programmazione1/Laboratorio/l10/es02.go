package main

import (
	. "fmt"
	"bufio"
	"os"
	"unicode"
	"sort"
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

func StampaIstogramma(occorrenze map[rune]int) {
	chiaviOrdinate := make([]string, len(occorrenze))

	var i int
	for c, _ := range occorrenze {
		chiaviOrdinate[i] = string(c)
		i++
	}

	sort.Strings(chiaviOrdinate)
	for _, e := range chiaviOrdinate {
		Print(string(e), ": ", strings.Repeat("*", occorrenze[rune(e[0])]), "\n")
	}
}

func main() {
	testo := LeggiTesto()
	occorrenze := Occorrenze(testo)
	StampaIstogramma(occorrenze)
}
