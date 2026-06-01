package main

import "fmt"

func main() {
	fmt.Println(Add(5, 3))
	fmt.Println(Substract(5, 3))
}

func Add(a, b int) int {
	return a + b
}

func Substract(a, b int) int {
	return a - b
}
