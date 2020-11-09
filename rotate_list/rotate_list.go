package main

/*
 * Problem statement
 * Given a singly linked list, rotate the list right by k
 */

import (
	"fmt"
	"math/rand"
)

type node struct { // Linked list node for our purpose.
	value int   // value in the node
	next  *node // next node in the list
}

func createList(n int) (s *node) {
	if n <= 0 {
		return nil
	}

	rand.Seed(0xdeadbeef)

	s = nil
	for i := 0; i < n; i++ {
		if t := new(node); t != nil {
			t.value = rand.Int() % 1000
			t.next = s
			s = t
		} else {
			fmt.Println("Allocation failed")
			break
		}
	}
	return
}

func printList(s *node) {
	for s != nil {
		fmt.Printf("%d ", s.value)
		s = s.next
	}
	fmt.Println()
}

func rotateList(s *node, k int) (ns *node) {

	// Initialize 2 pointers to the start of linked list
	p1 := s
	p2 := s

	// Initialize return value
	ns = nil

	// Place p2 at k-node distance apart
	for i := 0; i < k && p2 != nil; i++ {
		p2 = p2.next
	}

	// If we reach end of the list before we can place p2 k-node distance
	// it is an error, just return
	if p2 == nil {
		fmt.Println("Unable to place 2nd pointer at k-node distance")
		return
	}

	// Now move both pointers, till p2 reaches the last node
	for p2.next != nil {
		p2 = p2.next
		p1 = p1.next
	}

	// Rotate the list by reassigning pointers
	// We break the list and such that list starts from next node after p1

	// last node points to the current start of the list
	p2.next = s

	// New start location of the list is next node after p1
	s = p1.next

	// Current node pointed by p1 is the last node
	p1.next = nil

	// return value initialization
	ns = s
	return
}

func main() {

	fmt.Println("Original List")

	s := createList(10)
	printList(s)
	fmt.Println("Rotated List")
	s = rotateList(s, 2)
	printList(s)
}
