package main

import "fmt"

func main() {
	var str string = "adksf!背景"
	str2 := []rune(str) //slice

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	//for range能自动读出汉字，但是index会占用，用了rune后，index就不会跳开了
	for index, val := range str2 {
		fmt.Printf("%d %c \n", index, val)
	}
	for index, val := range str {
		fmt.Printf("%d %c \n", index, val)
	}

}
