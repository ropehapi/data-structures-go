package main

import "fmt"

type Person struct {
	Name     string
	Nickname string
	Age      int
}

func main() {
	list := List{}

	fmt.Println("---Append Linked List---")
	yoshimura := Person{"Pedro", "Yoshimura", 21}
	igarashi := Person{"Pedro", "Igarashi", 20}
	gii := Person{"Giovanna", "Zampa", 21}
	lorena := Person{"Lorena", "Guidoni", 22}
	list.Append(yoshimura)
	list.Append(igarashi)
	list.Append(gii)
	list.Append(lorena)
	list.Display()

	fmt.Println("---Search into Linked List---")
	gii = list.Search("Giovanna")
	fmt.Println(gii)

	fmt.Println("---Remove from Linked List---")
	list.Delete("Lorena")
	list.Display()

	fmt.Println("---Prepend Linked List---")
	list.Prepend(lorena)
	list.Display()

	fmt.Println("---Insert at given position on Linked List---")
	list.InsertAtGivenPosition(gii, 4)
	list.Display()

	fmt.Println("---Remove at given position on Linked List---")
	list.DeleteAtGivenPosition(4)
	list.Display()
}
