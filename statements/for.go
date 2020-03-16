package main

import "fmt"


//for的使用和java中基本是一样的，只是少了一个（）
//并且go当中没有while，使用for实现
//java中Doug Lea很喜欢这种写法
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	i := 1
	for i < 1000 {
		i += i
	}
	fmt.Println(i)

	for {
		break
	}
}
