package http_reserve_proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

/*
httputil.NewSingleHostReverseProxy(desUrl)这个方法定义了结构体ReverseProxy 的Director
这里我们实现自定义ModifyResponse，所以我们要自己实现NewSingleHostReverseProxy
*/

func httpProxy() {
	desAddr := "127.0.0.1:2003"
	desUrl, err := url.Parse(desAddr)
	if err != nil {
		log.Println(err)
	}
	//修改这个操作
	//proxy := httputil.NewSingleHostReverseProxy(desUrl)
	proxy := DefineNewSingleHostReverseProxy(desUrl)
	log.Println("Start httputil proxy server at :" + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}

func DefineNewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path, req.URL.RawPath = DefinejoinURLPath(target, req.URL)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	//添加modify

	modifyFunc := func(*http.Response) error {

	}
	return &httputil.ReverseProxy{Director: director}
}

func DefinejoinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
