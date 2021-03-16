package basic

import (
	"fmt"
	"unsafe"
)

/*

 */

func UnsafeTest() {
	a := 6
	double(&a)
	println("double a is : ", a)
}

//go 语言中的形参都是值传递，这里x 和 a 有两个不同的地址
func double(x *int) {
	*x += *x
	x = nil // 这个nil 赋值，不影响a ,因为他们是有不一样地址的两个参数
}

/*
上面的指针是安全的，但是有很多限制。有些指针是不安全的,没那么多限制，比如unsafe 包里的 unsafe.Pointer
任何类型的指针都可以和unsafe.Pointer互转，uintptr可以和unsafe.Pointer 互转，但不可以和其他类型互转，并且 uintptr 对内存的是不保护的，也就是在使用uintptr时，可能内存已经被回收了
*T <--> unsafe.Pointer <--> uintptr
unsafe.Pointer 的定义：
var ArbitraryType int
var Pointer *ArbitraryType

相关url:
https://www.cnblogs.com/qcrao-2018/p/10964692.html
*/

//如果通过unsafe.Pointer 计算数据的长度
//Len: &s => pointer => uintptr => pointer => *int => int
//Cap: &s => pointer => uintptr => pointer => *int => int
//&mp => pointer => **int => int
//len为什么加8 ，cap 为什么加16
func GetSliceLengthByUnsafe() {
	s := make([]int, 8, 30)
	//unsafe.Pointer 可以和任意类型的指针转换
	lenPoiter := unsafe.Pointer(uintptr(8) + uintptr(unsafe.Pointer(&s)))
	sLen := *(*int)(unsafe.Pointer(lenPoiter))
	fmt.Println(sLen, len(s))

	capPointer := unsafe.Pointer(uintptr(16) + uintptr(unsafe.Pointer(&s)))
	capLen := *(*int)(unsafe.Pointer(capPointer))
	fmt.Println(capLen, cap(s))

	mp := make(map[string]int)
	mp["qcra"] = 88
	mp["stefno"] = 66

	mpLen := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(mpLen, len(mp))

}
