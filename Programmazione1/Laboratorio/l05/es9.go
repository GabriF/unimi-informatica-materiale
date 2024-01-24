package main

import "fmt"

func main() {
	var s string

	fmt.Scan(&s)

	var vM, vm, cM, cm int

	for _, r := range s {
		switch r {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				if r >= 65 && r <= 90 {
					vM++
				} else {
					vm++;
				}
			default:
				if r >= 65 && r <= 90 {
					cM++
				} else {
					cm++;
				}
		}
	}

	fmt.Println("Vocali maiuscole:", vM)
	fmt.Println("Vocali minuscole:", vm)
	fmt.Println("Consonanti maiuscole:", cM)
	fmt.Println("Consonanti minuscole:", cm)
}
