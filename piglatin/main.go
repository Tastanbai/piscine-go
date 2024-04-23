package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		res := ""
		hasVowel := false // Flag to check if any vowel is found

		for i, w := range os.Args[1] {
			if w == 'a' || w == 'e' || w == 'i' || w == 'o' || w == 'u' || w == 'A' || w == 'E' || w == 'I' || w == 'O' || w == 'U' {
				res = os.Args[1][i:] + os.Args[1][:i] + "ay"
				hasVowel = true // Set the flag to true
				break
			}
		}

		if !hasVowel {
			res = "No vowels"
		}

		for _, w := range res {
			z01.PrintRune(w)
		}
		z01.PrintRune(10)
	}
}
