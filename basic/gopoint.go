package basic

import "fmt"

//存储地址的变量，就是指针  所以说指针也是变量 通过&vars 获取变量的指针
//释放了的变量对应的指针，就是野指针
/*
变量         值             地址
var1        100           0x000003
poiterVar1  0x000003      0x0000aa
*/
//指针也是个变量，它也有自己的内存地址

//指定指针的类型具有重要意义

//指针的运算（加减） 和多级运算

func PointerTest() {
	var intVar = 100
	fmt.Printf("value=%d,address =%v \n", intVar, &intVar)
	var pointerVar *int = &intVar
	fmt.Printf("value=%v,address =%v,var value =%v \n", pointerVar, &pointerVar, *pointerVar)
}
