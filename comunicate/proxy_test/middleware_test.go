package proxy_test

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"xiet16.com/golearn/comunicate/middleware"
	"xiet16.com/golearn/comunicate/proxy"
)

var addr = "127.0.0.1:2002"

func middleware_proxy_test() {
	//启动两个下游服务器
	reverseProxy := func(c *middleware.SliceRouterContext) http.Handler {
		desAddr1 := ""
		desUrl1, err1 := url.Parse(desAddr1)
		if err1 != nil {
			log.Println(err1)
		}

		desAddr2 := ""
		desUrl2, err2 := url.Parse(desAddr2)
		if err2 != nil {
			log.Println(err2)
		}

		urls := []*url.URL{desUrl1, desUrl2}
		return proxy.NewMultipleHostsReverseProxy(c, urls)
	}
	log.Println("starting httpserver at " + addr)

	//初始化方法数组路由器
	sliceRouter := middleware.NewSliceRouter()
	//中间件可充当业务逻辑代码
	sliceRouter.Group("/base").Use(middleware.TraceLogSliceMW(), func(c *middleware.SliceRouterContext) {
		c.Rw.Write([]byte("test func"))
	})

	sliceRouter.Group("/").Use(middleware.TraceLogSliceMW(), func(c *middleware.SliceRouterContext) {
		fmt.Println("reverseProxy")
		reverseProxy(c).ServeHTTP(c.Rw, c.Req)
	})
	routerHandler := middleware.NewSliceRouterHandler(nil, sliceRouter)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}
