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

/*
GO 语言中，指针的限制:
1、不能计算
2、不同类型的指针不能赋值
3、不同类型的指针不能相互转换
4、不同类型的指针不能使用 == 或 != 比较
*/

func CountLimit() {
	//a := 100
	//p := &a
	//p++
	//p = &a + 3

	// var fp *float32
	// f := float64(66.6)
}
