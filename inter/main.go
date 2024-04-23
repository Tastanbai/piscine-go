package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		mapa := make(map[string]bool)
		for _, a := range os.Args[1] {
			for _, b := range os.Args[2] {
				if a == b {
					if _, value := mapa[string(a)]; !value {
						mapa[string(a)] = true
						z01.PrintRune(a)

					}
				}

			}
		}
		z01.PrintRune('\n')
	}

}
