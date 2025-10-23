package main

import "fmt"

// 查找字符串数组的最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "" // 空数组返回空
	}
	// 以第一个字符串为基准
	base := strs[0]
	// 遍历基准字符串的每个字符位置
	for i := 0; i < len(base); i++ {
		// 检查其他所有字符串的相同位置
		for _, s := range strs[1:] {
			// 若字符串长度不足或字符不匹配，返回当前前缀
			if i >= len(s) || s[i] != base[i] {
				return base[:i]
			}
		}
	}
	// 所有字符均匹配，返回基准字符串
	return base
}

func main() {
	// 测试用例
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // 输出: "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // 输出: ""
	fmt.Println(longestCommonPrefix([]string{"a"}))                        // 输出: "a"
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))                  // 输出: "a"
	fmt.Println(longestCommonPrefix([]string{}))                           // 输出: ""
}
