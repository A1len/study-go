package linkname

import "fmt"
import _ "unsafe"


//下面这一行不是注释，而是golinkname的标准用法，类似java注解的标识意思
//格式为： go:linkname localname linkname

//go:linkname sayhello studygo/basicgrammar/linkname/outer.FlushICache
func sayhello(){
	fmt.Print("flush db")

}
