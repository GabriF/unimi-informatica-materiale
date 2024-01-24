package main

import (
	. "fmt"
	"os"
	"strconv"
)

func LeggiNumeri(sl []string) (numeri []int) {
	for _, c := range sl {
		intero, err := strconv.Atoi(c)

		if err == nil {
			numeri = append(numeri, intero)
		}
	}

	return
}

func Media (sl []int) float64 {
	if len(sl) == 0 {
		return 0
	}

	var s int

	for _, n := range sl {
		s += n
	}

	return float64(s) / float64(len(sl))
}

func Massimo(sl []int) int {
	if len(sl) == 0 {
		return 0
	}

	max := sl[0]

	for _, n := range sl[1:] {
		if n > max {
			max = n
		}
	}

	return max
}

func Minimo(sl []int) int {
	if len(sl) == 0 {
		return 0
	}

	min := sl[0]

	for _, n := range sl[1:] {
		if n < min {
			min = n
		}
	}

	return min
}

func main() {
	args := os.Args[1:]
	numeri := LeggiNumeri(args)

	Printf("Media: %.2f\n", Media(numeri))
	Println("Massimo: ", Massimo(numeri))
	Println("Minimo: ", Minimo(numeri))
}
