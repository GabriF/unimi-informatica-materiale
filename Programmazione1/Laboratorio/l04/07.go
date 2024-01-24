/*
n = 6

******
 *   *
  *  *
   * *
    **
     *
*/

package main

import . "fmt"

func main() {
	Print("Inserisci n: ")
	var n int
	Scan(&n)

	for i, col := 0, 1; i < n * 2; i++ {
		if i == n {
			col = 2
		}

		for j := 0; j < n * col - (col / 2); j++ {
			if i < n && (j == n - 1 || j == i || i == 0) {
				Print("*")
			} else if i >= n && (j == n - 1 || (j >= n && (i == n * 2 - 1 || j == i - 1)))  {
				Print("*")
			} else {
				Print(" ")
			}
		}

		Println()
	}
}
