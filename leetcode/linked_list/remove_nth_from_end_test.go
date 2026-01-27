package linkedlist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	for _, tt := range []struct {
		head     *ListNode
		n        int
		expected []int
	}{} {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			assert.Equal(t, removeNthFromEnd(tt.head, tt.n), tt.expected)
		})
	}
}

func NewLinkedListHead(items []int) *ListNode {
	first := items[0]
	node := &ListNode{
		Val: first,
	}

	for i, item := range items {
		if i == 0 {
			continue
		}
		cur := &ListNode{
			Val: item,
		}
		node.Next = cur
	}

	return node
}

func ToIntArr(head *ListNode) []int {
	arr := []int{}
	node := head
	for {
		if node.Next == nil {
			return arr
		}

		arr = append(arr, node.Val)
	}
}
