package main

import (
	"os"

	"github.com/01-edu/z01"
)

func print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func Atoi(s string) int {
	sing := 1
	num := 0

	for i, v := range s {
		if i == 0 && (v == '-' || v == '+') {
			if v == '-' {
				sing = -1
			}
			continue
		}
		if v < '0' || v > '9' {
			return 0
		}
		num = num*10 + int(v-'0')
	}
	return num * sing
}

func main() {
	if len(os.Args) == 2 {
		k := 1
		n := Atoi(os.Args[1])
		state := false
		for n >= k {
			if n == k {
				print("true")
				state = true
				break
			}
			k *= 2
		}
		if state == false {
			print("false")
		}
	}
}
