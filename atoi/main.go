package main

import "fmt"

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
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
