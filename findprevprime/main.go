package main

import "fmt"

func isprime(nb int) bool {
	for i := 0; i*i <= nb; i++ {
		if nb%2 == 0 {
			return false
		}

	}
	return true
}

func FindPrevPrime(nb int) int {
	for nb > 2 {
		if isprime(nb) {
			return nb
		}
		nb--
	}
	return nb

}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
