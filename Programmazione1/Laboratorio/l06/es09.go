package main

import "fmt"
import "strconv"
import "math"

type Frazione struct {
	n, d int
}

func Riduci(f *Frazione) {
	if f.n / f.d == 1.0 {
		f.n, f.d = 1, 1
	} else if f.n % f.d == 0 {
		f.n, f.d = f.n / f.d, 1
	} else {
		for i := 2; i < int(math.Min(float64(f.n), float64(f.d))); i++ {
			if f.n % i == 0 && f.d % i == 0 {
				f.n, f.d = f.n / i, f.d / i
				Riduci(f)
			}
		}
	}
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
	Riduci(f)
	fmt.Println(String(*f))
}
