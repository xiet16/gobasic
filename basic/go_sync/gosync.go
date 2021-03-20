package go_sync

import (
	"fmt"
	"sync"
	"time"
)

/*
sync 包提供了并发编程同步原语
相关url:
https://blog.csdn.net/kevin_tech/article/details/105935251
https://www.jianshu.com/p/4e2922f68991
常用的:
锁: sync.Locker()
等待
*/

func SyncLockTest()  {
	lock :=&sync.Mutex{}
	lock.Lock()
	defer  lock.Unlock()
	fmt.Println("sync 包中的锁，Lock,和unLock 是接口中的方法，Mutex中实现了它")
}

func SyncWaitAllTest()  {
	gwait:= &sync.WaitGroup{}

	for i:=5;i<10;i++ {
		gwait.Add(1)
		go func() {
			 it:=i
             fmt.Println("task begin ",it)
             time.Sleep(time.Second * time.Duration(it-5+1))
             gwait.Done()
			fmt.Println("task end ",it)
		}()
	}
	gwait.Wait()
	fmt.Println("wait end")
}
