package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		n := 0
		for _, v := range os.Args[1] {
			if v < '0' || v > '9' {
				fmt.Printf("ERROR: cannot convert to roman digit")
				return
			}
			n = n*10 + int(v-'0')
		}
		if n == 0 || n >= 4000 {
			fmt.Printf("ERROR: cannot convert to roman digit")
			return
		}
		num := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		romnum := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		mathnum := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
		var str, sum string // str := "" sum := ""
		for i := range num {
			temp := n / num[i]
			for temp > 0 {
				n -= num[i]
				str += romnum[i]
				sum += mathnum[i] + "+"
				temp--
			}
		}
		fmt.Printf("%s\n%s\n", sum[:len(sum)-1], str)
	}
}
