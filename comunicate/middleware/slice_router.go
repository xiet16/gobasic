package middleware

import (
	"context"
	"math"
	"net/http"
	"strings"
)

//最多63个中间件
const abortIndex int8 = math.MaxInt8 / 2

type HandlerFunc func(*SliceRouterContext)

type SliceRouter struct {
	groups []*SliceGroup
}

//路由注册结构图，用于承载路由和中间件的注册
type SliceGroup struct {
	*SliceRouter
	path     string
	handlers []HandlerFunc
}

//定义一个切片路由上下文，用于承载中间件的执行
type SliceRouterContext struct {
	Rw  http.ResponseWriter
	Req *http.Request
	Ctx context.Context
	*SliceGroup
	index int8
}

func newSliceRouterContext(rw http.ResponseWriter, req *http.Request, r *SliceRouter) *SliceRouterContext {
	newSliceGroup := &SliceGroup{}

	//最长url前缀匹配
	matchUrlLen := 0
	for _, group := range r.groups {
		//fmt.Println("req.RequestURI")
		//fmt.Println(req.RequestURI)
		if strings.HasPrefix(req.RequestURI, group.path) {
			pathLen := len(group.path)
			if pathLen > matchUrlLen {
				matchUrlLen = pathLen
				*newSliceGroup = *group //浅拷贝数组指针
			}
		}
	}

	c := &SliceRouterContext{Rw: rw, Req: req, SliceGroup: newSliceGroup, Ctx: req.Context()}
	c.Reset()
	return c
}

func (c *SliceRouterContext) Get(key interface{}) interface{} {
	return c.Ctx.Value(key)
}

func (c *SliceRouterContext) Set(key, val interface{}) {
	c.Ctx = context.WithValue(c.Ctx, key, val)
}

type SliceRouterHandler struct {
	coreFunc func(*SliceRouterContext) http.Handler
	router   *SliceRouter
}

func (w *SliceRouterHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := newSliceRouterContext(rw, req, w.router)
	if w.coreFunc != nil {
		c.handlers = append(c.handlers, func(c *SliceRouterContext) {
			w.coreFunc(c).ServeHTTP(rw, req)
		})
	}
	c.Reset()
	c.Next()
}
func NewSliceRouterHandler(coreFunc func(*SliceRouterContext) http.Handler, router *SliceRouter) *SliceRouterHandler {
	return &SliceRouterHandler{
		coreFunc: coreFunc,
		router:   router,
	}
}

//构造路由器
func NewSliceRouter() *SliceRouter {
	return &SliceRouter{}
}

//添加路由规则
func (g *SliceRouter) Group(path string) *SliceGroup {
	return &SliceGroup{
		SliceRouter: g,
		path:        path,
	}
}

//为路由添加中间件,同时确保两个SliceRouter SliceGroup 中的方法是一致的
func (g *SliceGroup) Use(middlewares ...HandlerFunc) *SliceGroup {
	g.handlers = append(g.handlers, middlewares...)
	existsFlag := false
	for _, oldGroup := range g.SliceRouter.groups {
		if oldGroup == g {
			existsFlag = true
		}
	}
	if !existsFlag {
		g.SliceRouter.groups = append(g.SliceRouter.groups, g)
	}
	return g
}

// 从最先加入中间件开始回调
func (c *SliceRouterContext) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		//fmt.Println("c.index")
		//fmt.Println(c.index)
		c.handlers[c.index](c)
		c.index++
	}
}

// 跳出中间件方法
func (c *SliceRouterContext) Abort() {
	c.index = abortIndex
}

// 重置回调
func (c *SliceRouterContext) Reset() {
	c.index = -1
}
