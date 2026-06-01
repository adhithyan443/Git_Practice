package main

import "fmt"

func main() {
	fmt.Println(Add(5, 3))
}

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}
