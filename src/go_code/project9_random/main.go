package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// //为了生产随机数，还需要给rand生成一个种子
	// rand.Seed(time.Now().Unix())
	// //1~100

	// n := rand.Intn(100) + 1 // [0,100)+1
	// fmt.Println(n)
	rand.Seed(time.Now().UnixNano())
	var count int = 0
	for {
		//为了生产随机数，还需要给rand生成一个种子

		//1~100

		n := rand.Intn(100) + 1 // [0,100)+1
		fmt.Println(n)
		count++
		if n == 99 {
			fmt.Printf("%v  %v", n, count)
			break
		}

	}

}
