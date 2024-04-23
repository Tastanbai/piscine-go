package main

import "fmt"

func StrRev(s string) string {
	newstr := ""
	for i := len(s) - 1; i >= 0; i-- {
		newstr += string(s[i])
	}
	return newstr
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
