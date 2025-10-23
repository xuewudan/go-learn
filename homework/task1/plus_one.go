package main

import "fmt"

// 大整数加1
func plusOne(digits []int) []int {
	carry := 1 // 初始进位为1（因为要加1）
	n := len(digits)

	// 从最后一位开始处理进位
	for i := n - 1; i >= 0 && carry > 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10 // 保留个位数
		carry = sum / 10     // 更新进位（0或1）
	}

	// 若仍有进位，说明所有位都进位了，需在最前面加1
	if carry > 0 {
		return append([]int{1}, digits...)
	}

	return digits
}

func main() {
	// 测试用例
	fmt.Println(plusOne([]int{1, 2, 3}))    // 输出: [1 2 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // 输出: [4 3 2 2]
	fmt.Println(plusOne([]int{9}))          // 输出: [1 0]
	fmt.Println(plusOne([]int{9, 9, 9}))    // 输出: [1 0 0 0]
	fmt.Println(plusOne([]int{1, 9, 9}))    // 输出: [2 0 0]
	fmt.Println(plusOne([]int{0}))          // 输出: [1]
}
