package main

import "fmt"

// 两数之和：返回和为target的两个元素的下标
func twoSum(nums []int, target int) []int {
	// 哈希表存储已遍历元素的值和下标（值: 下标）
	numMap := make(map[int]int)

	for i, num := range nums {
		// 计算当前元素需要的补数
		complement := target - num
		// 检查补数是否已在哈希表中
		if idx, exists := numMap[complement]; exists {
			// 找到答案，返回补数下标和当前下标
			return []int{idx, i}
		}
		// 补数不存在，将当前元素存入哈希表
		numMap[num] = i
	}

	// 题目保证有唯一答案，此处不会执行
	return nil
}

func main() {
	// 测试用例
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // 输出: [0 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // 输出: [1 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // 输出: [0 1]
	fmt.Println(twoSum([]int{5, 7, 3, 9}, 12))  // 输出: [1 2]（7+5=12，对应下标1和2）
}
