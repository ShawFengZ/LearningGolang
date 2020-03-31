package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	var x [58]int
	fmt.Printf("%T\n", x)
}
