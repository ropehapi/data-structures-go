package main

import "sync"

type HashTable struct {
	items map[int][]Person
	Lock  sync.RWMutex
}

func hash(name string) (key int) {
	for _, letter := range name {
		key = 21*key + int(letter)
	}
	return
}

func (ht *HashTable) Add(person Person) int {
	ht.Lock.Lock()
	defer ht.Lock.Unlock()

	if ht.items == nil {
		ht.items = make(map[int][]Person)
	}

	key := hash(person.Name)
	ht.items[key] = append(ht.items[key], person)
	return key
}

func (ht *HashTable) Remove(name string) {
	ht.Lock.Lock()
	defer ht.Lock.Unlock()

	key := hash(name)
	delete(ht.items, key)
}

func (ht *HashTable) Get(key int) []Person {
	ht.Lock.RLock()
	defer ht.Lock.RUnlock()

	return ht.items[key]
}

func (ht *HashTable) Search(name string) []Person {
	ht.Lock.RLock()
	ht.Lock.RUnlock()

	key := hash(name)
	return ht.items[key]
}
