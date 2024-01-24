package main

import (
	. "fmt"
	"os"
)

func isBalanced(sequence string) bool {
	var lv int

	for i := 0; i < len(sequence); i++ {
		if sequence[i] == '(' {
			lv++
		} else if sequence[i] == ')' {
			lv--
		}

		// Troppe chiuse
		if lv < 0 {
			return false
		}
	}

	// Troppe aperte non chiuse
	if lv != 0 {
		return false
	}

	return true
}

func subBalanced(sequence string) (lista []string) {
	for i := 0; i < len(sequence); i++ {
		for j := i + 2; j < len(sequence) + 1; j++ {
			if isBalanced(sequence[i:j]) {
				lista = append(lista, sequence[i:j])
			}
		}
	}

	return
}

func main() {
	stringa := os.Args[1]

	if isBalanced(stringa) {
		Println("bilanciata")
	} else {
		Println("non bilanciata")
	}

	for i, v :=  range subBalanced(stringa) {
		Print(i + 1, ") ", v, "\n")
	}
}
