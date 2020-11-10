package main

import (
	"fmt"
	"math/rand"
)

// Node of the linked list
type node struct {
	value int
	next  *node
}

// LinkedList is the head node of the linked list
type LinkedList struct {
	head   *node
	length int
}

// PrintList will print the nodes of the linked list to stdout
func (l *LinkedList) PrintList() {
	t := l.head
	for t != nil {
		fmt.Printf("%d ", t.value)
		t = t.next
	}
	fmt.Println()
}

// CreateList will create a list of n nodes
func (l *LinkedList) CreateList(n int) {
	// Validate the input.
	if n <= 0 {
		l.head = nil
		return
	}

	// If head is not nil, assume linked list is already initialized.
	if l.head != nil {
		fmt.Println("List already exists")
		return
	}

	var temp *node = nil
	rand.Seed(0x7eef10afdeadbeef)

	for i := 0; i < n; i++ {
		temp = new(node)

		temp.value = rand.Int() % 10000

		// Insert new node as start of the linked list
		if l.head == nil {
			temp.next = nil
		} else {
			temp.next = l.head
		}
		l.head = temp
	}

	l.length = n
}

// RotateRight will rotate the linked list to the right k times
func (l *LinkedList) RotateRight(k int) (err int) {
	err = -1
	if k <= 0 || l.head == nil {
		return
	}

	p1 := l.head
	p2 := l.head

	// Position p2 at k node distance from p1
	for i := 0; i < k && p2 != nil; i++ {
		p2 = p2.next
	}

	if p2 == nil {
		fmt.Println("Unable to place pointers k-node distance apart")
		err = -2
		return
	}

	// Move both pointers till p2 points to the last node
	for p2.next != nil {
		p1 = p1.next
		p2 = p2.next
	}

	// Reassign the pointers to create rotated list in place
	// Last node points to the current start of the list
	p2.next = l.head

	// Assign new start of the list
	l.head = p1.next

	// Current p1 node becomes the last node in the list
	p1.next = nil

	// If we reach here, then return success
	err = 0
	return
}

func main() {

	// Create head node object
	lst := LinkedList{head: nil, length: 0}

	lst.CreateList(10)
	fmt.Println("Original list")
	lst.PrintList()
	lst.RotateRight(5)
	fmt.Println("Rotated list")
	lst.PrintList()
}
