package linkedList

import "fmt"

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
			fmt.Println("Found : %v\n", name)
			return result
		}
	}
	return nil
}

// Delete node
func (list *SinglyLinkedList) Delete(name string) {
	node2del := list.Find(name)
	if node2del == nil { // There is no such name
		return
	} else if node2del == list.head && node2del == list.tail { // when there is single one node
		list.head = nil
		list.tail = nil
	} else if node2del == list.head { // deleting first node
		list.head = node2del.Next()
	} else { // deleting next one(even if the last node)
		previous_node := list.First()
		for previous_node != node2del {
			previous_node = previous_node.Next()
		}
		list.tail = previous_node
		list.tail.next = nil
	}

}

// Pop last item from list
func (list *SinglyLinkedList) Pop() (p Person, err error) {
	if list.tail == nil {
		err = errEmpty
	} else {
		previous_node := list.First()
		for previous_node != list.tail {
			previous_node = previous_node.Next()
		}
		list.tail = previous_node
		list.tail.next = nil
		p = list.tail.Person
	}
	return p, err
}
