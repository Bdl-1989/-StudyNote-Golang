package main

import "fmt"

func swap(n *int, m *int) {
	t := *n
	*n = *m
	*m = t
}

func main() {
	a := 10
	b := 13
	swap(&a, &b)
	fmt.Println(a, b)

}
