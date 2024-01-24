package main

import (
	. "fmt"
	"os"
	"strconv"
	_ "sort"
)

func EPrimo(n int) bool {
	for i := 2; i < n / 2; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func Occorrenze(numeri []int, n int) (occorrenze int) {
	for _, e := range numeri {
		if e == n {
			occorrenze++
		}
	}

	return
}

func main() {
	var primi []int

	if len(os.Args) < 2 {
		Println("Devi inserire un numero intero")
		return
	}

	numeroStringa := os.Args[1]

	if _, err := strconv.Atoi(numeroStringa); err != nil {
		Println("Devi inserire un intero")
		return
	}

	for i := 1; i <= 3; i++ {
		for j := 0; j < len(numeroStringa) - i + 1; j++ {
			Println(numeroStringa[0:j] + numeroStringa[i + j:len(numeroStringa)])
		}
	}


}
