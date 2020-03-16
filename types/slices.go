package main

import "fmt"

//https://blog.golang.org/go-slices-usage-and-internals slices看这篇博客就可以了
func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])
	//添加元素
	p=append(p,2)
	fmt.Println(p)

	//通过range遍历数据
	for i, v := range p {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
