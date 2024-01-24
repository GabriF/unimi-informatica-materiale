package main

import "fmt"

func main() {
	var m, q, x, y float64

	fmt.Print("Inserisci i valori m e q della retta y = mx + q: ")
	fmt.Scan(&m, &q)

	fmt.Print("Ora inserisci i valori x e y di un punto: ")
	fmt.Scan(&x, &y)

	if y == m * x + q {
		fmt.Println("Il punto appartiene alla retta")
	} else if y > m * x + q {
		fmt.Println("Il punto sta sopra la retta")
	} else if y < m * x + q {
		fmt.Println("Il punto sta sotto la retta")
	}
}
