package main

import "fmt"

func main() {

	nums := []int{2, 2, 9}

	fmt.Println(singleNumber1(nums))
}

func singleNumber1(nums []int) int {
	num2Count := map[int]int{}
	for _, num := range nums {
		if num2Count[num] > 0 {
			num2Count[num]++
		} else {
			num2Count[num] = 1
		}
	}
	fmt.Println(num2Count)
	for num, count := range num2Count {
		if count == 1 {
			return num
		}
	}
	return -1
}
