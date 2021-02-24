package basic

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// )

// /*
// 相关url:
// https://www.cnblogs.com/alin-qu/p/11087811.html
// https://studygolang.com/articles/1240
// https://www.jianshu.com/p/16210100d43d
// https://github.com/quguolin/goBase/blob/master/net/httpMiddleware

// 中间件像洋葱，一层层的往下执行
// go中要实现中间件，只要实现 func (http.Handler) http.Handler 就好了
// */

// func ChainMiddlewareTest() {
// 	http.Handle("/", loggerMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("do something")
// 	})))

// 	http.ListenAndServe(":2102", nil)
// }

// func loggerMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		fmt.Println("loggerMiddleware write start", r.Method, r.URL.Path, start)
// 		next.ServeHTTP(w, r)
// 		fmt.Println("loggerMiddleware write end")
// 	})
// }

// /*
// 实现了ServeHTTP方法,就是实现了http.Handler 接口
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }
// 用一个结构体来存储路由规则
// */

// func Middleware_V1Test() {
// 	router := NewMiddlewareRouter()
// 	router.Register("/bench", func(writer http.ResponseWriter, request *http.Request) {
// 		fmt.Println("bench")
// 	})
// 	router.Register("/hello", func(writer http.ResponseWriter, request *http.Request) {
// 		fmt.Println("hello")
// 	})

// 	http.ListenAndServe(":2102", router)
// }

// type Handle func(http.ResponseWriter, *http.Request)
// type MiddlewareRouter struct {
// 	routes map[string]Handle
// }

// //
// func (router *MiddlewareRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("do router servehttp", r.URL.String())
// 	if v, ok := router.routes[r.URL.String()]; ok {
// 		v(w, r)
// 		return
// 	}
// 	fmt.Println("not match path")
// }

// func NewMiddlewareRouter() *MiddlewareRouter {
// 	return &MiddlewareRouter{
// 		routes: make(map[string]Handle),
// 	}
// }

// func (router *MiddlewareRouter) Register_V1(path string, handle Handle) {
// 	router.routes[path] = handle
// }

// func (router *MiddlewareRouter) Register_V2(path string, handle Handle) {
// 	router.routes[path] = withMiddle(handle)
// }

// //又嵌套一层
// func (router *MiddlewareRouter) Register(path string, handle Handle) {
// 	router.routes[path] = withMiddLog(withMiddle(handle))
// }

// func withMiddle(handle Handle) Handle {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		now := time.Now()
// 		defer func() {
// 			fmt.Println("time span is", time.Since(now))
// 		}()
// 		handle(w, r)
// 	}
// }

// func withMiddLog(h Handle) Handle {
// 	return func(writer http.ResponseWriter, request *http.Request) {
// 		log.Printf("Request URL(%s) Method(%v) ", request.URL, request.Method)
// 		h(writer, request)
// 	}
// }
