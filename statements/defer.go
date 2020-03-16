package main

import (
	"fmt"
)

//defer的操作其实是一个压栈的方式调用
//先进后出FILO
//可以一次执行下面的三个例子

//另外通过例子3其实可以知道defer是有数据拷贝的
//换句话说执行defer的时候数据就保存了一个份
//后续变量怎么改变defer还是输出同样的信息

//TODO defer的使用场景后续补充，暂时先当try catch finally使用


func main() {
	fmt.Println("counting")

	a:="hello world"

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	//1.os.Exit(0)//直接推出不会执行defer
	//2.return//遇到return才返回
	defer fmt.Println(a)
	//3.a="hello defer"
	fmt.Println("done")//没有return则执行到底
}
