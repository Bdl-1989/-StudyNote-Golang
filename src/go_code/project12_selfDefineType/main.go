package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func mytype(funvar func(int, int) int, n1 int, n2 int) int {
	return funvar(n1, n2)
}

type abc func(int, int) int

func mytype2(funvar abc, n1 int, n2 int) int {
	return funvar(n1, n2)
}

func main() {
	res := mytype(getSum, 1, 4)
	fmt.Println(res)
}
