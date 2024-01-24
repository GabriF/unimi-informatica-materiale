package main

import "fmt"

func eTernaPitagorica(a, b, c int) bool {
	if c * c == a * a + b * b {
		return true
	}

	return false
}

func ternePitagoriche(soglia int) int {
	i := 0

	for c := 1; c < soglia; c++ {
		for a := 1; a < soglia; a++ {
			for b := a; b < soglia; b++ {
				if eTernaPitagorica(a, b, c) {
					fmt.Println("(", a, b, c, ") è una terna pitagorica")
					i++
				}
			}
		}
	}

	return i
}

func main() {
	fmt.Println(ternePitagoriche(100))
}
