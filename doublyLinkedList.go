package linkedList

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Node struct {
	Person
	next, prev *Node
}

type DoublyLinkedList struct {
	head, tail *Node
}

func (l *DoublyLinkedList) First() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

// Append Node
func (l *DoublyLinkedList) Push(p Person) *DoublyLinkedList {
	n := &Node{Person: p}
	if l.head == nil { // First Node
		l.head = n
	} else {
		l.tail.next = n // link to next
		n.prev = l.tail // link to previous
	}
	l.tail = n // update tail to new node
	return l
}

// Find Node
func (l *DoublyLinkedList) Find(name string) *Node {
	found := false
	var result *Node = nil
	for n := l.First(); n != nil && !found; n = n.Next() {
		if n.Person.Name == name {
			found = true
			result = n
			fmt.Printf("Found : %v\n", name)
			return result
		}
	}
	fmt.Printf("Not Found : %v\n", name)
	return result
}

// Delete Node
func (l *DoublyLinkedList) Delete(name string) bool {
	success := false
	node2del := l.Find(name)
	if node2del != nil {
		fmt.Println("Delete - Found: ", name)
		prev_node := node2del.prev
		next_node := node2del.next

		if prev_node != nil { // Not the First Node
			prev_node.next = next_node
		} else {
			l.head = next_node
		}
		if next_node != nil { // Not the Last Node
			next_node.prev = prev_node
		} else {
			l.tail = prev_node
		}
		success = true
	}
	if success {
		fmt.Printf("Removed: %v\n", name)
	} else {
		fmt.Printf("Failed Removing: %v\n", name)
	}
	return success
}

var errEmpty = errors.New("Error - List is empty")

// Pop last item from list
func (l *DoublyLinkedList) Pop() (p Person, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		p = l.tail.Person
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}
	return p, err
}
