package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, v := range os.Args[1] {
			if v >= 'a' && v < 'm' || v >= 'A' && v < 'M' {
				z01.PrintRune(v + 13)
			} else if v > 'm' && v <= 'z' || v > 'M' && v <= 'Z' {
				z01.PrintRune(v - 13)
			} else {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune('\n')
	}

}
