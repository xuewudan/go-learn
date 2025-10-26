package one

//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	fmt.Println("wudan")
//	fmt.Println("Google" + "Runoob")
//
//	var numbers []int
//
//	/* 允许追加空切片 */
//	numbers = append(numbers, 0, 1)
//	printSlice(numbers)
//
//	fmt.Println()
//
//	/* 同时添加多个元素 */
//	numbers = append(numbers, 2, 3, 4)
//	printSlice(numbers)
//
//	fmt.Println()
//
//	numbers = append(numbers, 2, 3, 4)
//	printSlice(numbers)
//
//	fmt.Println()
//
//	/* 创建切片 numbers1 是之前切片的两倍容量*/
//	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
//
//	/* 拷贝 numbers 的内容到 numbers1 */
//	copy(numbers1, numbers)
//	printSlice(numbers1)
//
//	array := [4]int{10, 20, 30, 40}
//	slice := array[0:2]
//	newSlice := append(slice, 50)
//	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
//	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
//	newSlice[1] += 10
//	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
//	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
//	fmt.Printf("After array = %v\n", array)
//
//	var a int
//	var ptr *int
//	var pptr **int
//
//	a = 3000
//
//	/* 指针 ptr 地址 */
//	ptr = &a
//
//	/* 指向指针 ptr 地址 */
//	pptr = &ptr
//
//	/* 获取 pptr 的值 */
//	fmt.Printf("变量 a = %d\n", a)
//
//	fmt.Printf("变量 &a = %v\n", &a)
//	fmt.Printf("变量 ptr = %v\n", ptr)
//
//	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
//
//	fmt.Printf("指针变量 &(*ptr) = %d\n", &(*ptr))
//
//	fmt.Printf("指针变量 pptr = %v\n", pptr)
//
//	fmt.Printf("指针变量 *pptr = %v\n", *pptr)
//
//	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
//
//	c := make(chan int, 11)
//	fmt.Println(cap(c))
//	go fibonacci(cap(c), c)
//	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
//	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
//	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
//	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
//	for i := range c {
//		fmt.Println(i)
//	}
//
//	c := Circle{Radius: 5}
//	var s Shape = c // 接口变量可以存储实现了接口的类型
//	fmt.Println("Area:", s.Area())
//	fmt.Println("Perimeter:", s.Perimeter())
//}
//
//// Shape 定义接口
//type Shape interface {
//	Area() float64
//	Perimeter() float64
//}
//
//// Circle 定义一个结构体
//type Circle struct {
//	Radius float64
//}
//
//// Area Circle 实现 Shape 接口
//func (c Circle) Area() float64 {
//	return math.Pi * c.Radius * c.Radius
//}
//
//func (c Circle) Perimeter() float64 {
//	return 2 * math.Pi * c.Radius
//}
//
//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//	close(c)
//}
//
//func printSlice(x []int) {
//	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
//}
