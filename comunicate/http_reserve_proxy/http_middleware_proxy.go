package http_reserve_proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"xiet16.com/golearn/comunicate/middleware"
	"xiet16.com/golearn/comunicate/proxy"
)

/*
定义切片路由middleware.NewSliceRouter()

*/
func middlewareProxy() {
	reverseProxy := func(c *middleware.SliceRouterContext) http.Handler {
		dst1 := ""
		url1, err1 := url.Parse(dst1)
		if err1 != nil {

		}

		dst2 := ""
		url2, err2 := url.Parse(dst2)
		if err2 != nil {
		}
		urls := []*url.URL{url1, url2}
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
