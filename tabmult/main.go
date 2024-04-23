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

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	hasword := false
	if n <= 0 {
		n = -n
		hasword = true
	}
	result := ""
	for n > 0 {
		digit := n % 10
		result += string(digit + '0')
		n /= 10
	}
	res := ""
	for i := len(result) - 1; i >= 0; i-- {
		res += string(result[i])
	}
	if hasword {
		res = "-" + res
		return res
	} else {
		return res
	}
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
		args := os.Args[1]
		n := Atoi(args)

		for i := 1; i <= 9; i++ {
			res := ""

			res += Itoa(i)
			res += " x "
			res += args
			res += " = "
			res += Itoa(i * n)
			print(res)
		}

	}
}
