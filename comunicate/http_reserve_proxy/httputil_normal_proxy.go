package http_reserve_proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

/*
 */

var (
	addr = "127.0.0.1:2002"
)

func httpultiProxy() {
	desAddr := "127.0.0.1:2003"
	desUrl, err := url.Parse(desAddr)
	if err != nil {
		log.Println(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(desUrl)
	log.Println("Start httputil proxy server at :" + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
