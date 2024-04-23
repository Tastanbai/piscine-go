package main

import "fmt"

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := 0
	for _, min := range a {
		if min > max {
			max = min
		}
	}
	return max
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
