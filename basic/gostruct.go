package basic

import (
	"fmt"
	"time"
)

//go 中的自定义类型 struct , 重定义类型

type interger int

type UserInfo struct {
	Name        string
	Age         int
	Birthday    time.Time
	Hobby       []string
	Description map[string]interface{}
}

func StructTest() {
	var intVar int = 1
	var intergerVar interger
	intergerVar = interger(intVar)
	fmt.Println(intergerVar)

	var peopleVar UserInfo
	peopleVar.Name = "xiet"
	peopleVar.Age = 18
	peopleVar.Birthday = time.Now()
	peopleVar.Hobby = []string{"爬山", "看书"}
	peopleVar.Description = map[string]interface{}{
		"company": "赛盒",
		"address": "宝安永丰社区",
	}
	fmt.Println(peopleVar)

	user := UserInfo{
		Name: "boge",
	}
	fmt.Printf("%v \n", user)

	var xiaoming *UserInfo
	xiaoming = new(UserInfo)
	(*xiaoming).Name = "xiaopeng"
	(*xiaoming).Age = 27
	fmt.Println(xiaoming)
}
