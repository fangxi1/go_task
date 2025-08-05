package task

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func LockTest(){
	var ck sync.Mutex
	count := 0
	for i:=0;i<10;i++{
		go func() {
			for i := 1; i <= 1000; i++ {
				ck.Lock()
				count++
				ck.Unlock()
			}
			
		}()
	}
	time.Sleep(3*time.Second)
	fmt.Println("last count=",count)
}


func LockTest2(){
	var count int64= 0
	for i:=0;i<10;i++{
		go func() {
			for i := 1; i <= 1000; i++ {
				atomic.AddInt64(&count,1)
			}
			
		}()
	}
	time.Sleep(3*time.Second)
	fmt.Println("last count=",count)
}