package main

import (
	. "fmt"
	"bufio"
	"os"
	"unicode"
	"io"
	"sort"
)

func LeggiParole(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var parole []string
	for scanner.Scan() {
		riga := scanner.Text()

		parola := ""
		for _, c := range riga {
			if unicode.IsSpace(c) {
				if parola != "" {
					parole = append(parole, parola)
				}
				parola = ""
			} else {
				parola += string(c)
			}
		}

		if parola != "" {
			parole = append(parole, parola)
		}
	}

	return parole
}

func main() {
	parole := LeggiParole(os.Stdin)

	var occorrenze map[string]int = make(map[string]int)
	for _, p := range parole {
		occorrenze[p]++
	}

	var paroleUniche []string
	for p, cont := range occorrenze {
		if cont == 1 {
			paroleUniche = append(paroleUniche, p)
		}
	}

	sort.Strings(paroleUniche)
	Println(paroleUniche)
}
