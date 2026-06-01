package main

import "fmt"

func main() {
	fmt.Println(Add(5, 3))
	fmt.Println(Substract(5, 3))
	fmt.Println(Multiply(5, 3))
	fmt.Println(Divide(10, 2))
}

func Add(a, b int) int {
	return a + b
}

func Substract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}
func Divide(a, b int) int {
	return a / b
}