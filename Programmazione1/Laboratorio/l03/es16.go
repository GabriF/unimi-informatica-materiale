package main

import "fmt"

func main() {
	fmt.Println("1) secondi -> ore")
	fmt.Println("2) secondi -> minuti")
	fmt.Println("3) minuti -> ore")
	fmt.Println("4) minuti -> secondi")
	fmt.Println("5) ore -> secondi")
	fmt.Println("6) ore -> minuti")
	fmt.Println("7) minuti -> giorni e ore")
	fmt.Println("8) minuti -> anni e giorni")

	var s int
	fmt.Print("Cosa vuoi fare?\n>")
	fmt.Scan(&s)

	var v float64

	if s >= 1 && s <= 8 {
		fmt.Print("Inserisci il valore su cui effettuare l'operazione: ")
		fmt.Scan(&v)

		if s == 1 {
			fmt.Println(v, "secondi corrispondono a ", v / 60 / 60, "ora/e")
		} else if s == 2 {
			fmt.Println(v, "secondi corrispondono a ", v / 60, "minuto/i")
		} else if s == 3 {
			fmt.Println(v , "minuti corrispondono a", v / 60, "ora/e")
		} else if s == 4 {
			fmt.Println(v, "minuti corrispondono a", v * 60, "secondo/i")
		} else if s == 5 {
			fmt.Println(v, "ore corrispondono a", v * 60 * 60, "secondo/i")
		} else if s == 6 {
			fmt.Println(v, "ore corrispondono a", v * 60, "minuto/i")
		} else if s == 7 {
			ore := v / 60
			giorni := ore / 24

			fmt.Println(v, "minuti corrispondono a ", ore, "ora/e e", giorni, "giorno/i")
		} else if s == 8 {
			giorni:= v / 60 / 24
			anni := giorni / 365

			fmt.Println(v, "minuti corrispondono a", giorni, "giorno/i e", anni, "anno/i")
		}
	} else {
		fmt.Println("Devi scegliere un valore tra 1 e 8")
	}
}
