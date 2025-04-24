package linkedlist

// add to head InsertAt, Delete, Search, Reverse

type Node struct {
	Val int
	Next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}
func (ll *LinkedList) Prepend(node *Node) *LinkedList {
	if node == nil {return nil}
	if ll.head == nil {
		ll.head = node
		ll.tail = node
		ll.size += 1
		return ll
	}
	tmp := ll.head
	ll.head = node
	ll.head.Next = tmp
	ll.size += 1
	return ll
}

func (ll *LinkedList) Append(node *Node) *LinkedList {
	if node == nil {return nil}
	node.Next = nil
	if ll.tail != nil {
		ll.tail.Next = node
		ll.tail = node
		ll.size += 1
		return ll
	} 
	return ll.Prepend(node)
}

func (ll *LinkedList) Search(val int) *Node {
	if ll.head == nil {
		return nil
	}
	current := ll.head 
	for current != nil {
		if current.Val == val {
			return current 
		}
		current = current.Next
	}
	return nil
}

func (ll *LinkedList) Contains(val int) bool {
	return ll.Search(val) != nil 
}

func (ll *LinkedList) Delete(val int) *LinkedList {
	if ll.head == nil {
		return ll
	}
	if ll.head.Val == val {
		ll.head = ll.head.Next
		return ll
	}
	prev := ll.head
	current := prev.Next
	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
			// If we delete the tail, set prev to tail.
			if ll.tail == current {
				ll.tail = prev
			}
			ll.size -= 1
			return ll
		}
		prev = current
		current = current.Next
	}
	return ll
}

func (ll *LinkedList) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

func (ll *LinkedList) Size() int {
    return ll.size
}

func (ll *LinkedList) InsertAt(index int, node *Node) {
    size := ll.Size()
    if index < 0 || index > size {
        panic("out of bounds") 
    }

    if index == 0 {
        ll.Prepend(node)
        return
    }

    if index == size {
        ll.Append(node)
        return
    }

    prev := ll.head
    current := prev.Next

    for i := 1; i < index; i++ {
        prev = current
        current = current.Next
    }
    prev.Next = node
    node.Next = current
    ll.size += 1
}

func (ll *LinkedList) Reverse() {
	var prev *Node
	current := ll.head
	ll.tail = ll.head // new tail will be the current head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.head = prev
}




func createNode(n int) *Node {
	return &Node{
		Val: n,
		Next: nil,
	}
	return nil
}