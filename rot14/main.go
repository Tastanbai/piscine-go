package main

import "github.com/01-edu/z01"

func Rot14(s string) string {
	runes := []rune(s)
	str := ""
	for _, v := range runes {
		if v >= 'a' && v < 'm' || v >= 'A' && v < 'M' {
			str += string(v + 14)
		} else if v > 'm' && v <= 'z' || v > 'M' && v <= 'Z' {
			str += string(v - 12)
		} else {
			str += string(v)
		}
	}
	return str
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
