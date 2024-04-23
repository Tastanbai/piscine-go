package main

import "fmt"

func Capitalize(s string) string {
	word := true
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if word {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] -= 'a' - 'A'
			}
			word = false
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 'z' - 'Z'
		}
		if !(runes[i] >= 'a' && runes[i] <= 'z') && !(runes[i] >= 'A' && runes[i] <= 'Z') && !(runes[i] >= '0' && runes[i] <= '9') {
			word = true
		}
	}
	return string(runes)
}
func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
