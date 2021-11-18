package main

import "fmt"

func main() {
	//init an array
	var stack []int

	//push some values
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)

	printStack(stack)
}

func printStack(stack []int) {
	for len(stack) > 0 {
		n := len(stack) - 1
		fmt.Println(stack[n])

		//pop
		stack = stack[:n]
	}
}
