package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci la quantità di numeri che vuoi inserire: ")
	fmt.Scan(&n)

	var s, min, max, pos, neg, nul int = 0, 0, 0, 0, 0, 0
	var ni int
	for i := 0; i < n; i++ {
		fmt.Print(i, " - inserisci un intero: ")
		fmt.Scan(&ni)

		if i == 0 {
			max = ni
			min = ni
		}

		if ni > max {
			max = ni
		} else if ni < min {
			min = ni
		}

		if ni > 0 {
			pos++
		} else if ni < 0 {
			neg++
		} else { // ni == 0
			nul++
		}

		s += ni
	}

	fmt.Println("Max: ", max)
	fmt.Println("Min: ", min)
	fmt.Println("Positivi: ", pos)
	fmt.Println("Negativi: ", neg)
	fmt.Println("Nulli: ", nul)
	fmt.Println("Somma: ", s)
}
