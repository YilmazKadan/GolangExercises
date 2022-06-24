package main

import "fmt"

func main() {
	const LENGTH int = 20
	const WIDTH int = 20
	// If you run the code below, you will get an error.
	// Because it's a constant, and its value cannot be changed while streaming.
	// WIDTH = 100
	var area int

	area = LENGTH * WIDTH
	// if you want to use something like code below. You should use "printf" instead of "print" or "println"
	fmt.Printf("Value of area : %d", area)

}
