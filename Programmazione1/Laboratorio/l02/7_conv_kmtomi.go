package main

import "fmt"

func main() {
	var km float64

	fmt.Print("Distanza (km) = ")
	fmt.Scan(&km)

	fmt.Println("Distanza (mi) = ", km * 0.62137)
}
