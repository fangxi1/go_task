package task

import (
	"fmt"
	"time"
)

func Task_1(a *int) int {
	*a += 10
	fmt.Println(*a)
	return *a
}
func Task_2(a []*int) []*int {
	for i, value := range a {
		fmt.Println(*value)
		*a[i] = *value * 2
	}
	return a
}

func Task_3(i int) {
	if i%2 == 1 {
		fmt.Println("奇数=", i)
	} else {
		fmt.Println("偶数=", i)
	}
}
// 入参是阻塞时间
func Task_4(date []int){
	fmt.Println(date)
	for i:=0;i<len(date);i++ {
		go func(){
			// 执行任务开始时间
			now := time.Now().UnixMilli()
			time.Sleep(time.Duration(date[i]) * time.Millisecond)
			// 执行任务结束时间
			endDate := time.Now().UnixMilli()
			fmt.Printf("任务i=%d执行耗时=%d \n",date[i],endDate-now)
		}()
	}
	
}
