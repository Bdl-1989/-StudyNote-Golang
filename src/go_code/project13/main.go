package main

import "fmt"

func sum(n int, args ...int) int {
	sum := n
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	fmt.Println(n)
	return sum
}
func main() {

}
