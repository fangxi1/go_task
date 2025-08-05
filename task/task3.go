package task

import (
	"fmt"
	"math"
)

type Shape interface{
	// 面积
	Area() float64
	// 周长
	Perimeter() float64
}


type Rectangle struct{
	Width float64// 宽
	Height float64//长

}

type Circle struct{
	Radius float64//半径

}

// 矩形面积
func (r Rectangle) Area() float64{
	return r.Width*r.Height
}
// 圆形面积
func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

// 矩形周长
func (r Rectangle) Perimeter() float64{
	return 2 * (r.Width+r.Height)
}

// 圆形周长
func (c Circle) Perimeter() float64{
	return 2*math.Pi*c.Radius
}

func ShapeImpl(shape Shape,name string) {
	fmt.Printf("名字=%s_area=%v \n",name,shape.Area())
	fmt.Printf("名字=%s_Perimeter=%v \n",name,shape.Perimeter())
}


type Person struct{
	Name string
	Age string
	
}

type Employee struct{
	EmployeeID int
	Person 
}

type Print interface{
	PrintInfo() 
}


func (e Employee) PrintInfo() {
	fmt.Printf("员工Id=%d,姓名=%s,年龄=%s",e.EmployeeID,e.Name,e.Age)
}
