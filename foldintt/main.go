package main

import "fmt"

func FoldInt(f func(int, int) int, a []int, n int) {
	if len(a) == 0 {
		return
	}

	res := f(n, a[0])
	for i := 1; i < len(a); i++ {
		res = f(res, a[i])
	}
	fmt.Println(res)
}
func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
func Add(n int, n1 int) int {
	return n + n1
}
func Mul(n int, n1 int) int {
	return n * n1
}
func Sub(n int, n1 int) int {
	return n - n1
}
