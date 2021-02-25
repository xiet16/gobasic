package basic

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/time/rate"
)

/*
 限流的两种方式，
 漏桶限流: 每次取时计算痛的流量，桶超过阈值则降级
 令牌桶限流：每次从桶里取令牌，取不到则降级
 rate limit 的三种状态
 Allow 判断当前是否可以取到token
 Wait 阻塞知道取到token
 Reverse 返回等待时间，到了时间再去取
*/

func RateLimitTest() {
	limiter := rate.NewLimiter(1, 5)
	log.Println(limiter.Limit, limiter.Burst())
	for i := 0; i < 100; i++ {
		//阻塞等待，直到取到一个token
		log.Println("before wait")
		c, _ := context.WithTimeout(context.Background(), time.Second*2)
		if err := limiter.Wait(c); err != nil {
			log.Println("limiter wait err:" + err.Error())
		}
		log.Println("after wait")
		r := limiter.Reserve()
		log.Printf("limeter Reserve: %d", r.Delay())

		if !limiter.Allow() {
			fmt.Println("limiter no allow")
		} else {
			fmt.Println("limiter allow")
		}
		time.Sleep(200 * time.Millisecond)
		log.Println(time.Now)
	}
}
