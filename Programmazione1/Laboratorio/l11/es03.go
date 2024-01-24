package main

import (
	. "fmt"
	"os"
	_ "unicode"
	"strings"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var numeri []string
	for _, v := range os.Args[1:] {
		numero := v
		segno := "+"

		if v[0] == '-' || v[0] == '+' {
			numero = v[1:]
			segno = string(v[0])
		}

		numeri = append(numeri, segno + numero)
	}

	sottoSequenze := make(map[string]int)
	for i := 0; i < len(numeri); i++ {
		for j := i + 2; j < len(numeri) + 1; j++ {
			sequenzaString := strings.Join(numeri[i:j], " ")
			sottoSequenze[sequenzaString]++
		}
	}

	var sequenzeValide []string
	for k, _ := range sottoSequenze {
		sequenzaSlice := strings.Split(k, " ")

		if sequenzaSlice[0] == sequenzaSlice[len(sequenzaSlice) - 1] {
			sequenzeValide = append(sequenzeValide, strings.Join(sequenzaSlice, " "))
		}
	}

	sort.Slice(sequenzeValide, func (i, j int) bool { return len(sequenzeValide[i]) < len(sequenzeValide[j]) })

	for _, v := range sequenzeValide {
		Println(v)
	}
}
