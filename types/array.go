package main

import "fmt"


//新建一个数组，和java差不多
func main() {
	var a [2]string
	b:=[]int{1,2,4,5}
	fmt.Print(b)
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}