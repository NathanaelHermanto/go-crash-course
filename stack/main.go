package main

import "fmt"

func main() {
	//init a stack
	var s stack

	//push some values
	s = s.push(5)
	s = s.push(2)
	s = s.push(1)

	s.print()
}

type stack []int

func (s stack) print() {
	for len(s) > 0 {
		//print
		fmt.Println(s[len(s)-1])
		//pop
		s = s.pop()
	}
}

func (s stack) pop() stack {
	return s[:len(s)-1]
}

func (s stack) push(i int) stack {
	return append(s, i)
}
