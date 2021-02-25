package proxytest

import (
	"log"
	"net/http"
	"net/url"

	"xiet16.com/golearn/comunicate/middleware"
	"xiet16.com/golearn/comunicate/proxy"
)

/*
  把限流操作当中间件注入,并且处于中间件的最外层
*/
func RateLimitTest() {
	//创建下游代理
	reverseproxy := func(c *middleware.SliceRouterContext) http.Handler {
		desAddr1 := "http://127.0.0.1:2003/base"
		url1, err1 := url.Parse(desAddr1)
		if err1 != nil {

		}

		desAddr2 := "http://127.0.0.1:2004/base"
		url2, err2 := url.Parse(desAddr2)
		if err2 != nil {

		}

		urls := []*url.URL{url1, url2}
		return proxy.NewMultipleHostsReverseProxy(c, urls)
	}

	log.Println("Starting httpserver at " + addr)
	sliceRouter := middleware.NewSliceRouter()
	//使用限流中间件
	sliceRouter.Group("/").Use(middleware.RateLimiter())
	routerHandler := middleware.NewSliceRouterHandler(reverseproxy, sliceRouter)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}
