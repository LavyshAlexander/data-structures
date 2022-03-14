package main

import (
	"fmt"

	"github.com/LavyshAlexander/data-structures/list/LinkedList"
)

func main() {
	list := LinkedList.LinkedList{}
	list.Append(10)
	fmt.Println(list.String())
}
