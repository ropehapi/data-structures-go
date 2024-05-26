package main

import "fmt"

func main() {
	queue := Queue{}
	queue.Enqueue("Pedro Yoshimura")
	queue.Enqueue("Giovanna Zampa")
	queue.Enqueue("Pedro Igarashi")
	queue.Enqueue("Lorena Guidoni")

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
