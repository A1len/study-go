package main

import "fmt"


//go当中声明一个变量可以通过var或者:=
//var可以作用在函数中也可以函数外
//:=只可以在函数中
//go有自动类型推断，如果有初始值，可以不带上变量类型


var initname="allen"


var name ,motherland  string

var defaultvar string

func main()  {
	var num int
	age:=18


	fmt.Println(defaultvar,num,age)

}
