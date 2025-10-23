package main

import (
	"fmt"
)

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArr2(&arr)
	fmt.Println(arr)

	var  ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr  )
	fmt.Printf("ptr 的值为 : %x\n", *ptr  )

}

func modifyArr2(a *[5]int) {
	a[1] = 20
}
