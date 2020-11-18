package main

import (
	"btree"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("In main")

	s := make([]int, 0)
	for i := 0; i < 20; i++ {
		s = append(s, rand.Int()%100)
	}

	fmt.Println(s)
	root := btree.CreateBTree(s)
	fmt.Print("Root Node: ")
	btree.PrintNodeValue(root)
	fmt.Println()

	btree.PrintTree(root, btree.Ascending)
	btree.PrintTree(root, btree.Descending)

	r := btree.CreateBTreeN(10)
	fmt.Println("New Tree:")
	btree.PrintTree(r, btree.Descending)
}
