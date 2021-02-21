package basic

import "fmt"

//数据是长度固定的相同类型的集合
//数组的内存地址是连续的
//数组定义但未赋值，也有默认值
//数组的三种定义方式
//go 语言中，数组是值类型

func ArrayTest() {
	arrayVar := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arrayVar)
	nochange(arrayVar)
	fmt.Println(arrayVar)
	change(&arrayVar)
	fmt.Println(arrayVar)
}

func nochange(arr [5]int) {
	arr[0] = 100
}

func change(arr *[5]int) {
	(*arr)[2] = 100
}
