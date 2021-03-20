package go_tag

import (
	"fmt"
	"reflect"
)

/*
Tag: 结构体字段声明后面跟着一个可选字符串，例如
type T struct {
	f1     string `f three`
    f2     string
    f3, f4 int64  `f four and five`
}
*/

/*
通过反射可以取出值
if t.Kind() != Struct {
	panic("reflect: FieldByName of non-struct type " + t.String())
}
tt := (*structType)(unsafe.Pointer(t))
return tt.FieldByName(name)
*/
type T struct {
	f1     string `f1 tag`
	f2     string
	f3, f4 int64 `custom_validate:"number", lowlength: "100"`
}

func GetTagInfoByReflection() {
	typeInfo := reflect.TypeOf(T{})
	field1, ok := typeInfo.FieldByName("f1")
	if !ok {
		return
	}

	fmt.Println("f1 tag is:", field1.Tag)

	field3, ok := typeInfo.FieldByName("f3")
	if !ok {
		return
	}

	value, _ := field3.Tag.Lookup("custom_validate")
	fmt.Println("f3 tag is:", field3.Tag)
	fmt.Printf("f3 validate_value is :%s \n", value)
}
