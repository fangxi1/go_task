package main

import (
	"fmt"
	"go_task/task"
	_ "go_task/task" // 引入 task1 包
)

func main() {
	fmt.Println("完成任务1：只出现一次的数字") 
	arr :=[9]int{1,2,3,4,5,1,3,4,5}
	result := task.Task1(arr)
	fmt.Println("只出现一次的数字是:", result)
}	