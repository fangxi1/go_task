package task

import (
	"fmt"
	"sort"
)

/*
只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，
然后再遍历 map 找到出现次数为1的元素。
*/
func Task1(arr [9]int) int {
	// 先声明个map 来记录每个元素出现的次数
	countMap := make(map[int]int)
	// 再遍历整数数组
	for _, num := range arr {
		// 如果countMap中没有这个元素，则把该元素加入map中，并设置key为该元素值，指为1
		if _, ok := countMap[num]; !ok {
			countMap[num] = 1
			continue
		}
		// 如果countMap中有该元素，则把该元素值加1
		countMap[num] += 1
	}
	fmt.Println("countMap:", countMap)
	// 再遍历Map，找到出现次数最少的元素
	for key, value := range countMap {
		fmt.Println("key:", key, "value:", value)
		if value == 1 {
			return key
		}
	}
	return 0
}

// 括号匹配：给定一个字符串，包含括号、花括号和方括号等符号，
func Task2(s string) bool {
	// 先把有效字符存入一个map中
	var strMap = make(map[rune]rune, 3)
	strMap['('] = ')'
	strMap['{'] = '}'
	strMap['['] = ']'

	fmt.Printf("s=%q strMap=%v \n", s, strMap)
	// 声明一个栈来存在符号
	stack := make([]rune, 0, len(s))

	fmt.Printf("stack len=%v cap=%d\n", len(stack), cap(stack))
	// 遍历字符串
	for _, char := range s {
		fmt.Printf("char=%q\n", char)
		// 如果是左括号，则入栈
		if _, ok := strMap[char]; ok {
			fmt.Println("符号入栈:", char)
			stack = append(stack, char)
			fmt.Printf("stack len=%v cap=%d\n", len(stack), cap(stack))
			continue
		}
		// 如果不是左边的符号，则去栈中取出存储的符号，做匹配
		if len(stack) == 0 {
			fmt.Println("栈中没有左括号，符号无效")
			return false
		}
		// 取出栈里面元素
		str := stack[0]
		if strMap[str] == char {
			fmt.Println("符号匹配成功:", str, "和", char)
			// 匹配成功，出栈
			stack = stack[1:]
			fmt.Printf("stack len=%v cap=%d\n", len(stack), cap(stack))
			continue
		} else {
			fmt.Println("符号匹配失败:", str, "和", char)
			return false
		}
	}
	return true
}

// 最长公共前缀
func Task3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var prefix string
	for i := 0; i < len(strs[0]); i++ {
		flag := false
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 {
				// 如果第一个数组的第一个字母大于第二个数组的长度，则跳出循环
				fmt.Println("第一个数组的第一个字母大于第二个数组的长度，跳出循环")
				return prefix

			}
			// 第一个数组的第一个字母和第二个数组的第一个字符对比
			if strs[0][i] == strs[j][i] && j == len(strs)-1 {
				flag = true
			}
		}
		if flag {
			prefix = prefix + string(strs[0][i])
		}
	}
	return prefix
}

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func Task4(digits []uint) []uint {
	var num = 0
	// 遍历数组
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			num = 1 // 如果是最后一个元素，则需要加1
		}
		if (int(digits[i]) + num) == 10 {
			num = 1
			digits[i] = 0
		} else if num == 1 {
			digits[i] = digits[i] + 1
			num = 0
		}
		if num == 1 && i == 0 {
			// 如果是第一个元素，且需要进位，则在前面添加1
			digits = append([]uint{1}, digits...)
		}
	}
	return digits
}

/*
删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，
使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，
你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，
一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，
并将 i 后移一位。
*/
func Task5(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0 // 慢指针，指向已处理的最后一个不重复元素
    for j := 1; j < len(nums); j++ { // 快指针，遍历数组
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j]
        }
    }
    return i + 1 // 新数组长度
}

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，
遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
如果没有重叠，则将当前区间添加到切片中。
*/
func Task6 (intervals [][]int) [][] int{
	sort.Slice(intervals,func(i int,j int) bool{
		return intervals[i][0]<intervals[j][0]
	})
	intelvalsNew := [][]int{intervals[0]}
	fmt.Println("intervlasNew=",intelvalsNew)
	for i :=1;i<len(intervals);i++{
		curr := intervals[i]
		last :=intelvalsNew[len(intelvalsNew)-1]
		if(curr[0]<=last[1]){
			// 有重叠，合并区间
			if(curr[1]>last[1]){
				last[1] = curr[1]
			}
			intelvalsNew[len(intelvalsNew)-1] = last
		}else {
			intelvalsNew = append(intelvalsNew, curr)
		}
	}
	return intelvalsNew
}


// 两数之和 
func Task7(arr []int, target int) []int{
	var nums = [2]int{}
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			sum :=arr[i]+arr[j]
			if(sum == target){
				nums[0] = arr[i]
				nums[1] = arr[j]
				break 
			}
		}
	}
	return nums[:]
}


