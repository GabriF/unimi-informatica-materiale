package main

import "fmt"

func main() {
	var n1, n2, n3, n4 int

	n, err := fmt.Scan(&n1, &n2, &n3, &n4)
	fmt.Println(n, err)
	fmt.Println(n1, n2, n3, n4)
}
