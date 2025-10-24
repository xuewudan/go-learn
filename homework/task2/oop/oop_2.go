package main

import "fmt"

// Person 结构体包含基本的姓名和年龄信息
type Person struct {
	Name string
	Age  int
}

// Employee 结构体通过组合Person来复用其字段，并添加员工ID
type Employee struct {
	Person            // 组合Person结构体（匿名嵌入，实现字段和方法的复用）
	EmployeeID string // 员工独有的ID字段
}

// PrintInfo 为Employee实现打印信息的方法
func (e *Employee) PrintInfo() {
	// 可以直接访问组合的Person结构体中的字段（Name和Age）
	fmt.Printf("员工信息：\n")
	fmt.Printf("  ID: %s\n", e.EmployeeID)
	fmt.Printf("  姓名: %s\n", e.Name)
	fmt.Printf("  年龄: %d\n", e.Age)
}

func main() {
	// 创建Employee实例（通过组合Person初始化）
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "EMP001",
	}

	// 调用PrintInfo()方法输出员工信息
	emp.PrintInfo()

	// 也可以直接访问组合的字段和Employee自身的字段
	fmt.Println("\n单独访问字段：")
	fmt.Println("ID:", emp.EmployeeID)
	fmt.Println("姓名:", emp.Name) // 等价于 emp.Person.Name
	fmt.Println("年龄:", emp.Age)  // 等价于 emp.Person.Age
}
