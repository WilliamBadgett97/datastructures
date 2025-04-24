package linkedlist

import (
	"testing"
)

func TestLinkedList_BasicOperations(t *testing.T) {
	ll := &LinkedList{}

	// Test Append
	ll.Append(createNode(1))
	ll.Append(createNode(2))
	ll.Append(createNode(3))

	if ll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size())
	}

	if ll.head.Val != 1 || ll.tail.Val != 3 {
		t.Errorf("Head or tail not set correctly")
	}

	// Test Prepend
	ll.Prepend(createNode(0))
	if ll.head.Val != 0 {
		t.Errorf("Expected head to be 0, got %d", ll.head.Val)
	}

	// Test Search
	found := ll.Search(2)
	if found == nil || found.Val != 2 {
		t.Errorf("Search failed to find value 2")
	}

	// Test Contains
	if !ll.Contains(3) {
		t.Errorf("Contains failed to find value 3")
	}

	// Test Delete
	ll.Delete(2)
	if ll.Contains(2) {
		t.Errorf("Delete failed to remove value 2")
	}
	if ll.Size() != 3 {
		t.Errorf("Size should be 3 after delete, got %d", ll.Size())
	}

	// Test Reverse
	ll.Reverse()
	if ll.head.Val != 3 || ll.tail.Val != 0 {
		t.Errorf("Reverse failed: head %d, tail %d", ll.head.Val, ll.tail.Val)
	}
}

func TestLinkedList_InsertAt(t *testing.T) {
	ll := &LinkedList{}
	ll.Append(createNode(1))
	ll.Append(createNode(3))
	ll.InsertAt(1, createNode(2))

	if ll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size())
	}

	values := []int{1, 2, 3}
	current := ll.head
	for _, v := range values {
		if current == nil || current.Val != v {
			t.Errorf("Expected %d, got %d", v, current.Val)
		}
		current = current.Next
	}
}
