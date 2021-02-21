package basic

import (
	"fmt"
	"time"
)

var (
	Name      string
	lastLogin time.Time
)

func Vartest() {
	Name = "xiet"
	lastLogin = time.Now()
	fmt.Println("变量测试")
	fmt.Println("变量的定义")
	var v1 int = 1
	v2 := "字符串"
	var v3 = 100
	fmt.Printf("%d,%s,%d,now time is %s \n", v1, v2, v3, lastLogin)
}
