package deletenode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNode(t *testing.T) {
	for i, tt := range []struct {
		input    []int
		given    int
		expected []int
	}{
		{
			input:    []int{4, 5, 1, 9},
			given:    5,
			expected: []int{4, 1, 9},
		},
	} {
		t.Run(fmt.Sprintf("%dth test", i), func(t *testing.T) {
			// create a linked list
			head := &ListNode{}
			target := head
			node := &ListNode{}
			for i, num := range tt.input {
				if i == 0 {
					head.Val = num
					continue
				}
				n := &ListNode{Val: num}
				if num == tt.given {
					node = n
				}
				target.Next = n
				target = n
			}

			// delete node
			deleteNode(node)

			// eval outcome
			result := traverse(head)

			// then
			assert.Equal(t, tt.expected, result)
		})
	}
}

func traverse(node *ListNode) []int {
	result := []int{}

	for {
		result = append(result, node.Val)
		if node.Next == nil {
			break
		}

		node = node.Next
	}

	return result
}
