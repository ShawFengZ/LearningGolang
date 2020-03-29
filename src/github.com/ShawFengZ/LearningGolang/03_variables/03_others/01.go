package main

import "fmt"

func main() {

	//declareVaiable()
	//declareManyAtOnce()
	//initManyAtOnce()
	//inferType()
	//inferMixedType()
}

func initShortedHand() {
	// you can only do this inside a func
	message := "Hello World!"
	a, b, c := 1, false, 3
	d := 4
	e := true

	fmt.Println(message, a, b, c, d, e)
}

func inferMixedType() {
	var message = "Hello World!"
	var a, b, c = 1, false, 3

	fmt.Println(message, a, b, c)
}

func inferType() {
	var message = "Hello World!"
	var a, b, c = 1, 2, 3

	fmt.Println(message, a, b, c)
}

func initManyAtOnce() {
	var message = "Hello World!"
	var a, b, c int = 1, 2, 3

	fmt.Println(message, a, b, c)
}

func declareManyAtOnce() {
	var message string
	var a, b, c int
	a = 1

	message = "Hello World!"

	fmt.Println(message, a, b, c)
}

//string
func declareVaiable() {
	var message string
	message = "Hello World."
	fmt.Println(message)
}
