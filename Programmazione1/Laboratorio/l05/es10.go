package main

import "fmt"

func main() {
	MENU := "\t1 - patatine\n\t2 - hamburger\n\t3 - coca cola\n\t0 - termina"

	var inp, totale uint
	for  {
		fmt.Println("Cosa vuoi ordinare?")
		fmt.Println(MENU)

		fmt.Scan(&inp)

		switch inp {
			case 1, 3:
				totale += 2
			case 2:
				totale += 5
		}

		if inp == 0 {
			break
		}
	}

	fmt.Println("Il totale è", totale + 2)
}
