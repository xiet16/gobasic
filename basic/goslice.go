package basic

import "fmt"

// 切片又称动态数组，长度是可变的
//切片由指针，长度，容量组成
//切片是引用类型，数组是值类型
//append 函数
//切片有自己的地址
//切片和数组的定义一样， var sliceVar []int,切片值得地址和数组值的地址是一致的，因为它指向的是底层的数组
//切片作为形参，实参与形参之间，指针指向同一个底层数组，但他们有两个不同的内存地址。形参在扩容时，如果超出了容量，会产生新的数组，这是形参和实参指向的地址就会不同

func SliceTest() {
	//定义一个容量和长度都为5的切片
	slice := make([]int, 5)
	fmt.Printf("slice pointer addr :%p \n", &slice)       //指针的地址
	fmt.Printf("slice pointer before addr :%p \n", slice) //底层数组的地址
	fmt.Printf("slice change before data : %v \n", slice)
	changedata(slice)
	fmt.Printf("slice pointer before addr :%p \n", slice)
	fmt.Printf("slice change after data: %v \n", slice)
}

func changedata(s []int) {
	fmt.Printf("s change pointer addr : %p \n", &s) //指针的地址
	fmt.Printf("s change before addr : %p \n", s)   //底层数组的地址
	fmt.Printf("s change before  data: %v \n", s)
	s[0] = 66
	fmt.Printf("s pointer after addr :%p \n", s)
	fmt.Printf("s change after data: %v \n", s)
}

func SliceTest2() {
	slice := make([]int, 5)
	fmt.Printf("slice pointer addr :%p \n", &slice)       //指针的地址
	fmt.Printf("slice pointer before addr :%p \n", slice) //底层数组的地址
	fmt.Printf("slice change before data : %v \n", slice)
	changecapicity(slice)
	fmt.Printf("slice pointer before addr :%p \n", slice)
	fmt.Printf("slice change after data: %v \n", slice)
}

//改变容量,改变后，形参指向底层的地址变了
func changecapicity(s []int) {
	fmt.Printf("s change pointer addr : %p \n", &s) //指针的地址
	fmt.Printf("s change before addr : %p \n", s)   //底层数组的地址
	fmt.Printf("s change before  data: %v \n", s)
	s = append(s, 66)
	fmt.Printf("s pointer after addr :%p \n", s)
	fmt.Printf("s change after data: %v \n", s)
}
