package basic

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// )

// /*
// 引入Context
// */

// type Context struct {
// 	Request *http.Request
// 	Writer  http.ResponseWriter
// }

// type Handle func(http.ResponseWriter, *http.Request)
// type HandlerFunc func(c *Context)

// type Server struct {
// 	routes map[string]Handle
// }

// func (r *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	path := req.URL.String()
// 	if v, ok := r.routes[path]; ok {
// 		v(w, req)
// 		return
// 	}
// 	fmt.Println("error")
// }

// func withMiddTime(h Handle) Handle {
// 	return func(writer http.ResponseWriter, request *http.Request) {
// 		t := time.Now()
// 		defer func() {
// 			log.Println("MiddleWare(withMiddTime) spend is ", time.Since(t))
// 		}()
// 		h(writer, request)
// 	}
// }

// func withMiddLog(h Handle) Handle {
// 	return func(writer http.ResponseWriter, request *http.Request) {
// 		log.Printf("MiddleWare(withMiddLog) Request URL(%s) Method(%v) ", request.URL, request.Method)
// 		h(writer, request)
// 	}
// }

// func (r *Server) createContext(w http.ResponseWriter, req *http.Request) *Context {
// 	return &Context{
// 		Request: req,
// 		Writer:  w,
// 	}
// }

// //注入方式，和执行方法的传参变了
// func (r *Server) Register(route string, f HandlerFunc) {
// 	r.routes[route] = withMiddLog(withMiddTime(func(writer http.ResponseWriter, request *http.Request) {
// 		f(r.createContext(writer, request))
// 	}))
// }

// func New() *Server {
// 	return &Server{
// 		routes: make(map[string]Handle),
// 	}
// }

// func Test() {
// 	s := New()
// 	s.Register("/bench", func(c *Context) {
// 		time.Sleep(time.Second)
// 		fmt.Println("bench sleep 1 second")
// 		c.Writer.Write([]byte("hello!\r\n"))
// 	})
// 	s.Register("/hello", func(c *Context) {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("hello sleep 2 second")
// 		_, err := c.Writer.Write([]byte("world\r\n"))
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	})
// 	http.ListenAndServe(":2102", s)
// }
