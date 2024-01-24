package main

import (
	. "fmt"
	"os"
	"strconv"
)

func CreaTavolaPitagorica(n int) [][]int {
	t := make([][]int, 0)

	for i := 1; i <= n; i++ {
		l := make([]int, 0)
		for j := 1; j <= 10; j++ {
			l = append(l, i * j)
		}
		t = append(t, l)
	}

	return t
}

func StampaTavolaPitagorica(s [][]int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			Printf("%4d", s[i][j])
		}
		Println()
	}
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		return
	}

	t := CreaTavolaPitagorica(n)
	StampaTavolaPitagorica(t)
}
