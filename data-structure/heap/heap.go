package heap

type Heap struct {
	lastIndex int
	list      []int
}

func New() *Heap {
	return &Heap{
		lastIndex: 0,
		list:      []int{0}, // 0번쨰 index는 비워둔다
	}
}

func (h *Heap) Insert(val int) {
	h.lastIndex = h.lastIndex + 1
	h.list = append(h.list, val)

	// 새 노드를 parent 노드와 비교하며 바꿔 올려보낸다
	current := h.lastIndex
	parent := current / 2
	for parent >= 1 && h.list[parent] < h.list[current] {
		h.list[parent], h.list[current] = h.list[current], h.list[parent]
		current = parent
		parent = current / 2
	}
}

func (h *Heap) RemoveMax() (int, bool) {
	if h.lastIndex == 0 {
		return 0, false
	}

	root := h.list[1]
	h.list[1] = h.list[h.lastIndex]
	h.list = h.list[:h.lastIndex] // TODO: check
	h.lastIndex = h.lastIndex - 1

	i := 1
	for i <= h.lastIndex {
		swap := i
		if 2*i <= h.lastIndex && (h.list[swap] < h.list[2*i]) {
			swap = 2 * i
		}
		if 2*i+1 <= h.lastIndex && (h.list[swap] < h.list[2*i+1]) {
			swap = 2*i + 1
		}

		if swap != i {
			h.list[i], h.list[swap] = h.list[swap], h.list[i]
			i = swap
		} else {
			break
		}
	}
	return root, true
}
