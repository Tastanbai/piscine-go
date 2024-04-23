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
func main() {
	if len(os.Args) == 2 {
		res := ""
		for _, v := range os.Args[1] {
			if v >= 'a' && v <= 'z' {
				for i := 0; i < int(v-'a'+1); i++ {
					res += string(v)
				}
			}
			if v >= 'A' && v <= 'Z' {
				for i := 0; i < int(v-'A'+1); i++ {
					res += string(v)
				}
			}
		}
		print(res)
	}
}
