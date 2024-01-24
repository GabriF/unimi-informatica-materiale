package main

import (
	. "fmt"
	"math/rand"
	"errors"
)

type Mazzo struct {
	carte []CartaDaGioco
	ncarte int
}

type CartaDaGioco struct {
	seme, simbolo string
}


func CreaCarta(seme, simbolo string) CartaDaGioco {
	return CartaDaGioco{seme:seme, simbolo:simbolo}
}

func CreaMazzo() Mazzo {
	const NCARTE = 40
	SEMI := [...]string{"quadri", "picche", "cuori", "fiori"}
	SIMBOLI := [...]string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette", "fante","donna", "re"}

	var mazzo Mazzo
	mazzo.ncarte = NCARTE
	mazzo.carte = make([]CartaDaGioco, mazzo.ncarte)

	for i, n := 0, 0; i < len(SEMI); i++ {
		for j := 0; j < len(SIMBOLI) && n < mazzo.ncarte; j++ {
			mazzo.carte[n] = CreaCarta(SEMI[i], SIMBOLI[j])
			n++
		}
	}

	return mazzo
}

func Mischia(m Mazzo) {
	 for i, _ := range m.carte {
		j := rand.Intn(i + 1)
		m.carte[i], m.carte[j] = m.carte[j], m.carte[i]
	}
}

func Preleva(m Mazzo) (CartaDaGioco, Mazzo, error) {
	if len(m.carte) == 0 {
		return CartaDaGioco{}, Mazzo{ncarte:m.ncarte}, errors.New("Mazzo vuoto")
	}

	c := m.carte[len(m.carte) - 1]
	m.carte = m.carte[:len(m.carte) - 1]
	mazzo := m

	return c, mazzo, nil
}

func Riponi(m Mazzo, c CartaDaGioco) (Mazzo, error) {
	if len(m.carte) == m.ncarte {
		return m, errors.New("Mazzo pieno")
	}

	m.carte = append(m.carte, c)
	mazzo := m

	return mazzo, nil
}

func main() {
	var m Mazzo
	var c CartaDaGioco
	var err error

	m = CreaMazzo()
	Println("Mazzo:", m)

	Mischia(m)
	Println("\nMazzo mischiato:", m, "\n")

	// Manda in errore
	for i := 0; i < 40; i++ {
		c, m, err = Preleva(m)
		Println("\nPreleva mazzo:", c, m, err)
	}

	// Manda in errore
	for i := 0; i < 41; i++ {
		m, err = Riponi(m, c)
		Println("\nRimetti nel mazzo:", m, err)
	}
}
