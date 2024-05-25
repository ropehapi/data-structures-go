package main

import "fmt"

type Person struct {
	Name     string
	Nickname string
	Age      int
	Gender   string
}

func main() {
	people := []Person{
		{"Pedro", "Yoshimura", 21, "M"},
		{"Pedro", "Igarashi", 20, "M"},
		{"Giovanna", "Zampa", 21, "F"},
		{"Lorena", "Guidoni", 22, "F"},
	}

	table := HashTable{}

	keys := make([]int, len(people))
	for i, person := range people {
		keys[i] = table.Add(person)
	}

	for _, key := range keys {
		ps := table.Get(key)
		for _, p := range ps {
			fmt.Println(p.Name, p.Nickname, p.Age, p.Gender)
		}
	}

	table.Remove("Lorena")

	pedro := table.Search("Pedro")
	fmt.Println(pedro)
}
