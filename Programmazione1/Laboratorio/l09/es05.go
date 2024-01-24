package main

import (
	. "fmt"
	"strconv"
	"os"
)

func GeneraTriangolo(n int) [][]int {
	if n < 1 {
		return nil
	}

	triangolo := make([][]int, n)

	triangolo[0] = []int{1}

	for i := 1; i < n; i++ {
		riga := make([]int, i + 1)

		riga[0] = 1
		riga[len(riga) - 1] = 1

		for j := 1; j < len(riga) - 1; j++ {
			s1 := triangolo[i - 1][j - 1]
			s2 := triangolo[i - 1][j]
			riga[j] = s1 + s2
		}

		triangolo[i] = riga
	}

	return triangolo
}

func main() {
	n, ok := strconv.Atoi(os.Args[1])

	if ok != nil {
		Println("Devi mettere un intero")
		return
	}

	triangolo := GeneraTriangolo(n)

	for i := 0; i < len(triangolo); i++ {
		for j := 0; j < len(triangolo[i]); j++ {
			Print(triangolo[i][j], " ")
		}
		Println()
	}
}
