package main

import (
	. "fmt"
	"os"
)

func main() {
	caratteri := os.Args[1:]

	var occorrenzeCount map[string]int = make(map[string]int)
	for i := 0; i < len(caratteri) - 3; i++ {
		for j := i + 3; j <= len(caratteri); j++ {
			occorrenza := caratteri[i:j]
			if occorrenza[0] == occorrenza[len(occorrenza) - 1] {
				var stringa string
				for _, e := range occorrenza {
					stringa += e + " "
				}
				occorrenzeCount[stringa[:len(stringa) - 1]]++
			}
		}
	}

	for chiave, valore := range occorrenzeCount {
		Print(chiave, " -> occorrenze: ", valore, "\n")
	}
}
