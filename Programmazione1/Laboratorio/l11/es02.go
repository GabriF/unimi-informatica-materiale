package main

import (
	. "fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type RGB struct {
	rosso uint8
	verde uint8
	blu uint8
}

type Immagine struct {
	larghezza uint
	altezza uint

	sequenzaPixel []RGB
}

func main() {
	file, _ := os.Open("immagine.ppm")
	scan := bufio.NewScanner(file)

	immagine := new(Immagine)

	scan.Scan() // salta prima riga

	// lettura larghezza e altezza
	scan.Scan()
	dimensioni := scan.Text()
	larghezza, _ := strconv.Atoi(string(dimensioni[0]))
	altezza, _ := strconv.Atoi(string(dimensioni[2]))
	immagine.larghezza = uint(larghezza)
	immagine.altezza = uint(altezza)
	immagine.sequenzaPixel = make([]RGB, immagine.larghezza * immagine.altezza)
	Println(len(immagine.sequenzaPixel))
	scan.Scan() // salta terza riga

	var i int
	for scan.Scan() {
		riga := scan.Text()
		valori := strings.Split(riga, " ")
		rosso, _ := strconv.Atoi(string(valori[0]))
		verde, _ := strconv.Atoi(string(valori[1]))
		blu, _ := strconv.Atoi(string(valori[2]))

		colori := RGB{rosso:uint8(rosso), verde:uint8(verde), blu:uint8(blu)}

		immagine.sequenzaPixel[i] = colori

		i++
	}

	Println("Immagine letta:", *immagine)
}
