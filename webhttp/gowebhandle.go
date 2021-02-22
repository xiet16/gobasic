package webhttp

import (
	"fmt"
	"net/http"
	"net/url"
)

//HandleFunc方式
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过http.HandleFunc启动一个http服务,要传入ServeHTTP 方法")
}

//自定义handler 的方式 要实现ServeHttp
type DefineHandle struct{}

func (handler *DefineHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自定义DefineHandle启动一个http服务")
	//获取url
	fmt.Fprintln(w, r.URL.RawQuery)
	fmt.Fprintln(w, r.URL.Host)
	fmt.Fprintln(w, r.URL.Path)

	rawQuery := r.URL.RawQuery
	va, _ := url.ParseQuery(rawQuery)
	fmt.Fprintln(w, "获取url当中rawquery的值", va.Get("name"))
}

func HttpHandleTest() {
	http.HandleFunc("/", ServeHTTP)
	defineHandle := DefineHandle{}
	http.Handle("/getuserinfo", &defineHandle)
	http.ListenAndServe("0.0.0.0:8380", nil)
}
