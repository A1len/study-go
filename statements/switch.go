package main

import (
	"fmt"
	"runtime"
	"time"
)


//switch用法比java更方便灵活，不用显示的在语句中加break
//如果要实现类似java的一个case成立继续执行后续的case可以使用fallthrough
//case也可以直接使用判断语句

//以及type switch

//todo 后续看一下switch的实现源码，感觉java的swtich语法糖和go比起来还是较弱，这也是java的历史原因
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		//fallthrough 打开注释看看结果
	case "linux":
		fmt.Println("Linux.")
		//fallthrough
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}


	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}


	var a interface{}

	a = "abc"



	switch t := a.(type) {

	case string:

		fmt.Printf("string %s\n", t)

	case int:

		fmt.Printf("int %d\n", t)

	default:

		fmt.Printf("unexpected type %T", t)

	}

}



