package main

import (
	"fmt"

	"xiet16.com/golearn/basic"
	"xiet16.com/golearn/pgs"
)

func main() {
	fmt.Println("learn start")
	pgs.Excute()

	/*basic包 基础语法学习*/
	//basic.Vartest()
	// basic.FloatTest()
	// basic.PointerTest()
	// basic.ArrayTest()
	// basic.StructTest()
	// basic.DataSplitTest()
	// basic.ChainMiddlewareTest()
	//basic.Middleware_V1Test()
	basic.Middleware_V3Test()
	//webhttp.HttpHandleTest()
}
