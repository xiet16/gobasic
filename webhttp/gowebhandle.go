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

	//获取header
	for key := range r.Header {
		fmt.Fprintf(w, "请求头%s的值为:%s\n", key, r.Header[key])
	}

	//获取body
	//body是一个接口，继承了 io.ReadCloser,  ReadCloser 继承了Reader、Closer , Reader接口定了了Read 方法
	//Read(p []byte) (n int, err error)

	//curl -id "action=search&resposne body" localhost:8380/getuserinfo
	len := r.ContentLength
	bodydata := make([]byte, len)
	r.Body.Read(bodydata)
	fmt.Fprintln(w, string(bodydata))
}

func HttpFormAndPostFrom(w http.ResponseWriter, r *http.Request) {
	//获取form 和postform This field is only available after ParseForm is called.
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.PostForm)
}

//这个是文件的Form 传输方式 postman 要怎么弄
func HttpMultiPartFrom(w http.ResponseWriter, r *http.Request) {
	//要定义最大内存
	r.ParseMultipartForm(1024)

}

func HttpHandleTest() {
	http.HandleFunc("/", ServeHTTP)
	http.HandleFunc("/getFormAndPostFrom", HttpFormAndPostFrom)

	defineHandle := DefineHandle{}
	http.Handle("/getuserinfo", &defineHandle)
	http.ListenAndServe("0.0.0.0:8380", nil)
}
