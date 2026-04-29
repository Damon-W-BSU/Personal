package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

// Node cointaing data within List
type node[T comparable] struct {
	prev *node[T]
	next *node[T]
	data T
}

// Linked list
type List[T comparable] struct {
	head *node[T]
	tail *node[T]
	len  int
}

// invalid index error
var errOOB = errors.New("index out of range")

// Initializes a new list
func New[T comparable]() *List[T] {

	L := List[T]{
		head: nil,
		tail: nil,
		len:  0,
	}

	return &L
}

// Initializes a new data node and returns its pointer
func newNode[T comparable](data T) *node[T] {

	N := node[T]{
		next: nil,
		prev: nil,
		data: data,
	}

	return &N
}

// Appends new data to end of list
func (L *List[T]) Append(data T) {

	n := newNode(data)

	// reassign pointers based on current length
	switch L.len {
	case 0:
		L.head = n
		L.tail = n
	case 1:
		if L.len == 1 {
			L.head.next = n
			L.tail = n
			n.prev = L.head
		}
	default:
		L.tail.next = n
		n.prev = L.tail
		L.tail = n
	}

	L.len++
}

// Inserts new data at index, error if index invalid
func (L *List[T]) Insert(data T, index int) error {

	// check for valid index
	if index < 0 || index >= L.len {
		err := fmt.Errorf("%w: index %d invalid for len %d", errOOB, index, L.len)
		return err
	}

	// empty list case
	if L.head == nil {
		L.Append(data)
		return nil
	}

	// traverse list to index
	curr := L.head
	for range index {
		curr = curr.next
	}

	// insert and reassign pointers
	n := newNode(data)
	if curr == L.head {
		L.head = n
		curr.prev = n
		n.next = curr
	} else {
		n.prev = curr.prev
		curr.prev.next = n
		curr.prev = n
		n.next = curr
	}

	L.len++
	return nil
}

// Returns data at index, error if index invalid
func (L *List[T]) Peek(index int) (T, error) {

	var out T

	// check for valid index
	if index < 0 || index >= L.len {
		err := fmt.Errorf("%w: index %d invalid for len %d", errOOB, index, L.len)
		return out, err
	}

	// traverse list to index
	curr := L.head
	for range index {
		curr = curr.next
	}

	out = curr.data
	return out, nil
}

// Removes and returns data at index
func (L *List[T]) Take(index int) (T, error) {

	var del *node[T] // node to delete
	var out T        // data in node

	// check for valid index
	if index < 0 || index >= L.len {
		err := fmt.Errorf("%w: index %d invalid for len %d", errOOB, index, L.len)
		return out, err
	}

	del = L.head

	// single node list
	if L.len == 1 {
		L.head = nil
		L.tail = nil
	} else {
		// traverse list to index
		for range index {
			del = del.next
		}

		// remove and reassign head/tail
		switch del {
		case L.head:
			L.head = del.next
			del.next.prev = nil
		case L.tail:
			L.tail = del.prev
			del.prev.next = nil
		default:
			del.prev.next = del.next
			del.next.prev = del.prev
			del.next = nil
			del.prev = nil
		}
	}

	// update len and return
	L.len--
	out = del.data
	return out, nil

}

// Searches for data in list, returns index or -1 if unfound
func (L *List[T]) IndexOf(data T) int {

	// track index
	i := 0

	// traverse list
	for n := L.head; n != nil; n = n.next {
		if n.data == data {
			return i
		}
		i++
	}

	// data not found
	return -1
}

// Removes earliest index containing Data, returns index or -1 if unfound
func (L *List[T]) Delete(data T) int {

	// track index
	i := 0

	// traverse list
	for n := L.head; n != nil; n = n.next {
		if n.data == data {
			L.Take(i) // covers removal and reassignnment
			return i
		}
		i++
	}

	// data not found
	return -1

}

// returns length of L
func (L *List[T]) Len() int {
	return L.len
}

// Returns string representation of list
func (L *List[T]) String() string {

	// string builder for concatention
	var sb strings.Builder
	sb.WriteString("[")

	// traverse and append data to output
	for n := L.head; n != nil; n = n.next {
		fmt.Fprintf(&sb, " %v", n.data)
	}

	// complete and return
	sb.WriteString(" ]")
	return sb.String()
}
