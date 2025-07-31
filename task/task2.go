package task

import "fmt"

func Task_1(a *int) int {
	*a = 10
	fmt.Println(*a)
	return *a
}
func Task_2(a []*int) []*int {
	for i,value := range a{
		fmt.Println(*value)
		*a[i] = *value * 2
	}
	return a
}