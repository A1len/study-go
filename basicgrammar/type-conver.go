package main

import (
	"fmt"
)
//go里面其实和java差不多需要显示调用强转类型
func main() {
	var x, y int = 3, 4
	var f  = float64(x)
	var z  = int(f)
	fmt.Println(x, y, z)
}
