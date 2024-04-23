package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		res := ""
		num, _ := strconv.Atoi(os.Args[1])
		if num == 0 {
			res = "00000000"
		}
		for len(res) != 8 {
			res = string(rune(num%2+48)) + res
			num /= 10
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
