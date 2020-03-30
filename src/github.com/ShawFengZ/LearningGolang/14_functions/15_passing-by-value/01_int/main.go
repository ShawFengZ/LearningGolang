package main

import "fmt"

func main() {
	age := 30
	changeAge(age)
	fmt.Println(age) //30
}

func changeAge(z int) int {
	fmt.Println(z)
	z = 20
	return z
}
