package main

import "fmt"

func StampaRigaInizioPiu(larghezza int) {
	for i := 0; i < larghezza; i++ {
		if i % 2 == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("*")
		}
	}
}

func StampaRigaInizioAsterisco(larghezza int) {
	for i := 0; i < larghezza; i++ {
		if i % 2 == 0 {
			fmt.Print("*")
		} else {
			fmt.Print("+")
		}
	}
}

func StampaScacchiera(dimensione int) {
	if dimensione <= 0 {
		return
	}

	for i := 0; i < dimensione; i++ {
		if i % 2 == 0 {
			StampaRigaInizioAsterisco(dimensione)
		} else {
			StampaRigaInizioPiu(dimensione)
		}
		fmt.Println()
	}
}

func main() {
	StampaScacchiera(10)
}
