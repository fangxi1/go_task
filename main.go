package main

import (
	"fmt"
	"go_task/task" // 引入 task1 包
)

func main() {
	fmt.Println("完成任务1：只出现一次的数字")
	arr :=[9]int{1,2,3,4,5,1,3,4,5}
	result := task.Task1(arr)
	fmt.Println("只出现一次的数字是:", result)

	fmt.Println("完成任务2：括号匹配")
	s := "((){}[]"
	result2 := task.Task2(s)
	if result2 {
		fmt.Println("符号有效")
	} else {
		fmt.Println("符号无效")
	}

	fmt.Println("完成任务3：匹配最长公共前缀")
	arrStr := []string{"flower", "flow", "flight"}
	result3 := task.Task3(arrStr)
	fmt.Println("最长公共前缀是:", result3)


	fmt.Println("完成任务4：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一")
	param := []uint{9, 8, 9, 6, 8}
	digits := task.Task4(param)
	fmt.Println("加一之后的数组为:", digits)


}
