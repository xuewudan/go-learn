package main

import "fmt"

// 原地删除重复元素，返回唯一元素的个数
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0 // 空数组直接返回0
	}
	i := 0 // 慢指针，指向当前最后一个唯一元素
	// 快指针j遍历数组，寻找新的唯一元素
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++               // 慢指针后移，准备接收新的唯一元素
			nums[i] = nums[j] // 更新唯一元素序列
		}
	}
	return i + 1 // 唯一元素个数为慢指针索引+1
}

func main() {
	// 测试用例
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Printf("k=%d, nums前k个元素: %v\n", k1, nums1[:k1]) // 输出: k=2, nums前k个元素: [1 2]

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Printf("k=%d, nums前k个元素: %v\n", k2, nums2[:k2]) // 输出: k=5, nums前k个元素: [0 1 2 3 4]

	nums3 := []int{}
	k3 := removeDuplicates(nums3)
	fmt.Printf("k=%d\n", k3) // 输出: k=0

	nums4 := []int{5}
	k4 := removeDuplicates(nums4)
	fmt.Printf("k=%d, nums前k个元素: %v\n", k4, nums4[:k4]) // 输出: k=1, nums前k个元素: [5]

	nums5 := []int{2, 2, 2}
	k5 := removeDuplicates(nums5)
	fmt.Printf("k=%d, nums前k个元素: %v\n", k5, nums5[:k5]) // 输出: k=1, nums前k个元素: [2]
}
