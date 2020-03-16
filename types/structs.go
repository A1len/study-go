package main

import "fmt"


//结构体，和java的自定义对象差不多

//todo 暂时先写这么一个，等后续看看别人项目以及自己项目的时候再回顾

type Wall struct {
	x int
	y int

}

func Setx(v *Wall) int{
		return v.x
}

func Getx(v *Wall) int{
	return v.x
}

func main() {
	v:=Vertex{1, 2}
	fmt.Print(v)
}
