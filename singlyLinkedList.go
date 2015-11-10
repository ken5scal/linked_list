package linkedList

import (
	"errors"
	"fmt"
)

type SinglyLinkedList struct {
	head, tail *Node
}

type SinglyLinkedListNumber struct {
	head, tail *NumNode
}

type NumNode struct {
	value int
	next  *NumNode
}

func (list *SinglyLinkedList) First() *Node {
	return list.head
}

// Append node
func (list *SinglyLinkedList) Push(p Person) {
	node := &Node{Person: p}
	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.tail = node
}

// Find node
func (list *SinglyLinkedList) Find(name string) *Node {
	var result *Node = nil
	for n := list.First(); n != nil; n = n.Next() {
		if n.Person.Name == name {
			result = n
			fmt.Printf("Found : %v\n", name)
		}
	}
	if result == nil {
		fmt.Printf("Not FOund : %v\n", name)
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
	list.head = list.tail
	list.tail = currentNode

	var previousNode *Node = nil
	for currentNode != nil {
		temp := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = temp
	}
}

func (list *SinglyLinkedList) RemoveDuplicates() {
	currentNode := list.First()
	var previousNode *Node = nil //非初期化(nil)の値は全型に対して共通の値なので、previousNode := nilだとNg

	nameExistMap := make(map[string]bool)

	for currentNode != nil { // tempを使わなくてもfor n:= list.First(); n!=nil; n = n.Next()でもいける
		temp := currentNode.next
		if _, ok := nameExistMap[currentNode.Name]; ok {
			//if nameExistMap[currentNode.Name] { // でもいける
			previousNode.next = currentNode.Next()
		} else {
			nameExistMap[currentNode.Name] = true
			previousNode = currentNode
		}
		currentNode = temp
	}
}

func (list *SinglyLinkedList) FindInReverseOrder(idx int) *Node {
	node1 := list.First()
	node2 := list.First()

	for i := 0; i < idx; i++ {
		if node2 == nil {
			return nil
		}
		node2 = node2.Next()
	}

	for node2.Next() != nil {
		node1 = node1.Next()
		node2 = node2.Next()
	}
	return node1
}

func (list *SinglyLinkedList) DeleteNonHeadNode(node *Node) {
	if node == nil || node.Next() == nil {
		// この場合、最終nodeは削除できない
		return
	}
	node.Person = node.Next().Person
	node.next = node.Next().Next()
}

func (list *SinglyLinkedList) insertNodeAndSortList(x string) *SinglyLinkedList {
	beforeList := new(SinglyLinkedList)
	afterList := new(SinglyLinkedList)

	currentNode := list.First()

	// listの分割
	for currentNode != nil {
		next := currentNode.Next()
		currentNode.next = nil

		if currentNode.Name < x {
			// 前半のリストの最後にnodeを挿入
			if beforeList.head == nil {
				beforeList.head = currentNode
				beforeList.tail = beforeList.First()
			} else {
				beforeList.tail.next = currentNode
				beforeList.tail = currentNode
			}
		} else {
			// 後半のリストの最後にnodeを挿入
			if afterList.head == nil {
				afterList.head = currentNode
				afterList.tail = afterList.First()

			} else {
				afterList.tail.next = currentNode
				afterList.tail = currentNode
			}
		}
		currentNode = next
	}

	if beforeList == nil {
		return afterList
	}

	// beforeListとafterListをマージ
	beforeList.tail.next = afterList.head
	return beforeList
}

func addNumList(l1 SinglyLinkedListNumber, l2 SinglyLinkedListNumber) *SinglyLinkedListNumber {
	list := new(SinglyLinkedListNumber)
	node1 := l1.head
	node2 := l2.head
	val := 0
	var node *NumNode = nil

	for &node1 != nil || &node2 != nil {
		if &node1.value != nil {
			val += node1.value
		}
		if &node2.value != nil {
			val += node2.value
		}

		node.value = val % 10
		node.next = nil

		if list.head == nil {
			list.head = node
		} else {
			list.tail.next = node
			list.tail = list.tail.next
		}

		node1 = node1.next
		node2 = node2.next
		val /= val

	}

	return list
}
