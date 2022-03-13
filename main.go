package main

import (
	"fmt"

	"github.com/LavyshAlexander/data-structures/List"
)

func main() {
	list := &List.LinkedList{Value: 10}

	list.Append(110)
	fmt.Println(list)
}
