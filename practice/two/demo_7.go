package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println("dongxusha")
	//fmt.Println("xiaojiejie")
	//str1 := fmt.Sprintf("%v hello world %v", "xuetao", "lidan")
	//fmt.Println(str1)
	//cc := "he"
	//fmt.Print(cc)

	//var a int = 20
	//var b int = 21
	//var c int
	//c=a+b
	//fmt.Printf("a = %d\n", c)

	//var a int = 21
	//var b int = 10
	//var c int
	//
	//c = a + b
	//fmt.Printf("第一行 - c 的值为 %d\n", c )
	//c = a - b
	//fmt.Printf("第二行 - c 的值为 %d\n", c )
	//c = a * b
	//fmt.Printf("第三行 - c 的值为 %d\n", c )
	//c = a / b
	//fmt.Printf("第四行 - c 的值为 %d\n", c )
	//c = a % b
	//fmt.Printf("第五行 - c 的值为 %d\n", c )
	//a++
	//fmt.Printf("第六行 - a 的值为 %d\n", a )
	//a=21   // 为了方便测试，a 这里重新赋值为 21
	//a--
	//fmt.Printf("第七行 - a 的值为 %d\n", a )

	//const (
	//	a = iota
	//	b = iota
	//	c = iota
	//)
	//fmt.Print(a, b, c)

	//const (
	//	a = "abc"
	//	b = len(a)
	//	c = unsafe.Sizeof(a)
	//)
	//
	//println(a, b, c)

	//var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//
	//// 从array取，左指针索引为0，右指针为5，切片是从array切的,
	//// 而且cap函数只计算左指针到原array最后的值的个数
	//var slice = array[0:5] // slice ==> {1, 2, 3, 4, 5}
	//i := cap(slice)        // == 9，因为左指针索引为0，到结尾有9个数，cap为9
	//fmt.Println(i)
	//slice = slice[2:] // slice ==> {3, 4, 5}
	//i2 := cap(slice)  // == 7 左指针偏移了2步，所以cap为9-2=7
	//fmt.Println(i2)

	//var a int = 20 /* 声明实际变量 */
	//var ip *int    /* 声明指针变量 */
	//
	//ip = &a /* 指针变量的存储地址 */
	//
	//fmt.Printf("a 变量的地址是: %x\n", &a)
	//
	///* 指针变量的存储地址 */
	//fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	//
	///* 使用指针访问值 */
	//fmt.Printf("*ip 变量的值: %d\n", *ip)

	//getInstance()
	//put("a", "a_put")
	//put("b", "b_put")
	//fmt.Println(get("a"))
	//fmt.Println(get("b"))
	//put("p", "p_put")
	//

	/*var a int64 = 3
	var b int32
	b = int32(a)
	fmt.Printf("b 为 : %d", b)*/

	//sqrt, err := Sqrt(-1)
	//fmt.Print(sqrt, err)

	//go say("world")
	//say("hello")

	//s := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int)
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // 从通道 c 中接收
	//
	//fmt.Println(x, y, x+y)

	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	//ch := make(chan int, 3)
	//
	//// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	//// 而不用立刻需要去同步读取数据
	//ch <- 1
	//ch <- 2
	//ch <- 3
	//
	//// 获取这两个数据
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//c := make(chan int, 10)
	//fmt.Printf("----%d-----\n", cap(c))
	//go fibonacci(cap(c), c)
	//// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	//// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	//// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	//// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	//for i := range c {
	//	fmt.Println(i)
	//}

	//c := make(chan int)
	//quit := make(chan int)
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//fibonacci(c, quit)

	//go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

}

//func spinner(delay time.Duration) {
//	for {
//		for _, r := range `-\|/` {
//			fmt.Printf("\r%c", r)
//			time.Sleep(delay)
//		}
//	}
//}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return f * f, nil
}

type HashMap struct {
	key      string
	value    string
	hashCode int
	next     *HashMap // 结构体指针
}

var table [16](*HashMap)

func initTable() {
	for i := range table {
		table[i] = &HashMap{"", "", i, nil}
	}
}

func getInstance() [16](*HashMap) {
	if table[0] == nil {
		initTable()
	}
	return table
}

func genHashCode(k string) int {
	if len(k) == 0 {
		return 0
	}
	var hashCode int = 0
	var lastIndex int = len(k) - 1
	for i := range k {
		if i == lastIndex {
			hashCode += int(k[i])
			break
		}
		hashCode += (hashCode + int(k[i])) * 31
	}
	return hashCode
}

func indexTable(hashCode int) int {
	return hashCode % 16
}

func indexNode(hashCode int) int {
	return hashCode >> 4
}

func put(k string, v string) string {
	var hashCode = genHashCode(k)
	var thisNode = HashMap{k, v, hashCode, nil}

	var tableIndex = indexTable(hashCode)
	var nodeIndex = indexNode(hashCode)

	var headPtr [16](*HashMap) = getInstance()
	var headNode = headPtr[tableIndex]

	if (*headNode).key == "" {
		*headNode = thisNode
		return ""
	}

	var lastNode *HashMap = headNode
	var nextNode *HashMap = (*headNode).next

	for nextNode != nil && (indexNode((*nextNode).hashCode) < nodeIndex) {
		lastNode = nextNode
		nextNode = (*nextNode).next
	}
	if (*lastNode).hashCode == thisNode.hashCode {
		var oldValue string = lastNode.value
		lastNode.value = thisNode.value
		return oldValue
	}
	if lastNode.hashCode < thisNode.hashCode {
		lastNode.next = &thisNode
	}
	if nextNode != nil {
		thisNode.next = nextNode
	}
	return ""
}

func get(k string) string {
	var hashCode = genHashCode(k)
	var tableIndex = indexTable(hashCode)

	var headPtr [16](*HashMap) = getInstance()
	var node *HashMap = headPtr[tableIndex]

	if (*node).key == k {
		return (*node).value
	}

	for (*node).next != nil {
		if k == (*node).key {
			return (*node).value
		}
		node = (*node).next
	}
	return ""
}
