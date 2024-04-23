package main

import (
	"fmt"
	"strconv"
)

func ReduceInt(a []int, f func(int, int) int) {
	res := f(a[0], a[1])

	for i := 2; i < len(a); i++ {
		res = f(res, a[1])
	}
	num := strconv.Itoa(res)
	fmt.Println(num)
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
