package task
 import "fmt"

 func Task2(s string) bool {
	// 先把有效字符存入一个map中
	var strMap = make(map[rune]rune,3)
	strMap['('] = ')'
	strMap['{'] = '}'
	strMap['['] = ']'
	
	fmt.Printf("s=%q strMap=%v \n", s,strMap)
	// 声明一个栈来存在符号
	stack := make([]rune, 0, len(s))
	
	fmt.Printf("stack len=%v cap=%d\n", len(stack),cap(stack))
	// 遍历字符串
	for _,char := range s {
		fmt.Printf("char=%q\n", char)
		// 如果是左括号，则入栈
		if _,ok := strMap[char];ok{
			fmt.Println("符号入栈:", char)
			stack = append(stack, char)
			fmt.Printf("stack len=%v cap=%d\n", len(stack),cap(stack))
			continue
		}
		// 如果不是左边的符号，则去栈中取出存储的符号，做匹配
		if len(stack)==0 {
			fmt.Println("栈中没有左括号，符号无效")
			return false
		}
		// 取出栈里面元素
		str := stack[0]
		if strMap[str] == char{
			fmt.Println("符号匹配成功:", str, "和", char)
			// 匹配成功，出栈
			stack = stack[1:]
			fmt.Printf("stack len=%v cap=%d\n", len(stack),cap(stack))
			continue
		}else{
			fmt.Println("符号匹配失败:", str, "和", char)
			return false
		}
	}
	return true;
 }