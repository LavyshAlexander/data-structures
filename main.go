package main

import (
	"fmt"

	"github.com/LavyshAlexander/data-structures/list"
)

func main() {
	l := &list.LinkedList{}

	l.Append(10)
	l.Append(20)
	l.Append(30)
	l.Append(42)
	fmt.Println(l)

	fmt.Println("Is there value equals 42?", l.Exist(42))

	err := l.Delete(42)
	fmt.Println(l, err)

	err = l.Delete(20)
	fmt.Println(l, err)

	err = l.Delete(10)
	fmt.Println(l, err)
}
