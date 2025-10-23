package main

import "fmt"

func main() {
	i := 6
	var ip *int
	ip = &i

	fmt.Println(*ip)
	fmt.Println(ip)
	ValuePointer(ip)
	fmt.Println(*ip)
	ValuePointer(&i)
	fmt.Println(i)
}

func ValuePointer(p *int) {
	*p += 10
}
