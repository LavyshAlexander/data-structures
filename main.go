package main

import "github.com/LavyshAlexander/data-structures/LinkedList"

func main() {
	l := LinkedList.LinkedList[int]{}
	l.Append(10)
	l.Append(20)
	l.Append(30)

	l.PrintReverse()
}
