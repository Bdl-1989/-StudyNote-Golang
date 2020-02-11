package main

import "fmt"

func main() {
	var i int = 9
	fmt.Println("i's address", &i)

	var ptr *int
	ptr = &i
	*ptr = 10
	fmt.Println("i's value", i)
	var abc_id int = 9
	fmt.Println(abc_id)
}
