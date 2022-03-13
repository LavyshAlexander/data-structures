package List

import "fmt"

type List interface {
	Append(value int)
	Delete(index int) error
	Find(value int) (List, int)
	Length() int
}

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (l *LinkedList) Append(value int) {
	for l.Next != nil {
		l = l.Next
	}

	l.Next = &LinkedList{Value: value}
}

func (l *LinkedList) String() string {
	str := "["
	for l.Next != nil {
		str += fmt.Sprintf("%v ", l.Value)
		l = l.Next
	}
	str += fmt.Sprint(l.Value) + "]"

	return str
}
