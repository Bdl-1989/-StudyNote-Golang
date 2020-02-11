package main

import "fmt"

func main() {
	var days int = 97
	var weeks int = days / 7
	var rest int = days % 7
	fmt.Printf("%d对应有%d个星期以及%d天", days, weeks, rest)

	
}
