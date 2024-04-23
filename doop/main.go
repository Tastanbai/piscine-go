package main

import (
	"os"
)

func isvalid(v1, v2 string) bool {
	if v1 == "-9223372036854775809" || v2 == "-9223372036854775809" {
		return false
	}
	for _, v := range v1 {
		if !(v >= '0' && v <= '9') && v != '-' {
			return false
		}
	}
	for _, v := range v2 {
		if !(v >= '0' && v <= '9') && v != '-' {
			return false
		}
	}
	return true
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
	if len(os.Args) == 4 {
		if isvalid(os.Args[1], os.Args[3]) {
			a := Atoi(os.Args[1])
			operator := os.Args[2]
			b := Atoi(os.Args[3])
			switch operator {
			case "+":
				res := a + b
				if res <= a {
					os.Stdout.WriteString(Itoa(res) + "\n")
				}
			case "-":
				res := a - b
				if res >= a {
					os.Stdout.WriteString(Itoa(res) + "\n")
				}
			case "/":
				if b == 0 {
					os.Stdout.WriteString("No division by 0\n")
				} else {
					res := a / b
					os.Stdout.WriteString(Itoa(res) + "\n")
				}
			case "*":
				res := a * b
				if res/a == b {
					os.Stdout.WriteString(Itoa(res) + "\n")
				}
			case "%":
				if b == 0 {
					os.Stdout.WriteString("No modulo by 0\n")
				} else {
					res := a * b
					os.Stdout.WriteString(Itoa(res) + "\n")
				}
			}
		}
	}
}
