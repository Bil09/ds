package methods

import (
	"fmt"
	"os"
)

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

type LinkList struct {
	head *Node
	tail *Node
}

// insert method: add any type of value in the linked list
func (ll *LinkList) Insert(v interface{}) {
	newNode := &Node{
		prev:  nil,
		next:  nil,
		value: v,
	}
	// check if there is a head
	if ll.head != nil {

		//set next to the current head
		nn := ll.head
		ll.head = newNode
		nn.prev = newNode
		ll.head.next = nn

	}

	ll.head = newNode
	t := ll.head
	// looking for the current next value of t value
	// if nil is th tail
	// if not nil set t to the prev next value of t and repeat
	for t.next != nil {
		t = t.next
	}

	ll.tail = t
}

// display method: show all values in the linked list to the console
func (ll *LinkList) Display() {
	list := ll.head
	if ll.head != nil {
		fmt.Printf("Head is: %v\n", ll.head.value)
		fmt.Printf("Tail is: %v\n", ll.tail.value)
		for list.next != nil {
			fmt.Printf(" %v => ", list.value)
			list = list.next
		}
		fmt.Printf("%v\n", ll.tail.value)
	} else {
		os.Exit(1)
	}

}

// reverse method: return a reversed linked list without changing the given list(the new list shold be stored)
func (ll *LinkList) Reverse(newList *LinkList) *LinkList {
	if ll.head != nil {
		var v []interface{}
		l := ll.head
		for l.next != nil {
			v = append(v, l.value)
			l = l.next
		}
		v = append(v, ll.tail.value)
		for _, val := range v {
			newList.Insert(val)
		}
		return newList

	}
	return newList
}
