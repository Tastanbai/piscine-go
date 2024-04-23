package main

import "github.com/01-edu/z01"

func main() {
	for h := 'z'; h >= 'a'; h-- {
		z01.PrintRune(h)
	}
	z01.PrintRune('\n')
}
