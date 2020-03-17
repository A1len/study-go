package outer



//隐式加入linkname包的调用，并且添加一个i.s文件否则无法编译通过
import (
	_ "studygo/basicgrammar/linkname"
)

//后续补充go:linkname
func FlushICache()



