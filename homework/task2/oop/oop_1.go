package main

import (
	"fmt"
	"math"
)

// Shape 接口定义了面积和周长的计算方法
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64 // 宽度
	Height float64 // 高度
}

// Area 计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 圆形结构体
type Circle struct {
	Radius float64 // 半径
}

// Area 计算圆形面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆形周长（ circumference ）
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建矩形实例（宽 5，高 3）
	rect := Rectangle{Width: 5, Height: 3}
	// 创建圆形实例（半径 4）
	circle := Circle{Radius: 4}

	// 直接调用矩形的方法
	fmt.Println("矩形:")
	fmt.Printf("  面积: %.2f\n", rect.Area())
	fmt.Printf("  周长: %.2f\n", rect.Perimeter())

	// 直接调用圆形的方法
	fmt.Println("\n圆形:")
	fmt.Printf("  面积: %.2f\n", circle.Area())
	fmt.Printf("  周长: %.2f\n", circle.Perimeter())

	// 利用接口切片统一处理不同形状（多态特性）
	shapes := []Shape{rect, circle}
	fmt.Println("\n通过接口处理:")
	for i, s := range shapes {
		fmt.Printf("  形状 %d 面积: %.2f\n", i+1, s.Area())
		fmt.Printf("  形状 %d 周长: %.2f\n", i+1, s.Perimeter())
	}
}
