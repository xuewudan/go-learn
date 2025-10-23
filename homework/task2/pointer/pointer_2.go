package main

import "fmt"

func main() {

	arr := []int{1, 3, 5}
	ArrPoint(&arr)
	fmt.Println(arr)
}

func ArrPoint(p *[]int) {
	fmt.Println(*p)
	for i, v := range *p {
		(*p)[i] = v * 2
	}
}
