package basic

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
理解：context 可以理解主要为了处理多个goroutine间保存、读取、中断、取消等操作。
规则：
主要方法：
1 WithValue : 保存信息
2 WithTimeout
3 WithCancel
4
问题： 它是如何传递到context 中的，它的修改有没有可能产生冲突
*/

func ContextWithValueTest() {
	ctx := context.WithValue(context.Background(), "trace_id", "11112233")
	ctx = context.WithValue(ctx, "cookie", "aasddget")
	GetContextValue(ctx)
	defer ctx.Done()
}

/*
func (c *cancelCtx) Value(key interface{}) interface{} {
	if key == &cancelCtxKey {
		return c
	}
	return c.Context.Value(key)
}
*/
//递归式取值
func GetContextValue(ctx context.Context) {
	v1 := ctx.Value("trace_id")
	v1 = 334355 //这样改变不了它的值
	v2 := ctx.Value("cookie")

	log.Println("trace_id value is ", v1)
	log.Println("cookie value is ", v2)
}

type ResponseResult struct {
	r   *http.Response
	err error
}

//可以用来控制goroutine超时，context包中提供的WithTimeout(本质上调用的是WithDeadline) 方法
func ContextWithTimeoutTest() {
	ctx, cancle := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancle()

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", "http://www.baidu.com/", nil)
	if err != nil {
		log.Println("get request err", err.Error())
		return
	}

	ch := make(chan ResponseResult)
	go func() {
		res, err := client.Do(req)
		pack := ResponseResult{r: res, err: err}
		ch <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-ch
		log.Println("request timeout:", res.err)
	case res := <-ch:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		log.Println("respose data ", out)
	}
}

//WithCancel返回一个继承的Context,这个Context在父Context的Done被关闭时关闭自己的Done通道
//或者在自己被Cancel的时候关闭自己的Done
//WithCancel同时还返回一个取消函数cancel，这个cancel用于取消当前的Context
func ContextWithCancelTest() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	intchan := GetCancelChan(ctx)
	for n := range intchan {
		if n == 5 {
			log.Println("break test,n value is:", n)
			break
		}
	}
}

func GetCancelChan(ctx context.Context) chan int {
	intChan := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("exit, n value is :", n)
			case intChan <- n:
				n++
			}
		}
	}()

	return intChan
}

//deadline保存了超时的时间，当超过这个时间，会触发cancel, 如果超过了过期时间，会自动撤销它的子context
func ContextWithDeadline() {
	t := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), t)
	defer cancel()

	select {
	case <-time.After(15 * time.Second):
		log.Println("time after")
	case <-ctx.Done():
		log.Println("ctx err :", ctx.Err())
	}
}

// func ContextTest() {
// 	ctx, cancle := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))

// }
