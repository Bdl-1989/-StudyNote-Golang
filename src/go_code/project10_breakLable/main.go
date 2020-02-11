package main

import "fmt"

func main() {
label1:
	for i := 0; i < 10; i++ {
		//label2:
		for j := 0; j < 20; j++ {
			if j == 14 {
				//break
				break label1
				//break label2
			}
			fmt.Println(i, j)
		}

	}

}
