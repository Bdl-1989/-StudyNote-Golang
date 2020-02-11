package main

import "fmt"

//type switch
func main() {
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("%T", i)
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	case bool, string:
		fmt.Println("bool or string")
	default:
		fmt.Println("unknow")
	}
}
