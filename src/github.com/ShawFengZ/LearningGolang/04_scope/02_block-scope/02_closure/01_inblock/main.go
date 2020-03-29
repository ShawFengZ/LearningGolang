package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "This is in block"
		fmt.Println(y)
	}
	//no access to y
	//fmt.Println(y)
}
