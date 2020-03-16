package main

import (
	"math/cmplx"
	"fmt"
)

/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

//一般需要用多大的可以选择多大的数字类型
//如果没有选择将会按照操作系统选择
//比如64位就会走int64


//go语言中的基础数据类型都有默认值，
//这很正常引用指针得指向一个内存空间，才算初始化了一个变量
//否则一个指针不初始化，很容易造成基础数据类型空指针。。。有点尴尬

byte // uint8 的别名

rune // int32 的别名,因为unicode是动态的所以选择32位可以完全满足，
		新版本java对于这一块做了自动判断，节省空间


float32 float64

complex64 complex128
 */
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}