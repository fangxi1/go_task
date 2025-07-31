package main

import (
	"fmt"
	"go_task/task" // 引入 task 包
)

func main() {
	// fmt.Println("完成任务1：只出现一次的数字")
	// arr :=[9]int{1,2,3,4,5,1,3,4,5}
	// result := task.Task1(arr)
	// fmt.Println("只出现一次的数字是:", result)

	// fmt.Println("完成任务2：括号匹配")
	// s := "((){}[]"
	// result2 := task.Task2(s)
	// if result2 {
	// 	fmt.Println("符号有效")
	// } else {
	// 	fmt.Println("符号无效")
	// }

	// fmt.Println("完成任务3：匹配最长公共前缀")
	// arrStr := []string{"flower", "flow", "flight"}
	// result3 := task.Task3(arrStr)
	// fmt.Println("最长公共前缀是:", result3)


	// fmt.Println("完成任务4：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一")
	// param := []uint{9, 8, 9, 6, 8}
	// digits := task.Task4(param)
	// fmt.Println("加一之后的数组为:", digits)

	// fmt.Println("完成任务5：删除有序数组中的重复项")
	// nums := []int{ 8, 8,9,9, 9}
	// task5 := task.Task5(nums)
	// fmt.Println("删除有序数组中的重复项,之后数组长度:", task5)

	// 找出两个数之和=target
	// arrys := []int{1,2,3,3,4,5}
	// target := 5
	// task6 := task.Task6(arrys,target)
	// fmt.Println("target=",task6)

	// a := 2
	// var p *int
	// p = &a
	// i := task.Task_1(p)
	// fmt.Println("进阶任务1=",i)

	arr := []int{1,4,3}
	var p []*int
	for _,value :=range arr{
		p = append(p, &value)
	}
	b := task.Task_2(p)
	if len(b) > 0 && b[0] != nil {
		fmt.Println("进阶任务2=", *b[0], *b[1], *b[2])
	} else {
		fmt.Println("进阶任务2= 无结果")
	}
}
