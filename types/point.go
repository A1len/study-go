package main

import "fmt"

//go里面的指针，个人暂时理解是减少数据拷贝
//众所周知java中全是值传递

//在下面的例子中setname函数和setpoint函数很好的说明了传递方式
//todo 数据的传递一般还是用指针吧。除非不想数据被修改，比如拷贝副本数据扩容等级制，不过说回来cow效率更高

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	fmt.Println(i)
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	man:=person{"allen"}
	fmt.Println(man)
	setname(man)
	fmt.Println(man)
	setnamebypoint(&man)
	fmt.Println(man)

}

func setname(p person) {
	  p.name="AE"
}


func setnamebypoint(p *person) {
	p.name="AE2"
}


type person struct {
	name string
}
