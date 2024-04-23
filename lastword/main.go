package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		word := ""
		lastword := ""
		hasword := false
		for _, v := range os.Args[1] {
			if v != ' ' {
				word += string(v)
			} else {
				if word != "" {
					lastword = word
					hasword = true
				}
				word = ""
			}
		}
		if word != "" {
			lastword = word
			hasword = true
		}
		if hasword {
			for _, v := range lastword {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		}
	}
}
