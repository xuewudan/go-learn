package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(palindrome(12321))
}

// 回文数判断
func palindrome(num int) bool {

	numStr := strconv.Itoa(num)

	return numStr == ReverseString(numStr)
}

// 辅助函数：反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
