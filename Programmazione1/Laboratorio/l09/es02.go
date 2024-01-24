package main

import (
	. "fmt"
	"os"
)

func Palindroma(s []rune) bool {
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s) - i - 1] {
			return false
		}
	}

	return true
}

func main() {
	stringa := []rune(os.Args[1])

	var palindrome []string

	for i := 0; i < len(stringa); i++ {
		for j := i + 2; j <= len(stringa); j++ {
			if Palindroma(stringa[i:j]) {
				palindrome = append(palindrome, string(stringa[i:j]))
			}
		}
	}

	Println(palindrome)
}
