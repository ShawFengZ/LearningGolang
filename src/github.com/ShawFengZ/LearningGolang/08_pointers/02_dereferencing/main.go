package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xc00000c0b8

	var b = &a
	fmt.Println(b)  // 0xc00000c0b8
	fmt.Println(*b) // 43
}
