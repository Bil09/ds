package main

import (
	"github.com/Bil09/ds/LinkedList/pkg/methods"
)

func main() {
	var l1 methods.LinkList
	l1.Insert(1)
	l1.Insert(17)
	l1.Insert(987)
	l1.Insert(8)
	l1.Display()

	l2 := l1.Reverse(&methods.LinkList{})
	l2.Display()

}
