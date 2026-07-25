package btree

import (
	"testing"
)

func TestCreate(t *testing.T) {

	bt := New()
	if bt == nil {
		t.Fail()
	}

}

func TestNewNode(t *testing.T) {

	bt := New()
	node := bt.newNode(true)
	t.Log(node)
	if node.isLeaf != true {
		t.Fail()
	}

}

func TestInsertEmpty(t *testing.T) {
	bt := New()
	err := bt.Insert(0)
	if err != nil {
		t.Fail()
	}
}

/*
	insertKey()
*/

func TestInsertKeyEmpty(t *testing.T) {
	bt := New()
	node := bt.newNode(true)
	node.insertKey(0)
	if node.keys[0] != 0 {
		t.Fail()
	}
}

func TestInsertTwoKeys(t *testing.T) {
	bt := New()
	node := bt.newNode(true)

	node.insertKey(0)
	node.insertKey(1)
	if node.keys[0] != 0 || node.keys[1] != 1 {
		t.Fail()
		t.Log(node.keys)
	}
}

func TestInsertThreeKeysAscending(t *testing.T) {
	bt := New()
	node := bt.newNode(true)

	node.insertKey(0)
	node.insertKey(1)
	node.insertKey(2)
	if node.keys[0] != 0 || node.keys[1] != 1 || node.keys[2] != 2 {
		t.Fail()
		t.Log(node.keys)
	}
}

func TestInsertThreeKeysDescending(t *testing.T) {
	bt := New()
	node := bt.newNode(true)

	node.insertKey(2)
	node.insertKey(1)
	node.insertKey(0)
	if node.keys[0] != 0 || node.keys[1] != 1 || node.keys[2] != 2 {
		t.Fail()
		t.Log(node.keys)
	}
}

func TestInsertSeveralKeys(t *testing.T) {

	nums := []int{5, 3, 6, 9, 8, 1, 2, 4, 0, 7}

	bt := New()
	node := bt.newNode(true)

	for _, val := range nums {
		node.insertKey(val)
		t.Log(node.keys)
	}

	for i := 0; i < len(nums)-1; i++ {
		if i > i+1 {
			t.Fail()
		}
	}

}
