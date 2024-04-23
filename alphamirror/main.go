package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, v := range os.Args[1] {
			if v >= 'a' && v <= 'z' {
				z01.PrintRune('a' - v + 'z')
			} else if v >= 'A' && v <= 'Z' {
				z01.PrintRune('A' - v + 'Z')
			} else {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune('\n')
	}
}
