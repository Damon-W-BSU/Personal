package btree

import (
	"errors"
)

const ORDER = 3

// Entry point and manager of tree structure
type BTree struct {
	root  *node
	order int
}

// BTree node containing keys and data
type node struct {
	isLeaf   bool
	keys     []int
	children []*node
}

// Initializes a new B-tree
func New() *BTree {
	return &BTree{
		root:  nil,
		order: ORDER,
	}
}

// Creates a new node
func (bt *BTree) newNode(isLeaf bool) *node {
	return &node{
		isLeaf:   isLeaf,
		keys:     make([]int, 0, bt.order),     // max M - 1 keys
		children: make([]*node, 0, bt.order+1), // max M children
	}
}

// inserts new data into tree
func (bt *BTree) Insert(data int) error {

	// tree is empty
	if bt.root == nil {
		bt.insertEmpty(data)
		return nil
	}

	currentNode := bt.root
	var visited []*node

	for !currentNode.isLeaf {
		// track pointers to traversed parent nodes
		visited = append(visited, currentNode)
		insertIndex := currentNode.keyInsertIndex(data)
		currentNode = currentNode.children[insertIndex]
	}

	if len(currentNode.keys) == cap(currentNode.keys) {
		bt.splitNode(currentNode, visited)
	} else {
		currentNode.insertKey(data)
	}

	return nil
}

// inserts data into empty tree
func (bt *BTree) insertEmpty(data int) {
	bt.root = bt.newNode(true)
	bt.root.insertKey(data)
}

// inserts a key into this node
func (n *node) insertKey(data int) error {

	if len(n.keys) == cap(n.keys) {
		return errors.New("node has reached key capacity")
	}

	if len(n.keys) == 0 {
		n.keys = append(n.keys, data)
		return nil
	}

	insertIndex := n.keyInsertIndex(data)

	// move array data
	last := n.keys[len(n.keys)-1]
	n.keys = append(n.keys, last)
	for i := len(n.keys) - 1; i > insertIndex; i-- {
		n.keys[i] = n.keys[i-1]
	}
	n.keys[insertIndex] = data
	return nil
}

// retrieves the index where key would be inserted
func (n *node) keyInsertIndex(data int) int {
	min := 0
	mid := len(n.keys) / 2
	max := len(n.keys)

	// binary search of key array
	for max != min {
		switch {
		case n.keys[mid] == data:
			max = mid
			min = mid
		case n.keys[mid] > data:
			max = mid
			mid = (max + min) / 2
		case n.keys[mid] < data:
			min = mid + 1
			mid = (max + min) / 2
		}
	}

	return min
}

func (bt *BTree) splitNode(node *node, visitedNodes []*node) {

	// root case gotta go here also

	newChild := bt.newNode(true)

	// save middle key and divide rest
	midIndex := (ORDER - 1) / 2
	midkey := node.keys[midIndex]
	newChild.keys = node.keys[:midIndex]
	node.keys = node.keys[midIndex+1:]

	// divide children between nodes
	midIndex = ORDER / 2
	newChild.children = node.children[:midIndex]
	node.children = node.children[midIndex:]

	// retrieve this node's parent
	parent := visitedNodes[len(visitedNodes)-1]

	// if parent is full, split parent node also
	if len(parent.keys) == cap(parent.keys) {
		bt.splitNode(parent, visitedNodes[:len(visitedNodes)-1])
	}

	parent.insertKey(midkey)

}

// TODO
func (bt *BTree) insertChild() int {
	return 0
}
