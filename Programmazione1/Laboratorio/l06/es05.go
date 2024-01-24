package main

import "fmt"
import "strconv"

func main() {
	var s string
	var som int

	for {
		fmt.Scan(&s)

		for _, c := range s {
			intero, err := strconv.Atoi(string(c))

			if err == nil {
				som += intero
			} else {
				fmt.Println(som)
				return
			}
		}
	}
}
