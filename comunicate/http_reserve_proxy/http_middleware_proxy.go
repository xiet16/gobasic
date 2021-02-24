package http_reserve_proxy

import (
	"net/http"
	"net/url"
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
		url2, err2 := url.Parse(dst1)
		if err2 != nil {
		}
		urls := []*url.URL{url1, url2}
		return proxy.NewMultipleHostsReverseProxy(c, urls)
	}
}
