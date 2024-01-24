package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Errore input")
		return
	}

	var s string
	_, err = fmt.Scan(&s)

	if err != nil || s == "" {
		fmt.Println("Errore input")
		return
	} else {
		for _, c := range s {
			if c == ' ' {
				fmt.Println("Errore input")
				return
			}
		}
	}

	for i := 0; i < n; i++ {
		if i != n - 1 {
			fmt.Print(s, "-")
		} else {
			fmt.Println(s)
		}
	}
}
