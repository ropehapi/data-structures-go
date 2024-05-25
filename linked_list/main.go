package main

import "fmt"

type Person struct {
	Name     string
	Nickname string
	Age      int
}

func main() {
	list := List{}

	yoshimura := Person{"Pedro", "Yoshimura", 21}
	igarashi := Person{"Pedro", "Igarashi", 20}
	gii := Person{"Giovanna", "Zampa", 21}
	lorena := Person{"Lorena", "Guidoni", 22}

	list.Append(yoshimura)
	list.Append(igarashi)
	list.Append(gii)
	list.Append(lorena)

	list.Display()

	fmt.Println("--------------------------")

	gii = list.Search("Giovanna")
	fmt.Println(gii)

	fmt.Println("--------------------------")

	list.Delete("Lorena")
	list.Display()

	fmt.Println("--------------------------")
}
