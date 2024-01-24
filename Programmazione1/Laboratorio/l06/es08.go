package main

import "fmt"
import "strconv"

type Frazione struct {
	n, d int
}

func String(f Frazione) string {
	return strconv.Itoa(f.n) + "\\" + strconv.Itoa(f.d)
}

func NuovaFrazione(n, d int) *Frazione {
	f := Frazione{n, d}

	return &f
}

func main() {
	var n, d int

	fmt.Print("Inserisci il numeratore: ")
	fmt.Scan(&n)
	fmt.Print("Inserisci il denominatore: ")
	fmt.Scan(&d)

	if d == 0 {
		fmt.Println("Il denominatore non può essere nullo")
		return
	}

	f := NuovaFrazione(n, d)

	fmt.Println(String(*f))
}
