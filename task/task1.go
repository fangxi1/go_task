package task

import "fmt"

/*
只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，
然后再遍历 map 找到出现次数为1的元素。
*/
func Task1(arr [9]int) int{
	// 先声明个map 来记录每个元素出现的次数
	countMap := make(map[int]int)
	// 再遍历整数数组
	for _,num :=range arr{
		// 如果countMap中没有这个元素，则把该元素加入map中，并设置key为该元素值，指为1
		if _,ok:=countMap[num]; !ok{
			countMap[num] = 1
			continue
		}
		// 如果countMap中有该元素，则把该元素值加1
		countMap[num] += 1
	}
	fmt.Println("countMap:", countMap)
	// 再遍历Map，找到出现次数最少的元素
	for key,value := range countMap{
		fmt.Println("key:", key, "value:", value)
		if value == 1 {
			return key
		}
	}
	return 0;
}