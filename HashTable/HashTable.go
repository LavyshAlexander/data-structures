package HashTable

import "github.com/LavyshAlexander/data-structures/LinkedList"

type (
	HashTable[T comparable] struct {
		table []LinkedList.LinkedList[T]
	}
)

func New[T comparable]() *HashTable[T] {
	return &HashTable[T]{table: make([]LinkedList.LinkedList[T], 0)}
}

// TODO: Think of hash function, hash to index
