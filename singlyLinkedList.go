package linkedList

import (
	"errors"
	"fmt"
)

type SinglyLinkedList struct {
	head, tail *Node
}

func (list *SinglyLinkedList) First() *Node {
	return list.head
}

// Append node
func (list *SinglyLinkedList) Push(p Person) *SinglyLinkedList {
	node := &Node{Person: p}
	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.tail = node
	return list
}

// Find node
func (list *SinglyLinkedList) Find(name string) *Node {
	var result *Node = nil
	for n := list.First(); n != nil; n = n.Next() {
		if n.Person.Name == name {
			result = n
			fmt.Printf("Found : %v菫ハ", name)
		}
	}
	if result == nil {
		fmt.Printf("Not FOund : %v菫ハ", name)
	}
	return result
}

// Delete node
func (list *SinglyLinkedList) Delete(name string) {
	node2del := list.Find(name)
	if node2del == nil { // There is no such name
		fmt.Println("There is nothing to Delete")
		return
	} else if node2del == list.head && node2del == list.tail { // when there is single one node
		list.head = nil
		list.tail = nil
	} else if node2del == list.head { // deleting first node
		list.head = node2del.Next()
	} else {
		previous_node := list.First()
		for node := list.First(); node != node2del; node = node.Next() {
			previous_node = node
		}
		if node2del == list.tail {
			fmt.Println("Deleting tail")
			list.tail = previous_node
			list.tail.next = nil
		} else {
			fmt.Println("Deleting non tail")
			previous_node.next = node2del.Next()
		}
	}
	// deleting next one(even if the last node)
	fmt.Printf("Removed: %v菫ハ", name)
}

// Pop last item from list
func (list *SinglyLinkedList) Pop() (p Person, err error) {
	var errEmpty = errors.New("Error - List is empty")

	if list.tail == nil {
		err = errEmpty
	} else if list.tail == list.head {
		p = list.head.Person
		list.head = nil
		list.tail = nil
	} else {
		previous_node := list.First()
		for node := list.First(); node != list.tail; node = node.Next() {
			previous_node = node
		}
		p = previous_node.next.Person

		list.tail = previous_node
		list.tail.next = nil
	}

	return p, err
}

func (list *SinglyLinkedList) Reverse() {
	currentNode := list.First()
	origHead := list.First()
	origTail := list.tail
	var previousNode *Node = nil
	for {
		if currentNode == nil {
			break
		}
		temp := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = temp
	}

	list.tail = origHead
	list.head = origTail
}
