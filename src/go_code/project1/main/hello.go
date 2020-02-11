package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

//
func main() {
	fmt.Println("hello world")
	var i int8 = 127
	fmt.Println(i)

	var n = 100
	var n2 = 10
	fmt.Printf("%T,%d \n", n, unsafe.Sizeof(n2)) //用于做格式化输出,以及返回某个变量的占用字节bit大小
	// %v 表示按照变量的值输出

	//用fmt.Sprintf转string
	var in int = 8
	var str string
	str = fmt.Sprintf("%d", in)
	fmt.Printf("%T  %q \n", str, str)

	//strconv
	//str = strconv.FormatInt(int64(in),10)
	str = strconv.Itoa(in)
	fmt.Printf("%T  %q \n", str, str)

	var in3 int64
	var str2 string = "99"
	in3, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("%T  %d \n", in3, in3)

	var str3 string = "123.423"
	var flo float64
	flo, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("flo type %T, flo = %v", flo, flo)

}
