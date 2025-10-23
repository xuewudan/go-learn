package main

import "fmt"

// 判断括号字符串是否有效
func isValid(s string) bool {
	// 存储右括号对应的左括号
	closingToOpening := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []rune{} // 用切片模拟栈

	for _, c := range s {
		if opening, ok := closingToOpening[c]; ok {
			// 当前字符是右括号，检查栈顶是否匹配
			if len(stack) == 0 {
				return false // 栈为空，无对应左括号
			}
			if stack[len(stack)-1] != opening {
				return false // 括号类型不匹配
			}
			stack = stack[:len(stack)-1] // 弹出栈顶元素
		} else {
			// 当前字符是左括号，压入栈中
			stack = append(stack, c)
		}
	}

	// 栈为空说明所有左括号都被正确闭合
	return len(stack) == 0
}

func main() {
	// 测试用例
	fmt.Println(isValid("()"))     // 输出: true
	fmt.Println(isValid("()[]{}")) // 输出: true
	fmt.Println(isValid("(]"))     // 输出: false
	fmt.Println(isValid("([)]"))   // 输出: false
	fmt.Println(isValid("{[]}"))   // 输出: true
	fmt.Println(isValid(""))       // 输出: true（空字符串有效）
	fmt.Println(isValid("("))      // 输出: false
}
