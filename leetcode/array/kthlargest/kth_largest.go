package kthlargest

func findKthLargest(nums []int, k int) int {
	h := newHeap()
	for _, num := range nums {
		h.insert(num)
	}

	var result int
	var ok bool
	for range k {
		result, ok = h.removeMax()
		if !ok {
			break
		}
	}

	return result
}

type heap struct {
	array []int
}

func newHeap() *heap {
	return &heap{
		array: []int{0}, // 0번째 index는 비워둔다
	}
}

func (h *heap) insert(v int) {
	last := len(h.array) // -1 + 1
	h.array = append(h.array, v)

	// 새 노드를 말단에 추가하고 current와 비교하며 크면 올려보낸다
	cur := last
	parent := cur / 2
	for parent >= 1 && h.array[parent] < h.array[cur] {
		h.array[parent], h.array[cur] = h.array[cur], h.array[parent]
		cur = parent
		parent = cur / 2
	}
}

func (h *heap) removeMax() (int, bool) {
	if len(h.array) == 1 { // 0번째 값만 있음
		return 0, false // zero value, false 반환
	}

	// len >= 2
	root := h.array[1]

	// root에 할당: 우선 가장 마지막 노드 할당하고 자리 바꿔 내려가기
	last := len(h.array) - 1
	h.array[1] = h.array[last]
	h.array = h.array[:last]

	last = last - 1
	i := 1
	for i <= last {
		swap := i
		if idx := 2 * i; idx <= last && (h.array[swap] < h.array[idx]) {
			// children이 더 크면 swap
			swap = idx
		}
		if idx := 2*i + 1; idx <= last && (h.array[swap] < h.array[idx]) {
			swap = idx
		}

		if swap != i {
			h.array[i], h.array[swap] = h.array[swap], h.array[i]
			i = swap
		} else {
			break
		}
	}

	return root, true
}
