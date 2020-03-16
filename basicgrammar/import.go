package main


import (
	"fmt"
	"math"
)
//有两种方式可以倒入包，上述第一种方式是打包语句方式,也可以通过下面方式
//import "fmt"
//import "math"

func main() {
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}
