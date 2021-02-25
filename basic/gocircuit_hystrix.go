package basic

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func HystrixCircuitTest() {
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":2111", hystrixStreamHandler)

	hystrix.ConfigureCommand("aaa", hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  1, //最大并发数
		RequestVolumeThreshold: 1,
		SleepWindow:            5000,
		ErrorPercentThreshold:  1,
	})

	for i := 0; i < 100000; i++ {
		//异步调用并使用 hystrix.go
		err := hystrix.Do("aaa", func() error {
			//并发测试
			if i == 0 {
				return errors.New("service first do")
			}
			log.Println("do something")
			return nil
		}, nil)

		if err != nil {
			log.Println("hystrix error:" + err.Error())
			time.Sleep(1 * time.Second)
			log.Println("sleep 1 second")
		}
	}

	time.Sleep(100 * time.Second)
}

var Number int

var Result string

func HystrixTest2() {

	config := hystrix.CommandConfig{
		Timeout:                2000, //超时时间设置  单位毫秒
		MaxConcurrentRequests:  8,    //最大请求数
		SleepWindow:            5000, //过多长时间，熔断器再次检测是否开启。单位毫秒
		ErrorPercentThreshold:  30,   //错误率
		RequestVolumeThreshold: 5,    //请求阈值  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
	}
	hystrix.ConfigureCommand("test", config)
	cbs, _, _ := hystrix.GetCircuit("test")
	defer hystrix.Flush()
	for i := 0; i < 50; i++ {
		start1 := time.Now()
		Number = i
		hystrix.Go("test", run, getFallBack)
		fmt.Println("请求次数:", i+1, ";用时:", time.Now().Sub(start1), ";请求状态 :", Result, ";熔断器开启状态:", cbs.IsOpen(), "请求是否允许：", cbs.AllowRequest())
		time.Sleep(1000 * time.Millisecond)
	}
	time.Sleep(20 * time.Second)
}

func run() error {
	Result = "RUNNING1"
	if Number > 10 {
		return nil
	}
	if Number%2 == 0 {
		return nil
	} else {
		return errors.New("请求失败")
	}
}

func getFallBack(err error) error {
	Result = "FALLBACK"
	return nil
}
