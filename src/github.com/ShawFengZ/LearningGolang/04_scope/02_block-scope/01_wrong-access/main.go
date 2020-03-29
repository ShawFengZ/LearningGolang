package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)

}

func foo() {
	//no access to x
	//fmt.Println(x)
}
