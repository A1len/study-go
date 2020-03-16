package main

import ("fmt" "unsafe")


//函数可以返回一个或者多个值，具体看add和multipl函数
//两个或多个连续的函数命名参数是同一类型，可以只需要最后一个参数定义类型
//Go 的返回值可以被命名，并且像变量那样使用。
//返回值的名称应当具有一定的意义，可以作为文档使用
//没有参数的 return 语句返回结果的当前值。也就是`直接`返回
func add(x int, y int)(int){
	return x+y
}

func multiple(x , y int)(string ,int){
	return "hello",x
}


func multiple2(x , y int)(str string ,z int){
	str="hello"
	z=x
	return
}

//外部实现,这里其实可以不定义结构体
//后续补充go:linkname
func FlushICache(begin,end int)



func main() {
	x ,y :=multiple(42, 13)
	fmt.Print(x,y)

}
