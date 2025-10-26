package one

import (
	"fmt"
	"math"
)

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义一个结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func test() {
	c := Circle{Radius: 5}
	var s Shape = c // 接口变量可以存储实现了接口的类型
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
