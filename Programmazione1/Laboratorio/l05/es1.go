/*
Scrivere un programma che legga da standard input una sequenza di numeri
interi compresi tra 1 e 9 (estremi inclusi) e stampi per ognuno di essi la tabellina
corrispondente. Il programma si interrompe quando viene inserito in input
un numero non valido (inferiore a 1 o superiore a 9) stampando Programma
terminato..

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti
funzioni: * Tabellina(numero int) bool che riceve in input un valore intero
nel parametro numero. Se numero è compreso tra 1 e 9 (estremi inclusi), la fun-
zione stampa la tabellina associata e restituisce un valore logico true. Altrimenti,
la funzione non stampa nulla e restituisce un valore logico false
*/

package main

import "fmt"

func tabellina(numero int) bool {
	if ! (numero >= 1 && numero <= 9) {
		return false
	}

	for i := 1; i <= 10; i++ {
		fmt.Print(numero * i, " ")
	}
	fmt.Println()

	return true
}

func main() {
	var n int = 1

	for n >= 1 && n <= 9 {
		fmt.Scan(&n)

		fmt.Print("Tabellina del ", n, ": ")
		tabellina(n)
	}
}
