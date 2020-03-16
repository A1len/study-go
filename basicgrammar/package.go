package main

import (
	"fmt"
	"math/rand"
)

//每个 Go 程序都是由包组成的。
//程序运行的入口是包 `main`。
//这个程序使用并导入了包 "fmt" 和 `"math/rand"`。



func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}