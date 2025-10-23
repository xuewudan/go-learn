package main

import (
	"fmt"
	"sort"
)

// 合并重叠区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{} // 空输入直接返回空
	}

	// 按区间起始值升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果，加入第一个区间
	res := [][]int{intervals[0]}

	// 遍历剩余区间，合并重叠部分
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := res[len(res)-1] // 结果中最后一个区间

		if current[0] <= last[1] {
			// 重叠或相邻，合并区间（更新结束值为较大者）
			if current[1] > last[1] {
				res[len(res)-1][1] = current[1]
			}
		} else {
			// 不重叠，直接加入结果
			res = append(res, current)
		}
	}

	return res
}

func main() {
	// 测试用例
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // 输出: [[1 6] [8 10] [15 18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    // 输出: [[1 5]]
	fmt.Println(merge([][]int{{4, 7}, {1, 4}}))                    // 输出: [[1 7]]
	fmt.Println(merge([][]int{{1, 5}, {2, 3}}))                    // 输出: [[1 5]]
	fmt.Println(merge([][]int{}))                                  // 输出: []
	fmt.Println(merge([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}})) // 输出: [[1 16]]
}
