package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := len(os.Args[1:])
	z01.PrintRune(rune('0' + args))
	z01.PrintRune('\n')
}
