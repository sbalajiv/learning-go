package btree

/*
 * Naive implementation of a binary search tree
 * the code is part of my learning Go programming language.
 */

import (
    "fmt"
    "math/rand"
)

// BNode : Node of a Binary Tree
type BNode struct {
    value int
    lnode *BNode
    rnode *BNode
}

// Sequence of integers
type _Sequence []int

// SortFlag : For custom enum flags
type SortFlag int

// Custom sort flags, like enums
const (
    _                  = iota
    Ascending SortFlag = iota
    Descending
)

// Find a node where new value can be placed
func findNodeToInsertOnBinarySearchTree(root *BNode, v int) *BNode {
    t := root
    // Traverse the tree
    for t != nil {
        if v > t.value {
            if t.rnode == nil {
                break
            } else {
                t = t.rnode
            }
        } else {
            if t.lnode == nil {
                break
            } else {
                t = t.lnode
            }
        }
    }

    return t
}

// Insert the value v at the determined node
func insertNodeInTreeAt(node *BNode, v int) {
    // create new node and insert into the tree
    n := new(BNode)
    n.value = v
    n.lnode, n.rnode = nil, nil
    if v > node.value {
        node.rnode = n
    } else {
        node.lnode = n
    }
}

// CreateBTree : Create a binary search tree from a sequence of integers
func CreateBTree(s _Sequence) (root *BNode) {
    root = nil
    if len(s) == 0 {
        return
    }

    // Allocate for memory for root node
    root = new(BNode)
    root.lnode = nil
    root.rnode = nil
    if len(s) > 1 {
        // choose a root value
        // just taking the middle element of the sequence
        rindex := len(s) / 2
        root.value = s[rindex]

        // Go through the sequence of numbers
        for i, v := range s {
            if i == rindex {
                continue
            }

            t := findNodeToInsertOnBinarySearchTree(root, v)

            if t != nil {
                insertNodeInTreeAt(t, v)
            } else {
                fmt.Println("Unable to find node to insert")
                return nil
            }

        }
    } else {
        root.value = s[0]
    }
    return
}

// PrintNodeValue : Write to console the 'value' inside the node
func PrintNodeValue(n *BNode) {
    if n != nil {
        fmt.Printf("%d ", n.value)
    }
}

// printTreeAscending : In-order traversal of the tree prints in Ascending sorted order
func printTreeAscending(root *BNode) {
    if root == nil {
        return
    }

    printTreeAscending(root.lnode)
    PrintNodeValue(root)
    printTreeAscending(root.rnode)
}

// printTreeDescending : Post-order traversal of the tree prints in Descending sorted order
func printTreeDescending(root *BNode) {
    if root == nil {
        return
    }

    printTreeDescending(root.rnode)
    PrintNodeValue(root)
    printTreeDescending(root.lnode)
}

// PrintTree : User callable API into this package to print the contents of the tree
func PrintTree(root *BNode, flag SortFlag) {
    if root == nil {
        return
    }

    switch flag {
    case Ascending:
        printTreeAscending(root)
    case Descending:
        printTreeDescending(root)
    }
    fmt.Println()
}

// makeSequence : Creating for list of integers
func makeSequence(n int, withinRange int) (s _Sequence) {
    s = make(_Sequence, 0)

    for i := 0; i < n; i++ {
        s = append(s, rand.Int()%withinRange)
    }

    return
}

// CreateBTreeN : Create a binary tree of 'n' nodes
func CreateBTreeN(n int) *BNode {
    seq := makeSequence(n, 1000)
    return CreateBTree(seq)
}

// Package/module initialization function
func init() {
    rand.Seed(0x1232aeddeadbeef)
}
