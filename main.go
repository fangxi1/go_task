package main

import (
	"fmt"
	"go_task/task" // 引入 task 包
	// "time"
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

	// 合并区间
	// intervals := [][]int{{6,8},{2,6},{1,3}}
	// task6 := task.Task6(intervals)
	// fmt.Println("target=",task6)


	// 找出两个数之和=target
	// arrys := []int{1,2,3,3,4,5}
	// target := 5
	// task6 := task.Task7(arrys,target)
	// fmt.Println("target=",task6)

	// a := 2
	// var p *int
	// p = &a
	// i := task.Task_1(p)
	// fmt.Println("进阶任务1=",i)

	// arr := []int{1,4,3}
	// var p []*int
	// for _,value :=range arr{
	// 	p = append(p, &value)
	// }
	// b := task.Task_2(p)
	// if len(b) > 0 && b[0] != nil {
	// 	fmt.Println("进阶任务2=", *b[0], *b[1], *b[2])
	// } else {
	// 	fmt.Println("进阶任务2= 无结果")
	// }

	// task_2_arr := []int{1,2,3,4,5,6,7,8,9,10}
	// for _, i := range task_2_arr {
	// 	go func() {
	// 		task.Task_3(i)
	// 	}()
	// }
	// time.Sleep(5*time.Second)
	// fmt.Println("进阶任务3完成")

	// task_2_4_arr := []int{30,20,40,50,60,90,100,200,300,10}
	// task.Task_4(task_2_4_arr)
	// time.Sleep(5*time.Second)
	// fmt.Println("进阶任务4完成")
	// 矩形面积
	// retangle := task.Rectangle{Width:1.2,Height:3.1}
	// task.ShapeImpl(retangle,"retangle")// 矩形
	// circle := task.Circle{Radius: 3.4}
	// task.ShapeImpl(circle,"circle")// 圆形
	// fmt.Println("进阶任务5完成")

	// p := task.Person{Age: "32",Name: "测试"}
	// e := task.Employee{EmployeeID: 21,Person: p}
	// task.Print.PrintInfo(e)
	// fmt.Println("进阶任务6完成")

	// var ch chan int
	// ch = make(chan int)
	// task.ChannelTest1(ch)
	// time.Sleep(5*time.Second)
	// fmt.Println("进阶任务7完成")
	
	
	// cha := make(chan int,30)
	// task.ChannelTest2(cha)
	// time.Sleep(5*time.Second)
	// fmt.Println("进阶任务8完成")

	task.LockTest()
	fmt.Println("进阶任务9完成")

	
	task.LockTest2()
	fmt.Println("进阶任务10完成")
}
