package main

import (
	"fmt"
	"strings"
)

// type abc func(string, string) string

// //
// func makeSuffix() abc {
// 	return func(s string, suffix string) string {
// 		return s + suffix
// 	}
// }

// func main() {
// 	f := makeSuffix
// 	suffix := ".jpg"
// 	filename := "abc.igk"
// 	if strings.HasSuffix(filename, suffix) {
// 		fmt.Println(filename)
// 	} else {
// 		fmt.Println(f(filename, suffix))

// 	}

// }

type abc func(string) string

//这里输入应该是公用部分
func makeSuffix(suffix string) abc {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix

	}
}

func main() {
	suffix := ".jpg"
	name := "sdfaldjf.gje"
	f := makeSuffix(suffix)
	fmt.Println(f(name))
}
