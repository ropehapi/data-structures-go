package main

import "fmt"

func main() {
	stack := Stack{Size: 2}
	stack.Push("Pedro")
	stack.Push("Giovanna")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
