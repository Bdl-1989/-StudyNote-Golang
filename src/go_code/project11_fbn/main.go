package main

import "fmt"

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	//fmt.Println(fbn(n - 1))
	return fbn(n-2) + fbn(n-1)

}
func main() {
	//fmt.Println(fbn(55))
	var n int = 1222
	fmt.Println(fbn(4), n)
}
