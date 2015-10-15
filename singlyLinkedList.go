package linkedList

type singlyLinkedList struct {
	head, tail *Node
}

func (list *singlyLinkedList) First() *Node {
	return list.head
}

// Append node
func (list *singlyLinkedList) Push(p Person) *singlyLinkedList {
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
func (list *singlyLinkedList) Find(name string) *Node {
	var result *Node = nil
	for n := list.First(); n != nil; n = n.Next() {
		if n.Person.Name == name {
			result = n
			return result
		}
	}
	return nil
}
