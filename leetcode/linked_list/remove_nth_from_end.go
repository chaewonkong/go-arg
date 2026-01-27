package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func _removeNthFromEnd(head *ListNode, n int) *ListNode {
	arr := []*ListNode{}
	node := head
	for {
		arr = append(arr, node)
		if node.Next == nil {
			break
		}
		node = node.Next
	}

	nOffset := len(arr) - n
	var prev, next *ListNode
	if nOffset+1 < len(arr) {
		next = arr[+1]
	}
	if nOffset == 0 { //first node
		return next
	}
	prev = arr[nOffset-1]

	prev.Next = next
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	// Move fast n steps ahead
	for range n {
		fast = fast.Next
	}

	// Move both until fast reaches last node
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Remove slow.Next
	slow.Next = slow.Next.Next

	return dummy.Next
}
