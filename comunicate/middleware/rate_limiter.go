package middleware

import (
	"fmt"

	"golang.org/x/time/rate"
)

func RateLimiter() func(c *SliceRouterContext) {
	limiter := rate.NewLimiter(1, 2)
	return func(c *SliceRouterContext) {
		if !limiter.Allow() {
			c.Rw.Write([]byte(fmt.Sprintf("rate limiter no allow , limit:%v, burst:%v ", limiter.Limit(), limiter.Burst())))
			c.Abort()
			return
		} else {
			fmt.Println("limiter allow")
		}
		c.Next()
	}
}
