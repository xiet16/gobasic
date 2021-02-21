package basic

import (
	"fmt"
	"unsafe"
)

//go 默认float64
//复数 complex
// float---第31位(占1bit)---第30-23位(占8bit)----第22-0位(占23bit)
// double--第63位(占1bit)---第62-52位(占11bit)---第51-0位(占52bit)

func FloatTest() {
	floatVar := .1415926
	fmt.Printf("%s,%T,%d", floatVar, floatVar, unsafe.Sizeof(floatVar))
}
