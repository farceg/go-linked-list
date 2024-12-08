package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) print() {
	node := list.Head
	if node == nil {
		fmt.Println("Empty list")
		return
	}
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}

func (list *LinkedList) append(num int) {
	newNode := &Node{Value: num}
	if list.Head == nil {
		list.Head = newNode
	} else {
		node := list.Head
		for node.Next != nil {
			node = node.Next
		}
		node.Next = newNode
	}
}

func (list *LinkedList) deleteLast() {
	if list.Head == nil {
		fmt.Println("Empty list")
		return
	}

	if list.Head.Next == nil {
		list.Head = nil
		return
	}

	node := list.Head
	for node.Next.Next != nil {
		node = node.Next
	}
	node.Next = nil

}

func (list *LinkedList) deleteFirst() {
	if list.Head == nil {
		fmt.Println("Empty list")
		return
	}

	list.Head = list.Head.Next
}

func (list *LinkedList) reverse() {
	var prev, curr, next *Node
	curr = list.Head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	list.Head = prev
}

func main() {
	list := LinkedList{}
	list.append(4)
	list.append(3)
	list.append(2)
	list.append(1)

	fmt.Println("Linkedlist:")
	list.print()

	list.reverse()
	fmt.Println("Reversed Linkedlist:")
	list.print()

	list.deleteLast()
	fmt.Println("Deleted last element from Linkedlist:")
	list.print()

	list.deleteFirst()
	fmt.Println("Deleted first element from Linkedlist:")
	list.print()

	list.deleteFirst()
	fmt.Println("Deleted first element from Linkedlist:")
	list.print()

	list.deleteLast()
	fmt.Println("Deleted last element from Linkedlist:")
	list.print()
}
