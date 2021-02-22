package main

import (
	"fmt"

	"xiet16.com/golearn/basic"
	"xiet16.com/golearn/pgs"
	"xiet16.com/golearn/webhttp"
)

func main() {
	fmt.Println("learn start")
	pgs.Excute()

	/*basic包 基础语法学习*/
	basic.Vartest()
	// basic.FloatTest()
	// basic.PointerTest()
	// basic.ArrayTest()
	// basic.StructTest()

	webhttp.HttpHandleTest()
}
