package main

import (
	"os"
)

func main() {
	if len(os.Args) == 3 {
		a := os.Args[1]
		b := os.Args[2]
		k := 0
		res := ""
		for i := 0; i < len(a); i++ {
			for j := k; j < len(b); j++ {
				if a[i] == b[j] {
					res += string(b[j])
					k = j
					break
				}
			}
		}
		if res == a {
			os.Stdout.Write([]byte(a + "\n"))
		}
	}
}
