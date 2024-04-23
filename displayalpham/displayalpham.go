package main

import "github.com/01-edu/z01"

func main() {
	a := "aBcDeFgHiJkLmNoPqRsTuVwXyZ"
	for _, s := range a {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
