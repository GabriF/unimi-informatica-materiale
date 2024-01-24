package main

import (
	. "fmt"
	"os"
	"strconv"
)

func Fattoriali(n int) (f []int) {
	f = append(f, 1)

	for i := 2; i <= n; i++ {
		f = append(f, f[i - 2] * i)
	}

	return
}

func main() {
	args := os.Args[1:]

	numero, err := strconv.Atoi(args[0])

	if err == nil && numero >= 1 {
		Println(Fattoriali(numero))
	} else {
		Println("Inserisci un numero intero maggiore o uguale a 1")
	}


}
